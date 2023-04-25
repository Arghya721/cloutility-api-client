package cloutapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type clientOptionSets struct {
	Href   string            `json:"href"`
	Total  int               `json:"total"`
	Offset int               `json:"offset"`
	First  string            `json:"first"`
	Items  []ClientOptionSet `json:"items"`
}

type ClientOptionSet struct {
	Href        string    `json:"href"`
	Name        string    `json:"name"`
	ID          int       `json:"id"`
	CreatedDate time.Time `json:"createdDate"`
}

func (c *AuthenticatedClient) GetClientOptionSet(bUnitID int) ([]ClientOptionSet, error) {
	var (
		list clientOptionSets
		sets []ClientOptionSet
	)

	// validate the base url to create the endpoint
	path := "/v1/bunits" + strconv.Itoa(bUnitID) + "/defaultserver/clientoptionsets"
	baseURL, err := url.Parse(c.BaseURL)

	if err != nil {
		return nil, fmt.Errorf("error parsing base url: %s", err)
	}

	baseURL.Path = path
	endpoint := baseURL.String()

	resp, err := c.apiRequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, fmt.Errorf("error requesting clientOptionSets: %s", err)
	}
	if err := json.Unmarshal([]byte(resp), &list); err != nil {
		return nil, fmt.Errorf("failed to decode clientOptionSets: %s", err)
	}

	sets = append(sets, list.Items...)

	return sets, nil
}
