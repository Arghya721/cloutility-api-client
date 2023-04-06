package cloutapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type NodeOperatingSystem struct {
	Href   string                 `json:"href"`
	Total  int                    `json:"total"`
	Offset int                    `json:"offset"`
	First  string                 `json:"first"`
	Items  []NodeOperatingSystems `json:"items"`
}

type NodeOperatingSystems struct {
	Href               string               `json:"href"`
	Name               string               `json:"name"`
	ShortName          string               `json:"shortName"`
	SupportedNodeTypes []SupportedNodeTypes `json:"supportedNodeTypes"`
	ID                 int                  `json:"id"`
	CreatedDate        time.Time            `json:"createdDate"`
}

type SupportedNodeTypes struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CreatedDate time.Time `json:"createdDate"`
	ShortName   string    `json:"shortName"`
}

func (c *AuthenticatedClient) GetNodeOperatingSystem() ([]NodeOperatingSystems, error) {
	var osResponse NodeOperatingSystem
	var osList []NodeOperatingSystems

	endpoint := c.BaseURL + "/v1/nodeoperatingsystems"

	resp, err := c.apiRequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, fmt.Errorf("error requesting nodeOperatingSystems: %s", err)
	}
	if err := json.Unmarshal([]byte(resp), &osResponse); err != nil {
		return nil, fmt.Errorf("failed to decode nodeOperatingSystems: %s", err)
	}

	osList = append(osList, osResponse.Items...)

	return osList, nil
}
