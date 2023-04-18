package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "consumer subcommand",
	Long: `
The consumer subcommand is used for managing consumers and backup nodes.
	`,
}

func init() {
	rootCmd.AddCommand(consumerCmd)
}
