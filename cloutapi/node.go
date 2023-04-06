package cloutapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (c *AuthenticatedClient) CreateNode(bUnitID int, consumerID int) (Node, error) {
	var newNode Node

	endpoint := c.BaseURL + "/v1/bunits/" + strconv.Itoa(bUnitID) + "/consumers/" + strconv.Itoa(consumerID) + "/node"

	// TODO: Parameterize
	data := map[string]interface{}{
		"OperatingSystem": map[string]int{
			"ID": 1,
		},
		"Type": map[string]int{
			"ID": 1,
		},
		"Domain": map[string]int{
			"ID": 1,
		},
		"ClientOptionSet": map[string]int{
			"ID": 1,
		},
		"contact":  "Daniel",
		"CpuCount": 1,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return Node{}, fmt.Errorf("failed to encode json payload: %s", err)
	}

	resp, err := c.apiRequest(endpoint, http.MethodPost, payload)
	if err != nil {
		return Node{}, fmt.Errorf("failed to create node: %s", err)
	}

	if err := json.Unmarshal([]byte(resp), &newNode); err != nil {
		return Node{}, fmt.Errorf("failed to decode nodedata: %s", err)
	}

	return newNode, nil
}

func (c *AuthenticatedClient) DeleteNode(id int) (Node, error) {
	var node Node

	endpoint := c.BaseURL + "/v1/bunits/17/consumers/" + strconv.Itoa(id) + "/node?deleteAssociations=True"
	fmt.Println(endpoint)

	resp, err := c.apiRequest(endpoint, http.MethodDelete, nil)
	if err != nil {
		return Node{}, fmt.Errorf("error requesting nodedata: %s", err)
	}
	if err := json.Unmarshal([]byte(resp), &node); err != nil {
		return Node{}, fmt.Errorf("failed to decode nodedata: %s", err)
	}
	return node, nil
}

func (c *AuthenticatedClient) GetNode(userID, consumerID int) (Node, error) {

	var node Node

	endpoint := c.BaseURL + "/v1/bunits/" + strconv.Itoa(userID) + "/consumers/" + strconv.Itoa(consumerID) + "/node"

	resp, err := c.apiRequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return Node{}, fmt.Errorf("error requesting nodedata: %s", err)
	}
	if err := json.Unmarshal([]byte(resp), &node); err != nil {
		return Node{}, fmt.Errorf("failed to decode nodedata: %s", err)
	}
	return node, nil
	// XXX needs conf or code to use your bUnit/node instead
}
