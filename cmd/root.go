/*
 * Copyright (c) Blue Safespring AB - Jan Johansson <jj@safespring.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/safespring-community/cloutility-api-client/cloutapi"
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
	activate        bool
	bunitId         int
	cfgFile         string
	client          *cloutapi.AuthenticatedClient
	clientOptionSet int
	clientType      int
	consumerId      int
	contact         string
	domain          int
	name            string
	osType          int
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
