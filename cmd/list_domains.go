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

// listDomainsCmd represents the listDomains command
var listDomainsCmd = &cobra.Command{
	Use:   "domains",
	Short: "list domains will list the available backup domains",
	Long: `
The list domains command will list all available backup domains supported by
the server.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		listDomains()
	},
}

func listDomains() {
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

	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "ID", "Name", "Description", "Url")
	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "--", "----", "-----------", "---")

	domains, err := client.GetDomains(user.BusinessUnit.ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, domain := range domains {
		fmt.Fprintf(twriter, "%v\t%s\t%s\t%s\n", domain.ID, domain.Name, domain.Description, domain.Href)
	}
}

func init() {
	listCmd.AddCommand(listDomainsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listDomainsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listDomainsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
