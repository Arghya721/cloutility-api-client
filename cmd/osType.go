/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var osTypeCmd = &cobra.Command{
	Use:   "ostype",
	Short: "ostype subcommand",
	Long: `
The ostype subcommand is used for OSType operations.
	`,
}

func init() {
	rootCmd.AddCommand(osTypeCmd)
}
