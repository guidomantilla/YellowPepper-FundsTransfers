package cmd

import (
	"github.com/spf13/cobra"
)

func CreateRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "app",
		Short: "Application Description",
	}
}
