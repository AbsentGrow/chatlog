package pgmigrate

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func execStatements(db *sql.DB, ddl string) error {
	stmts := strings.Split(ddl, ";")
	for _, s := range stmts {
		q := strings.TrimSpace(s)
		if q == "" {
			continue
		}
		if _, err := db.Exec(q); err != nil {
			return err
		}
	}
	return nil
}

func ImportToPostgres(root string, pgURI string) error {
	ddl, err := GenerateDDL(root)
	if err != nil {
		return err
	}
	pgdb, err := sql.Open("postgres", pgURI)
	if err != nil {
		return err
	}
	defer pgdb.Close()
	if err := execStatements(pgdb, ddl); err != nil {
		return err
	}
	dbFiles, err := ListDBFiles(root)
	if err != nil {
		return err
	}
	for _, dbPath := range dbFiles {
		schema := schemaNameFromPath(root, dbPath)
		sdb, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			continue
		}
		rows, err := sdb.Query(`SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%'`)
		if err != nil {
			_ = sdb.Close()
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
			infoRows, err := sdb.Query(`PRAGMA table_info(` + QuoteIdent(tname) + `)`)
			if err != nil {
				continue
			}
			var cols []string
			var pkCols = make(map[int]string)
			for infoRows.Next() {
				var cid int
				var name string
				var dtype string
				var notnull int
				var dflt sql.NullString
				var pk int
				_ = infoRows.Scan(&cid, &name, &dtype, &notnull, &dflt, &pk)
				cols = append(cols, name)
				if pk > 0 {
					pkCols[pk] = name
				}
			}
			_ = infoRows.Close()
			var pkOrdered []int
			for k := range pkCols {
				pkOrdered = append(pkOrdered, k)
			}
			var pk []string
			for _, k := range pkOrdered {
				pk = append(pk, pkCols[k])
			}
			selectSQL := fmt.Sprintf(`SELECT %s FROM %s`, strings.Join(quoteAll(cols), ", "), QuoteIdent(tname))
			selRows, err := sdb.Query(selectSQL)
			if err != nil {
				continue
			}
			for {
				if !selRows.Next() {
					break
				}
				vals := make([]interface{}, len(cols))
				ptrs := make([]interface{}, len(cols))
				for i := range vals {
					ptrs[i] = &vals[i]
				}
				if err := selRows.Scan(ptrs...); err != nil {
					break
				}
				placeholders := make([]string, len(cols))
				for i := range cols {
					placeholders[i] = fmt.Sprintf("$%d", i+1)
				}
				conflict := ""
				if len(pk) > 0 {
					conflict = fmt.Sprintf(` ON CONFLICT (%s) DO NOTHING`, strings.Join(pk, ", "))
				}
				ins := fmt.Sprintf(`INSERT INTO "%s"."%s" (%s) VALUES (%s)%s`,
					schema, tname, strings.Join(quoteAll(cols), ", "), strings.Join(placeholders, ", "), conflict)
				_, _ = pgdb.Exec(ins, vals...)
			}
			_ = selRows.Close()
		}
		_ = sdb.Close()
	}
	return nil
}

func schemaNameFromPath(root, path string) string {
	rel, err := filepathRelSafe(root, path)
	if err != nil {
		return "public"
	}
	parts := strings.Split(rel, string(filepathSeparator()))
	if len(parts) >= 2 {
		return parts[len(parts)-2]
	}
	return "public"
}

func quoteAll(cols []string) []string {
	out := make([]string, len(cols))
	for i, c := range cols {
		out[i] = QuoteIdent(c)
	}
	return out
}

func filepathSeparator() rune {
	return []rune(string(osPathSeparator()))[0]
}

func osPathSeparator() byte {
	var s = "/"
	if strings.Contains(fmt.Sprintf("%T", s), "windows") {
		return '\\'
	}
	return '/'
}

func filepathRelSafe(root, path string) (string, error) {
	r := strings.TrimSuffix(root, string(osPathSeparator()))
	p := strings.TrimSuffix(path, string(osPathSeparator()))
	if strings.HasPrefix(p, r) {
		return strings.TrimPrefix(p, r+string(osPathSeparator())), nil
	}
	return "", fmt.Errorf("not a child")
}
