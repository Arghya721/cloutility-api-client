/*
Copyright 2022-2023 (C) Blue Safespring AB

		Programmed by Jan Johansson
	        Contributions by Daniel Oquiñena and Patrik Lundin
		All rights reserved for now, will have liberal
		license later
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/safespring/cloutility-api-client/cloutapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "cloutility-api-client",
	Short: "client used for managing resources in Safespring BaaS 2.0",
	Long: `cloutility-api-client is used for managing resources in
Safespring BaaS 2.0 using the Cloutility REST API.`,
}

// Global variables
var (
	name            string
	contact         string
	osType          int
	clientType      int
	domain          int
	cfgFile         string
	consumerId      int
	clientOptionSet int
	client          *cloutapi.AuthenticatedClient
	activate        bool
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initCloutilityApi() {
	c, err := cloutapi.Init(
		context.Background(),
		viper.GetString("client_id"),
		viper.GetString("client_origin"),
		viper.GetString("username"),
		viper.GetString("password"),
		viper.GetString("url"),
	)
	if err != nil {
		log.Fatalf("error initializing client: %s", err)
		os.Exit(1)
	}
	client = c
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("cloutility-api-client")
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initCloutilityApi)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./cloutility-api-client.properties)")
	// rootCmd.PersistentFlags().Bool("debug", false, "print debug information")
	// rootCmd.PersistentFlags().Bool("dry-run", false, "do not actually create anything")

	// Link cobra with viper
	err := viper.BindPFlags(rootCmd.PersistentFlags())
	if err != nil {
		panic(fmt.Errorf("error parsing flags: %w", err))
	}
}
