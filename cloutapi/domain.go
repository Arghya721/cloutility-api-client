package cloutapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Domain struct {
	Href   string    `json:"href"`
	Total  int       `json:"total"`
	Offset int       `json:"offset"`
	First  string    `json:"first"`
	Items  []Domains `json:"items"`
}

type Domains struct {
	Href             string    `json:"href"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	BackupRetention  int       `json:"backupRetention"`
	ArchiveRetention int       `json:"archiveRetention"`
	MissingInTsm     bool      `json:"missingInTsm"`
	ID               int       `json:"id"`
	CreatedDate      time.Time `json:"createdDate"`
}

func (c *AuthenticatedClient) GetDomains(bUnitID int) ([]Domains, error) {
	var domain Domain
	var domains []Domains

	endpoint := c.BaseURL + "/v1/bunits/" + strconv.Itoa(bUnitID) + "/defaultserver/domains"

	resp, err := c.apiRequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, fmt.Errorf("error requesting domains: %s", err)
	}
	if err := json.Unmarshal([]byte(resp), &domain); err != nil {
		return nil, fmt.Errorf("failed to decode domaindata: %s", err)
	}

	domains = append(domains, domain.Items...)

	return domains, nil
}
