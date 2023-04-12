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

	// consumers, err := client.GetConsumers(user.BusinessUnit.ID)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// for _, consumer := range consumers {
	// 	fmt.Println(consumer.ID)
	// }

	if err := client.DeleteConsumer(user.BusinessUnit.ID, consumerID); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("SUCCESS: Node %v (ID: %v) deleted\n", consumerID, consumerID)
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
