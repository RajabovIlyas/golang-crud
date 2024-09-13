package cmd

import (
	"github.com/RajabovIlyas/golang-crud/internal/pkg/migration"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var initMigrateCMD = &cobra.Command{
	Use:   "migrate:init",
	Short: "init migrate command",
	Long:  "command to initialize the database by command: migrate:init",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := migration.MigrateConfig()

		if err != nil {
			log.Fatal().Err(err).Msg("migrate init failed")
			return
		}

		migration.InitMigrations(db)

		log.Info().Msg("migrate init done")
	},
}
