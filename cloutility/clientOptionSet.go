package cloutility

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type clientOptionSets struct {
	Href   string            `json:"href"`
	First  string            `json:"first"`
	Items  []ClientOptionSet `json:"items"`
	Total  int               `json:"total"`
	Offset int               `json:"offset"`
}

type ClientOptionSet struct {
	CreatedDate time.Time `json:"createdDate"`
	Href        string    `json:"href"`
	Name        string    `json:"name"`
	ID          int       `json:"id"`
}

func (c *AuthenticatedClient) GetClientOptionSet(bUnitID int) ([]ClientOptionSet, error) {
	var (
		list clientOptionSets
		sets []ClientOptionSet
	)

	// validate the base url to create the endpoint
	endpoint := "/v1/bunits/" + fmt.Sprintf("%d", bUnitID) + "/defaultserver/clientoptionsets"
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
