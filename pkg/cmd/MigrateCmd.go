package cmd

import (
	"github.com/spf13/cobra"
)

func CreateMigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
