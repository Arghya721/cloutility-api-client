package main

import (
	"context"
	"fmt"
	"os"

	"github.com/safespring-community/cloutility-api-client/cloutapi"
)

const (
	// Set the OS to Linux, ID = 3
	osTypeID = int(3)
	// Set the nodeType to Fileserver, ID = 1
	nodeTypeID = int(1)
	// Use standard domain (180 Days backup retention), ID = 1
	domainID = int(1)
	// Set the clientOptionSet to "DEDUP_AND_COMPRESS", ID = 2
	clientOptionSetID = int(2)
	// Set contact info to IT Departament
	contact = "Company IT Departement"
)

func main() {
	// Get the hostname of the current Workstation
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	// Initilize client using predefined URL and environment variables
	client, err := cloutapi.Init(
		context.Background(),
		os.Getenv("CLIENT_ID"),
		os.Getenv("APIKEY_ORIGIN"),
		os.Getenv("SAFESPRING_USERNAME"),
		os.Getenv("SAFESPRING_PASSWORD"),
		"https://portal-api.backup.sto2.safedc.net",
	)
	if err != nil {
		panic(err)
	}

	// Retrieve userinfo to determine business-unit in which to create backup-node
	user, err := client.GetUser()
	if err != nil {
		panic(err)
	}

	// Get a list of consumers within the current business-unit
	consumers, err := client.GetConsumers(user.UserBUnit.ID)
	if err != nil {
		panic(err)
	}

	// Loop through names and abort if consumer with the name of 'hostname' exists
	for _, consumer := range consumers {
		if consumer.Name == hostname {
			fmt.Printf("A consumer with the name %s already exists\n", hostname)
			os.Exit(1)
		}
	}

	// Proceed to create a new consumer within the context of the username business-unit
	consumer, err := client.CreateConsumer(user.UserBUnit.ID, hostname)
	if err != nil {
		panic(err)
	}

	// Create a backup node within the consumer we just created
	backupNode, err := client.CreateNode(
		user.UserBUnit.ID,
		consumer.ID,
		osTypeID,
		nodeTypeID,
		domainID,
		clientOptionSetID,
		contact,
	)
	if err != nil {
		panic(err)
	}

	// Activate the node
	backupNode, err = client.ActivateNode(user.UserBUnit.ID, consumer.ID)
	if err != nil {
		panic(err)
	}

	// Finally print the userID and password for backupNode separated by a newline
	fmt.Printf("%s\n%s\n", backupNode.TsmName, backupNode.TsmPassword)
}
