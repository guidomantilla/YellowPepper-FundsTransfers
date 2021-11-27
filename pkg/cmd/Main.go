package cmd

import (
	"log"
)

func Execute() {
	rootCmd := CreateRootCmd()
	serveCmd := CreateServeCmd()
	migrateCmd := CreateMigrateCmd()
	rootCmd.AddCommand(serveCmd, migrateCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
