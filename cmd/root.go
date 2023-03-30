/*
Copyright 2022-2023 (C) Blue Safespring AB

	Programmed by Jan Johansson
        Contributions by Daniel Oquiñena and Patrik Lundin
	All rights reserved for now, will have liberal
	license later
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cloutility-api-client",
	Short: "client used for managing resources in Safespring BaaS 2.0",
	Long: `cloutility-api-client is used for managing resources in
Safespring BaaS 2.0 using the Cloutility REST API.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().Bool("debug", false, "print debug information")
	rootCmd.PersistentFlags().Bool("dry-run", false, "do not actually create anything")

	// Link cobra with viper
	err := viper.BindPFlags(rootCmd.PersistentFlags())
	if err != nil {
		panic(fmt.Errorf("error parsing flags: %w", err))
	}

	viper.SetConfigName("config")
	viper.SetConfigType("properties")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config file: %w", err))
	}

}
