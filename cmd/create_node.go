package cmd

import (
	"context"
	"fmt"
	"os"

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
		createNode()
	},
}

var (
	name       string
	contact    string
	osType     int
	clientType int
	domain     int
)

func createNode() {
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

	user, err := client.GetUser()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	consumer, err := client.CreateConsumer(user.BusinessUnit.ID, name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(osType)
	fmt.Println(contact)
	node, err := client.CreateNode(user.BusinessUnit.ID, consumer.ID, osType, clientType, domain, int(1), contact)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("SUCCESS: Node %s (ID: %v) created\n", consumer.Name, node.ID)
}

func init() {
	// Init cmd
	createCmd.AddCommand(createNodeCmd)

	// Add flags
	createNodeCmd.Flags().StringVarP(&name, "name", "n", "", "name of the node (required)")
	createNodeCmd.Flags().StringVar(&contact, "contact", "Safespring", "Name to be set as contact")
	createNodeCmd.Flags().IntVar(&osType, "ostype", 3, "Set the os type")
	createNodeCmd.Flags().IntVar(&clientType, "clientype", 2, "Set the client type")
	createNodeCmd.Flags().IntVar(&domain, "domain", 6, "Set the domain to be used")

	// Mark --name as required
	err := createNodeCmd.MarkFlagRequired("name")
	if err != nil {
		fmt.Println("error marking name flag as required: %w", err)
		os.Exit(1)
	}

	// Link cobra with viper
	// err = viper.BindPFlag("node-name", createNodeCmd.Flags().Lookup("name"))
	// if err != nil {
	// 	panic(fmt.Errorf("error parsing node-name flag: %w", err))
	// }
}
