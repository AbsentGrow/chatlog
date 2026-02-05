package chatlog

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/sjzar/chatlog/internal/wechatdb/pgmigrate"
)

func init() {
	rootCmd.AddCommand(pgImportCmd)
	pgImportCmd.Flags().StringVarP(&pgImportDir, "dir", "d", "", "sqlite root dir")
	pgImportCmd.Flags().StringVarP(&pgImportURI, "pg", "p", "", "postgres uri")
}

var (
	pgImportDir string
	pgImportURI string
)

var pgImportCmd = &cobra.Command{
	Use:   "pg-import",
	Short: "import sqlite data into postgresql",
	Run: func(cmd *cobra.Command, args []string) {
		if len(pgImportDir) == 0 || len(pgImportURI) == 0 {
			log.Error().Msg("missing dir or pg uri")
			return
		}
		if err := pgmigrate.ImportToPostgres(pgImportDir, pgImportURI); err != nil {
			log.Err(err).Msg("import failed")
			return
		}
	},
}
