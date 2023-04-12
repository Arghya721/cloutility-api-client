/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list subcommand",
	Long: `
The list subcommand is used for listing various resources available on the backup server.
	`,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
