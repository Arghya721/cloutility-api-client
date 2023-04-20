package cmd

import (
	"github.com/spf13/cobra"
)

var clientOptionSetCmd = &cobra.Command{
	Use:   "clientoptionset",
	Short: "The clientoptionset subcommand",
	Long:  `The clientoptionset subcommand is used for operations regarding ClientOptionSets.`,
}

func init() {
	rootCmd.AddCommand(clientOptionSetCmd)
}
