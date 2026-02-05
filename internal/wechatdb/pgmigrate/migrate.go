package pgmigrate

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Column struct {
	Name       string
	Type       string
	NotNull    bool
	Default    sql.NullString
	PKPosition int
}

type Index struct {
	Name    string
	Unique  bool
	Columns []string
}

type Table struct {
	Schema     string
	Name       string
	Columns    []Column
	PrimaryKey []string
	Indexes    []Index
}

func ListDBFiles(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToLower(d.Name()), ".db") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(files)
	return files, nil
}

func ScanSchemas(root string) ([]Table, error) {
	dbFiles, err := ListDBFiles(root)
	if err != nil {
		return nil, err
	}
	var tables []Table
	for _, dbPath := range dbFiles {
		schema := filepath.Base(filepath.Dir(dbPath))
		if schema == "." || schema == "" {
			schema = "public"
		}
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			continue
		}
		rows, err := db.Query(`SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%'`)
		if err != nil {
			_ = db.Close()
			continue
		}
		var tnames []string
		for rows.Next() {
			var name string
			_ = rows.Scan(&name)
			if len(name) > 0 {
				tnames = append(tnames, name)
			}
		}
		_ = rows.Close()
		for _, tname := range tnames {
			tiRows, err := db.Query(`PRAGMA table_info(` + QuoteIdent(tname) + `)`)
			if err != nil {
				continue
			}
			var cols []Column
			var pkCols = make(map[int]string)
			for tiRows.Next() {
				var cid int
				var name string
				var dtype string
				var notnull int
				var dflt sql.NullString
				var pk int
				_ = tiRows.Scan(&cid, &name, &dtype, &notnull, &dflt, &pk)
				col := Column{
					Name:       name,
					Type:       dtype,
					NotNull:    notnull == 1,
					Default:    dflt,
					PKPosition: pk,
				}
				cols = append(cols, col)
				if pk > 0 {
					pkCols[pk] = name
				}
			}
			_ = tiRows.Close()
			var pkOrdered []int
			for k := range pkCols {
				pkOrdered = append(pkOrdered, k)
			}
			sort.Ints(pkOrdered)
			var pk []string
			for _, k := range pkOrdered {
				pk = append(pk, pkCols[k])
			}
			ilRows, err := db.Query(`PRAGMA index_list(` + QuoteIdent(tname) + `)`)
			if err != nil {
				ilRows = nil
			}
			var idxs []Index
			if ilRows != nil {
				for ilRows.Next() {
					var seq int
					var iname string
					var unique int
					var origin string
					var partial int
					_ = ilRows.Scan(&seq, &iname, &unique, &origin, &partial)
					iiRows, err := db.Query(`PRAGMA index_info(` + QuoteIdent(iname) + `)`)
					var icolumns []string
					if err == nil {
						for iiRows.Next() {
							var icid int
							var icolorder int
							var icolname string
							_ = iiRows.Scan(&icid, &icolorder, &icolname)
							icolumns = append(icolumns, icolname)
						}
						_ = iiRows.Close()
					}
					idxs = append(idxs, Index{
						Name:    iname,
						Unique:  unique == 1,
						Columns: icolumns,
					})
				}
				_ = ilRows.Close()
			}
			tables = append(tables, Table{
				Schema:     schema,
				Name:       tname,
				Columns:    cols,
				PrimaryKey: pk,
				Indexes:    idxs,
			})
		}
		_ = db.Close()
	}
	return tables, nil
}

func GenerateDDL(root string) (string, error) {
	tables, err := ScanSchemas(root)
	if err != nil {
		return "", err
	}
	var b strings.Builder
	schemas := make(map[string]bool)
	for _, t := range tables {
		if !schemas[t.Schema] {
			b.WriteString(fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS "%s";`+"\n", t.Schema))
			schemas[t.Schema] = true
		}
		b.WriteString(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS "%s"."%s" (`, t.Schema, t.Name))
		var colDefs []string
		for _, c := range t.Columns {
			pt := MapType(c.Type, c.PKPosition)
			def := fmt.Sprintf(`"%s" %s`, c.Name, pt)
			if c.NotNull {
				def += ` NOT NULL`
			}
			if c.Default.Valid {
				def += ` DEFAULT ` + NormalizeDefault(c.Default.String)
			}
			colDefs = append(colDefs, def)
		}
		if len(t.PrimaryKey) > 0 {
			colDefs = append(colDefs, fmt.Sprintf(`PRIMARY KEY (%s)`, QuoteJoin(t.PrimaryKey)))
		}
		b.WriteString(strings.Join(colDefs, ", "))
		b.WriteString(");\n")
		for _, idx := range t.Indexes {
			in := fmt.Sprintf(`"%s_%s_%s"`, t.Schema, t.Name, idx.Name)
			unique := ""
			if idx.Unique {
				unique = "UNIQUE "
			}
			if len(idx.Columns) == 0 {
				continue
			}
			b.WriteString(fmt.Sprintf(`CREATE %sINDEX IF NOT EXISTS %s ON "%s"."%s" (%s);`+"\n",
				unique, in, t.Schema, t.Name, QuoteJoin(idx.Columns)))
		}
	}
	return b.String(), nil
}

func MapType(sqliteType string, pkPos int) string {
	t := strings.ToUpper(strings.TrimSpace(sqliteType))
	if pkPos > 0 && (t == "INTEGER" || strings.Contains(t, "INT")) {
		return "BIGINT"
	}
	if strings.Contains(t, "INT") {
		return "BIGINT"
	}
	if strings.Contains(t, "REAL") || strings.Contains(t, "DOUBLE") || strings.Contains(t, "FLOAT") {
		return "DOUBLE PRECISION"
	}
	if strings.Contains(t, "BLOB") {
		return "BYTEA"
	}
	if strings.Contains(t, "TEXT") || strings.Contains(t, "CHAR") || strings.Contains(t, "CLOB") || t == "" {
		return "TEXT"
	}
	if strings.Contains(t, "NUMERIC") || strings.Contains(t, "DECIMAL") {
		return "NUMERIC"
	}
	return "TEXT"
}

func QuoteIdent(s string) string {
	if s == "" {
		return `""`
	}
	return `"` + strings.ReplaceAll(s, `"`, `""`) + `"`
}

func QuoteJoin(cols []string) string {
	var parts []string
	for _, c := range cols {
		parts = append(parts, QuoteIdent(c))
	}
	return strings.Join(parts, ", ")
}

func NormalizeDefault(d string) string {
	trim := strings.TrimSpace(d)
	if trim == "" {
		return "NULL"
	}
	if strings.EqualFold(trim, "NULL") || strings.EqualFold(strings.Trim(trim, `'`), "NULL") {
		return "NULL"
	}
	fnRe := regexp.MustCompile(`^[A-Z_][A-Z0-9_]*\(.+\)$`)
	numRe := regexp.MustCompile(`^-?\d+(\.\d+)?$`)
	if strings.HasPrefix(trim, "'") && strings.HasSuffix(trim, "'") {
		return trim
	}
	if fnRe.MatchString(strings.ToUpper(trim)) {
		return trim
	}
	if numRe.MatchString(trim) {
		return trim
	}
	return `'` + strings.ReplaceAll(trim, `'`, `''`) + `'`
}
