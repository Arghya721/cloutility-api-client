/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var bunitCmd = &cobra.Command{
	Use:   "bunit",
	Short: "bunit subcommand",
	Long: `
The bunit subcommand is used for backup-unit operations.
	`,
}

func init() {
	rootCmd.AddCommand(bunitCmd)
}
