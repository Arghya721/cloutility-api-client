package cloutapi

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Temporary function for testing purposes
func RunClient() {
	var (
		user me
		// myConsumer consumer
		// myNode     node
	)

	// Initialize client by passing username, password and client_id from config file
	c, err := Init(
		context.TODO(),
		viper.GetString("client_id"),
		viper.GetString("client_origin"),
		viper.GetString("username"),
		viper.GetString("password"),
		viper.GetString("url"),
	)
	if err != nil {
		log.Printf("Error authenticating: %s", err)
		os.Exit(1)
	}

	// if viper.GetBool("debug") {
	// 	log.Println("Token type:", c.TokenType)
	// 	log.Println("Expires:", c.Expires)
	// 	log.Println("Refresh token:", c.RefreshToken)
	// 	log.Println("Access token:", c.AccessToken)
	// }

	user, err = c.GetUser()
	if err != nil {
		log.Println("Error retrieving userdata: ", err)
	}

	fmt.Println("USER: ", user, "\n\n")

	// if viper.GetBool("debug") {
	// 	log.Println(user.Name)
	// 	log.Println(user.BusinessUnit.Name)
	// 	log.Println(user.BusinessUnit.ID)
	// }

	// node, err := c.GetNode(user.BusinessUnitID, user.ID)
	// if err != nil {
	// 	log.Println("Error retrieving nodedata: ", err)
	// }n
	// fmt.Println("NODE1: ", node, "\n\n")

	if viper.GetBool("dry-run") {
		log.Println("Running in dry-run mode, exiting")
		os.Exit(0)
	}

	consumer, err := c.CreateConsumer(user.BusinessUnit.ID, "testar")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(user.BusinessUnit.ID, " ", consumer.ID)
	fmt.Println("CONSUMER: ", consumer, "\n\n")
	node, err := c.CreateNode(user.BusinessUnit.ID, consumer.ID)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("NODE2: ", node, "\n\n")
	node, err = c.DeleteNode(node.ID)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("NODE DELETED: ", node, "\n\n")
	// log.Println(myNode)
	fmt.Println(c.AccessToken)
	ok, err := c.DeleteConsumer(user.BusinessUnit.ID, consumer.ID)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(user.BusinessUnit.ID, " ", consumer.ID)
	fmt.Println("DETETED: ", ok, "\n\n")

	osType, err := c.GetNodeOperatingSystem()
	if err != nil {
		log.Println(err)
	}
	for _, v := range osType {
		fmt.Println(v.Name, v.ID)
	}
	domains, err := c.GetDomains(user.BusinessUnit.ID)
	if err != nil {
		log.Println(err)
	}
	for _, v := range domains {
		fmt.Println(v.Name, v.ID)
	}
}
