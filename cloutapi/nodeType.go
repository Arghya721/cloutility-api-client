package cloutapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type nodeTypes struct {
	Href   string     `json:"href"`
	Total  int        `json:"total"`
	Offset int        `json:"offset"`
	First  string     `json:"first"`
	Items  []NodeType `json:"items"`
}

type NodeType struct {
	Href        string    `json:"href"`
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	ShortName   string    `json:"shortName"`
	CreatedDate time.Time `json:"createdDate"`
}

func (c *AuthenticatedClient) GetNodeType() ([]NodeType, error) {
	var (
		list  nodeTypes
		nodes []NodeType
	)

	// validate the base url to create the endpoint
	path := "/v1/nodetypes"
	baseURL, err := url.Parse(c.BaseURL)

	if err != nil {
		return nil, fmt.Errorf("error parsing base url: %s", err)
	}

	baseURL.Path = path
	endpoint := baseURL.String()

	resp, err := c.apiRequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, fmt.Errorf("error requesting NodeTypes: %s", err)
	}
	if err := json.Unmarshal([]byte(resp), &list); err != nil {
		return nil, fmt.Errorf("failed to decode NodeTypes: %s", err)
	}

	nodes = append(nodes, list.Items...)

	return nodes, nil
}
