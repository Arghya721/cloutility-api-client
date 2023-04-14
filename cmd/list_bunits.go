/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/safespring/cloutility-api-client/cloutapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listBUnitsCmd = &cobra.Command{
	Use:   "bunits",
	Short: "list bunits will list the available business units",
	Long: `
The command will list all the available Operating System types supported by the 
backup server.	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		listBUnits()
	},
}

var bunitId int

func listBUnits() {
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

	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "ID", "Name", "Type", "Group Name")
	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "--", "----", "----", "----------")

	if bunitId == 0 {
		currentUser, err := client.GetUser()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		bunitId = currentUser.UserBUnit.ID
	}

	bunit, err := client.GetBusinessUnit(bunitId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Fprintf(twriter, "%v\t%s\t%s\t%s\n", bunit.ID, bunit.Name, "Parent", bunit.GroupName)

	for _, subUnit := range bunit.BusinessUnits {
		fmt.Fprintf(twriter, "%v\t%s\t%s\t%s\n", subUnit.ID, subUnit.Name, "Decendant", subUnit.GroupName)
	}
}

func init() {
	listCmd.AddCommand(listBUnitsCmd)
	listBUnitsCmd.Flags().IntVarP(&bunitId, "bunit-id", "i", 0, "ID of business unit to list")
}
