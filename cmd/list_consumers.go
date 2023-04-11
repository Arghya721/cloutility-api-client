/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/safespring/cloutility-api-client/cloutapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// consumersCmd represents the consumers command
var consumersCmd = &cobra.Command{
	Use:   "consumers",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		listConsumers()
	},
}

func listConsumers() {
	client, err := cloutapi.Init(
		context.Background(),
		viper.GetString("client_id"),
		viper.GetString("client_origin"),
		viper.GetString("username"),
		viper.GetString("password"),
		viper.GetString("url"),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	twriter := new(tabwriter.Writer)
	twriter.Init(os.Stdout, 8, 8, 1, '\t', 0)
	defer twriter.Flush()

	user, err := client.GetUser()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "ID", "Name", "Creation date", "Url")
	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "--", "----", "-------------", "---")

	cUnits, err := client.GetConsumers(user.BusinessUnit.ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, cUnit := range cUnits {
		fmt.Fprintf(twriter, "%v\t%s\t%s\t%s\n", cUnit.ID, cUnit.Name, cUnit.CreatedDate.Format(time.ANSIC), cUnit.Href)
	}
}

func init() {
	listCmd.AddCommand(consumersCmd)
}
