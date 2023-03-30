package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var deleteNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Delete existing backup node",
	Long:  `This command deletes an existing backup node.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete node called")
	},
}

func init() {
	deleteCmd.AddCommand(deleteNodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteNodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteNodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
