package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create subcommand",
	Long: `
The create subcommand is used for creating resources like consumers and backup nodes.
	`,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
