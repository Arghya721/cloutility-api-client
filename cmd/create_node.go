package cmd

import (
	"fmt"

	"github.com/safespring/cloutility-api-client/cloutapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// nodeCmd represents the node command
var createNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Create new backup node",
	Long:  `This command creates a new backup node that you can then use for TSM backups.`,
	Run: func(cmd *cobra.Command, args []string) {
		cloutapi.RunClient()
	},
}

func init() {
	createCmd.AddCommand(createNodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createNodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createNodeCmd.Flags().String("name", "", "name of the node")
	err := createNodeCmd.MarkFlagRequired("name")
	if err != nil {
		panic(fmt.Errorf("error marking name flag as required: %w", err))
	}

	// Link cobra with viper
	err = viper.BindPFlag("node-name", createNodeCmd.Flags().Lookup("name"))
	if err != nil {
		panic(fmt.Errorf("error parsing node-name flag: %w", err))
	}
}
