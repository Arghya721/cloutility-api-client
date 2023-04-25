package cloutapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
	var (
		list     nodeOperatingSystems
		nodeOSes []NodeOperatingSystem
	)

	// validate the base url to create the endpoint
	path := "/v1/nodeoperatingsystems"
	baseURL, err := url.Parse(c.BaseURL)

	if err != nil {
		return nil, fmt.Errorf("error parsing base url: %s", err)
	}

	baseURL.Path = path
	endpoint := baseURL.String()

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
