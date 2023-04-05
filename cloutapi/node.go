package cloutapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (c *AuthenticatedClient) CreateNode(myid int, myConsumer int) (node, error) {

	var newNode node
	var nodestr string
	createstr := "/v1/bunits/" + strconv.Itoa(myid) + "/consumers/" +
		strconv.Itoa(myConsumer) + "/node"
	data := map[string]interface{}{
		"operatingSystem": map[string]string{
			"name": "Linux",
		},
		"type": map[string]string{
			"name": "File server",
		},
		"server": map[string]string{
			"name": "tsm12.backup.sto2.safedc.net",
		},
		"clientOptionSet": map[string]string{
			"name": "STANDARD",
		},
		"contact":  "Someone",
		"cpuCount": 1,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return node{}, fmt.Errorf("failed to encode json payload: %s", err)
	}

	nodestr, err = c.apiRequest(createstr, http.MethodPost, payload)
	if err != nil {
		return node{}, fmt.Errorf("failed to create node: %s", err)
	}
	if err := json.Unmarshal([]byte(nodestr), &newNode); err != nil {
		return node{}, fmt.Errorf("failed to decode nodedata: %s", err)
	}

	return newNode, nil

}

func (c *AuthenticatedClient) DeleteNode() (string, error) {
	return "", nil
}

func (c *AuthenticatedClient) GetNode() (string, error) {
	node, err := c.apiRequest("/v1/bunits/17/consumers/31/node", "GET", nil)
	if err != nil {
		return "", fmt.Errorf("error requesting nodedata: %s", err)
	}
	return node, nil
	// XXX needs conf or code to use your bUnit/node instead
}
