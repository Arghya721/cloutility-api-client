package cloutapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type nodeOperatingSystems struct {
	Href   string                `json:"href"`
	Total  int                   `json:"total"`
	Offset int                   `json:"offset"`
	First  string                `json:"first"`
	Items  []NodeOperatingSystem `json:"items"`
}

type NodeOperatingSystem struct {
	Href               string               `json:"href"`
	Name               string               `json:"name"`
	ShortName          string               `json:"shortName"`
	SupportedNodeTypes []supportedNodeTypes `json:"supportedNodeTypes"`
	ID                 int                  `json:"id"`
	CreatedDate        time.Time            `json:"createdDate"`
}

type supportedNodeTypes struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CreatedDate time.Time `json:"createdDate"`
	ShortName   string    `json:"shortName"`
}

func (c *AuthenticatedClient) GetNodeOperatingSystem() ([]NodeOperatingSystem, error) {
	var list nodeOperatingSystems
	var nodeOSes []NodeOperatingSystem

	endpoint := c.BaseURL + "/v1/nodeoperatingsystems"

	resp, err := c.apiRequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, fmt.Errorf("error requesting nodeOperatingSystems: %s", err)
	}
	if err := json.Unmarshal([]byte(resp), &list); err != nil {
		return nil, fmt.Errorf("failed to decode nodeOperatingSystems: %s", err)
	}

	nodeOSes = append(nodeOSes, list.Items...)

	return nodeOSes, nil
}
