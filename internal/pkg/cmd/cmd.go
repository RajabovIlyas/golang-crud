package cmd

import (
	"github.com/RajabovIlyas/golang-crud/internal/pkg/app"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/logger"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(initMigrateCMD, resetMigrateCmd)
}

var rootCmd = &cobra.Command{
	Use:   "golang-crud",
	Short: "Brief description of a CRUD application with authorization and file uploading.",
	Long:  `This application allows users to perform basic operations on data in a database. In this scenario`,
	Run: func(cmd *cobra.Command, args []string) {
		appLogger, err := logger.InitLogger()

		if err != nil {
			log.Fatal().Err(err).Msg("Error initializing logger")
			return
		}

		err = app.Run(appLogger)

		if err != nil {
			appLogger.Fatal().Msg(err.Error())
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error().Err(err).Msgf("cmd.Execute() failed")
		os.Exit(1)
	}
}
