package cloutapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type domains struct {
	Href   string   `json:"href"`
	Total  int      `json:"total"`
	Offset int      `json:"offset"`
	First  string   `json:"first"`
	Items  []Domain `json:"items"`
}

type Domain struct {
	Href             string    `json:"href"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	BackupRetention  int       `json:"backupRetention"`
	ArchiveRetention int       `json:"archiveRetention"`
	MissingInTsm     bool      `json:"missingInTsm"`
	ID               int       `json:"id"`
	CreatedDate      time.Time `json:"createdDate"`
}

func (c *AuthenticatedClient) GetDomains(bUnitID int) ([]Domain, error) {
	var (
		list    domains
		domains []Domain
	)

	// validate the base url to create the endpoint
	path := "/v1/bunits/" + strconv.Itoa(bUnitID) + "/defaultserver/domains"
	baseURL, err := url.Parse(c.BaseURL)

	if err != nil {
		return nil, fmt.Errorf("error parsing base url: %s", err)
	}

	baseURL.Path = path
	endpoint := baseURL.String()

	resp, err := c.apiRequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, fmt.Errorf("error requesting domains: %s", err)
	}
	if err := json.Unmarshal([]byte(resp), &list); err != nil {
		return nil, fmt.Errorf("failed to decode domaindata: %s", err)
	}

	domains = append(domains, list.Items...)

	return domains, nil
}
