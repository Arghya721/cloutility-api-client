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
var deleteConsumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Delete existing consumer and associated backup node",
	Long: `
The command delete consumer deletes an existing consumer and associated 
backup node.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteConsumer()
	},
}

var consumerID int

func deleteConsumer() {
	var selectedConsumer cloutapi.Consumer
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

	consumers, _ := client.GetConsumers(user.UserBUnit.ID)
	for _, consumer := range consumers {
		if consumer.ID == consumerID {
			selectedConsumer = consumer
		}
	}

	if err := client.DeleteConsumer(user.UserBUnit.ID, consumerID); err != nil {
		fmt.Fprintf(twriter, "%v\t%s\t%s\t%s\n", consumerID, selectedConsumer.Name, err, selectedConsumer.Href)
		os.Exit(1)
	}
	fmt.Fprintf(twriter, "%v\t%s\t%s\t%s\n", consumerID, selectedConsumer.Name, "DELETED", selectedConsumer.Href)
}

func init() {
	deleteCmd.AddCommand(deleteConsumerCmd)

	deleteConsumerCmd.Flags().IntVarP(&consumerID, "id", "i", -1, "ID of consumption-unit to delete")

	// Mark --id as required
	err := deleteConsumerCmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Println("error marking id flag as required: %w", err)
		os.Exit(1)
	}
}
