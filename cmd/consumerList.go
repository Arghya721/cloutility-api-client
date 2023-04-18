/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

// consumersCmd represents the consumers command
var consumerListCmd = &cobra.Command{
	Use:   "list",
	Short: "consumer list will list the available consumers / consumption-units",
	Long: `
The command 'consumer list' will list all the available consumers / consumption-units 
for the current user account if no business unit ID is provided`,
	Run: func(cmd *cobra.Command, args []string) {
		consumerList()
	},
}

func consumerList() {
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
	consumerCmd.AddCommand(consumerListCmd)
	consumerListCmd.Flags().IntVar(&bunitId, "bunit-id", 0, "ID of business unit in which to list consumers")
}
