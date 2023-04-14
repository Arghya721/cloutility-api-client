// Business units make up the heirarchical center of data in Cloutility.
// Their goal is to represent companies, divisions or any other meaningful
// entity or logical group.

package cloutapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type BusinessUnit struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	GroupName      string         `json:"groupName"`
	ReportRemotely bool           `json:"reportRemotely"`
	BusinessUnits  []BusinessUnit `json:"businessUnits"`
	InvoiceDay     int            `json:"invoiceDay"`
	// Tags           []any          `json:"tags"`
}

func (c *AuthenticatedClient) CreateBusinessUnit() error {
	return nil
}

func (c *AuthenticatedClient) DeleteBusinessUnit() error {
	return nil
}

func (c *AuthenticatedClient) GetBusinessUnit(bunitID int) (BusinessUnit, error) {
	var bunit BusinessUnit

	endpoint := c.BaseURL + "/v1/bunits?bunitId=" + strconv.Itoa(bunitID)

	resp, err := c.apiRequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return BusinessUnit{}, fmt.Errorf("error retrieving business units: %s", err)
	}
	if err := json.Unmarshal([]byte(resp), &bunit); err != nil {
		return BusinessUnit{}, fmt.Errorf("error decoding business unit response: %s", err)
	}

	return bunit, nil
}
