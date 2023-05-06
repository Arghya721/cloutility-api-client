package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/safespring-community/cloutility-api-client/cloutility"
	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var consumerDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete existing consumer and associated backup node",
	Long: `
The consumer delete command deletes an existing consumer and associated 
backup node.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		consumerDelete()
	},
}

func consumerDelete() {
	var selectedConsumer cloutility.Consumer

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

	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\t%s\n", "Consumer ID", "Business-unit ID", "Name", "Status", "Url")
	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\t%s\n", "-----------", "----------------", "----", "------", "---")

	consumers, _ := client.GetConsumers(bunitId)
	for _, consumer := range consumers {
		if consumer.ID == consumerId {
			selectedConsumer = consumer
		}
	}

	if selectedConsumer.ID == 0 {
		fmt.Fprintf(twriter, "%v\tUser Default: %v\t%s\t%s\t%s\n", consumerId, bunitId, "NOT FOUND", "NOT FOUND", "NOT FOUND")
		return
	}

	if err := client.DeleteConsumer(bunitId, consumerId); err != nil {
		fmt.Fprintf(twriter, "%v\t%v\t%s\t%s\t%s\n", consumerId, bunitId, selectedConsumer.Name, err, selectedConsumer.Href)
		return
	}
	fmt.Fprintf(twriter, "%v\t%v\t%s\t%s\t%s\n", consumerId, bunitId, selectedConsumer.Name, "DELETED", selectedConsumer.Href)
}

func init() {
	consumerCmd.AddCommand(consumerDeleteCmd)

	consumerDeleteCmd.Flags().IntVar(&consumerId, "id", 0, "ID of consumption-unit to delete")
	consumerDeleteCmd.Flags().IntVar(&bunitId, "bunit-id", 0, "ID of business-unit in which the consumption-unit resides")

	// Mark --id as required
	err := consumerDeleteCmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Println("error marking id flag as required: %w", err)
		os.Exit(1)
	}
}
