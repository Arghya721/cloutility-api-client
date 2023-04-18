package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "The domain subcommand",
	Long:  `The domain subcommand is used for operations regarding backup domains.`,
}

func init() {
	rootCmd.AddCommand(domainCmd)
}
