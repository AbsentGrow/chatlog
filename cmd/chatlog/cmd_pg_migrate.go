package chatlog

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/sjzar/chatlog/internal/wechatdb/pgmigrate"
)

func init() {
	rootCmd.AddCommand(pgMigrateCmd)
	pgMigrateCmd.Flags().StringVarP(&pgDir, "dir", "d", "", "sqlite root dir")
	pgMigrateCmd.Flags().StringVarP(&pgOut, "out", "o", "", "output ddl file")
}

var (
	pgDir string
	pgOut string
)

var pgMigrateCmd = &cobra.Command{
	Use:   "pg-migrate",
	Short: "generate PostgreSQL ddl from sqlite files",
	Run: func(cmd *cobra.Command, args []string) {
		if pgDir == "" {
			pgDir = filepath.Join("misc", "decrypt_log", "db_storage")
		}
		ddl, err := pgmigrate.GenerateDDL(pgDir)
		if err != nil {
			log.Err(err).Msg("generate ddl failed")
			return
		}
		if pgOut != "" {
			if err := os.WriteFile(pgOut, []byte(ddl), 0644); err != nil {
				log.Err(err).Msg("write ddl file failed")
				return
			}
			fmt.Println(pgOut)
			return
		}
		fmt.Println(ddl)
	},
}
