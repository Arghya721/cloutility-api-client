package cloutility

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type nodeTypes struct {
	Href   string     `json:"href"`
	First  string     `json:"first"`
	Items  []NodeType `json:"items"`
	Total  int        `json:"total"`
	Offset int        `json:"offset"`
}

type NodeType struct {
	CreatedDate time.Time `json:"createdDate"`
	Href        string    `json:"href"`
	Name        string    `json:"name"`
	ShortName   string    `json:"shortName"`
	ID          int       `json:"id"`
}

func (c *AuthenticatedClient) GetNodeType() ([]NodeType, error) {
	var (
		list  nodeTypes
		nodes []NodeType
	)

	endpoint := "/v1/nodetypes"
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
