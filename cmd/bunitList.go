/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var bunitListCmd = &cobra.Command{
	Use:   "list",
	Short: "bunit list will list the available business units",
	Long: `
The command will list all the available business-units and decendant business-units
which your user account has access to.	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		bunitList()
	},
}

var bunitId int

func bunitList() {
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
	bunitCmd.AddCommand(bunitListCmd)
	bunitListCmd.Flags().IntVar(&bunitId, "bunit-id", 0, "ID of business unit to list")
}
