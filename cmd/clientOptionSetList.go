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

var clientOptionSetListCmd = &cobra.Command{
	Use:   "list",
	Short: "clientoptionset list will list the available ClientOptionSets.",
	Long: `
The command will list all the available Operating System types supported by the 
backup server.	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		clientOptionSetList()
	},
}

func clientOptionSetList() {
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

	fmt.Fprintf(twriter, "%s\t%s\t%s\n", "ID", "Name", "Url")
	fmt.Fprintf(twriter, "%s\t%s\t%s\n", "--", "----", "---")

	sets, err := client.GetClientOptionSet(bunitId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, set := range sets {
		fmt.Fprintf(twriter, "%v\t%s\t%s\n", set.ID, set.Name, set.Href)
	}
}

func init() {
	clientOptionSetCmd.AddCommand(clientOptionSetListCmd)

	clientOptionSetCmd.Flags().IntVar(&bunitId, "bunit-id", 0, "ID of business unit in which to search for ClientOptionSets")
}
