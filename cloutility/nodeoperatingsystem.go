package cloutility

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type nodeOperatingSystems struct {
	Href   string                `json:"href"`
	First  string                `json:"first"`
	Items  []NodeOperatingSystem `json:"items"`
	Total  int                   `json:"total"`
	Offset int                   `json:"offset"`
}

type NodeOperatingSystem struct {
	CreatedDate        time.Time            `json:"createdDate"`
	Href               string               `json:"href"`
	Name               string               `json:"name"`
	ShortName          string               `json:"shortName"`
	SupportedNodeTypes []supportedNodeTypes `json:"supportedNodeTypes"`
	ID                 int                  `json:"id"`
}

type supportedNodeTypes struct {
	CreatedDate time.Time `json:"createdDate"`
	Name        string    `json:"name"`
	ShortName   string    `json:"shortName"`
	ID          int       `json:"id"`
}

func (c *AuthenticatedClient) GetNodeOperatingSystem() ([]NodeOperatingSystem, error) {
	var (
		list     nodeOperatingSystems
		nodeOSes []NodeOperatingSystem
	)

	endpoint := "/v1/nodeoperatingsystems"
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
