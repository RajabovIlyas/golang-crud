package cmd

import (
	"github.com/RajabovIlyas/golang-crud/internal/pkg/migration"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var resetMigrateCmd = &cobra.Command{
	Use:   "migrate:reset",
	Short: "reset migrate command",
	Long:  "command to reset the database by command: migrate:reset",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := migration.MigrateConfig()

		if err != nil {
			log.Fatal().Err(err).Msg("migrate reset failed")
			return
		}

		migration.ResetMigrations(db)

		log.Info().Msg("migrate reset success")
	},
}
