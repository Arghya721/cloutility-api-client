package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var consumerCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new consumer and/or associated backup node",
	Long: `
The consumer create command creates a new consumer / consumtion-unit. If additional
information regarding OS, domain, contactperson etc is added an associated backup 
node is also created that can be used for TSM backups.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		createConsumer()
	},
}

func createConsumer() {
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

	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "ID", "Name", "Status", "Url")
	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "--", "----", "------", "---")

	consumer, err := client.CreateConsumer(bunitId, name)
	if err != nil {
		fmt.Fprintf(twriter, "%v\t%s\t%s\t%s\n", "N/A", name, err, "N/A")
		return
	}

	_, err = client.CreateNode(bunitId, consumer.ID, osType, clientType, domain, int(1), contact)
	if err != nil {
		fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "N/A", name, err, "N/A")
		return
	}

	fmt.Fprintf(twriter, "%v\t%s\t%s\t%s\n", consumer.ID, consumer.Name, "CREATED", consumer.Href)
}

func init() {
	consumerCmd.AddCommand(consumerCreateCmd)

	// Add flags
	consumerCreateCmd.Flags().StringVar(&name, "name", "", "Name of backup node (required)")
	consumerCreateCmd.Flags().StringVar(&contact, "contact", "Safespring", "Name of contact")
	consumerCreateCmd.Flags().IntVar(&osType, "os-type", 0, "ID of OS Type")
	consumerCreateCmd.Flags().IntVar(&clientType, "client-type", 0, "ID of client type")
	consumerCreateCmd.Flags().IntVar(&domain, "domain", 0, "ID of domain")
	consumerCreateCmd.Flags().IntVar(&bunitId, "bunit-id", 0, "ID of business unit in which to create consumer")

	// Mark --name as required
	err := consumerCreateCmd.MarkFlagRequired("name")
	if err != nil {
		fmt.Println("error marking name flag as required: %w", err)
		os.Exit(1)
	}
}
