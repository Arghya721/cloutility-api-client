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
var listConsumersCmd = &cobra.Command{
	Use:   "consumers",
	Short: "list consumers will list the available consumers / consumption-units",
	Long: `
The command 'list consumers' will list all the available consumers / consumption-units 
for the current user account if no business unit ID is provided`,
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

	if bunitId == 0 {
		user, err := client.GetUser()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		bunitId = user.UserBUnit.ID
	}

	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "ID", "Name", "Creation date", "Url")
	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "--", "----", "-------------", "---")

	cUnits, err := client.GetConsumers(bunitId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, cUnit := range cUnits {
		fmt.Fprintf(twriter, "%v\t%s\t%s\t%s\n", cUnit.ID, cUnit.Name, cUnit.CreatedDate.Format(time.ANSIC), cUnit.Href)
	}
}

func init() {
	listCmd.AddCommand(listConsumersCmd)
	listConsumersCmd.Flags().IntVar(&bunitId, "bunit-id", 0, "ID of business unit in which to list consumers")
}
