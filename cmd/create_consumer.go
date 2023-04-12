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

// nodeCmd represents the node command
var createNodeCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Create new consumer and associated backup node",
	Long: `
The create consumer command creates a new consumer / consumtion-unit and an 
associated backup node that you can be used for TSM backups.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		createConsumer()
	},
}

var (
	name       string
	contact    string
	osType     int
	clientType int
	domain     int
)

func createConsumer() {
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

	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "ID", "Name", "Status", "Url")
	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "--", "----", "------", "---")

	consumer, err := client.CreateConsumer(user.BusinessUnit.ID, name)
	if err != nil {
		fmt.Fprintf(twriter, "%v\t%s\t%s\t%s\n", "N/A", name, err, "N/A")
		os.Exit(1)
	}

	_, err = client.CreateNode(user.BusinessUnit.ID, consumer.ID, osType, clientType, domain, int(1), contact)
	if err != nil {
		fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "N/A", name, err, "N/A")
		os.Exit(1)
	}

	fmt.Fprintf(twriter, "%v\t%s\t%s\t%s\n", consumer.ID, consumer.Name, "CREATED", consumer.Href)
}

func init() {
	createCmd.AddCommand(createNodeCmd)

	// Add flags
	createNodeCmd.Flags().StringVarP(&name, "name", "n", "", "name of the node (required)")
	createNodeCmd.Flags().StringVar(&contact, "contact", "Safespring", "Name to be set as contact")
	createNodeCmd.Flags().IntVar(&osType, "ostype", 3, "Set the os type")
	createNodeCmd.Flags().IntVar(&clientType, "clienttype", 2, "Set the client type")
	createNodeCmd.Flags().IntVar(&domain, "domain", 6, "Set the domain to be used")

	// Mark --name as required
	err := createNodeCmd.MarkFlagRequired("name")
	if err != nil {
		fmt.Println("error marking name flag as required: %w", err)
		os.Exit(1)
	}
}
