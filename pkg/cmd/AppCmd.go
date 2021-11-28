package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func ExecuteAppCmd() {
	appCmd := CreateAppCmd()
	serveCmd := CreateServeCmd()
	migrateCmd := CreateMigrateCmd()
	appCmd.AddCommand(serveCmd, migrateCmd)

	if err := appCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}

func CreateAppCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "app",
		Short: "Application Description",
	}
}
