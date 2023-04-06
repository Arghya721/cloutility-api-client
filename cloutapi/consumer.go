package cloutapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type consumer struct {
	AllowNoActivity bool `json:"allowNoActivity"`
	BusinessUnit    struct {
		Type      string `json:"$type"`
		Addresses struct {
			Type  string        `json:"$type"`
			Href  string        `json:"href"`
			Items []interface{} `json:"items"`
			Total int           `json:"total"`
		} `json:"addresses"`
		BillingStorageType   int `json:"billingStorageType"`
		BillingStorageTypeID int `json:"billingStorageTypeId"`
		BusinessUnits        struct {
			Type  string        `json:"$type"`
			Href  string        `json:"href"`
			Items []interface{} `json:"items"`
			Total int           `json:"total"`
		} `json:"businessUnits"`
		ClientOptionSetFilter            []interface{} `json:"clientOptionSetFilter"`
		Consumers                        []interface{} `json:"consumers"`
		CreatedDate                      time.Time     `json:"createdDate"`
		DomainFilter                     []interface{} `json:"domainFilter"`
		FinalDeleteRequestApprover       bool          `json:"finalDeleteRequestApprover"`
		Href                             string        `json:"href"`
		ID                               int           `json:"id"`
		InvoiceDay                       int           `json:"invoiceDay"`
		Name                             string        `json:"name"`
		NodeLimit                        int           `json:"nodeLimit"`
		PasswordExpirationDays           int           `json:"passwordExpirationDays"`
		RegistrationNumber               string        `json:"registrationNumber"`
		ReportRemotely                   bool          `json:"reportRemotely"`
		RequiredApproversOfDeleteRequest int           `json:"requiredApproversOfDeleteRequest"`
		StorageLimit                     int           `json:"storageLimit"`
		SupportResponsible               bool          `json:"supportResponsible"`
		Tags                             []interface{} `json:"tags"`
		TransferLimit                    int           `json:"transferLimit"`
		UseScheduleBindings              bool          `json:"useScheduleBindings"`
		Users                            struct {
			Type  string        `json:"$type"`
			Href  string        `json:"href"`
			Items []interface{} `json:"items"`
			Total int           `json:"total"`
		} `json:"users"`
		UsersCanApproveOwnRequests bool `json:"usersCanApproveOwnRequests"`
	} `json:"businessUnit"`
	CreatedDate                 time.Time `json:"createdDate"`
	DataSourceIsPotentialParent bool      `json:"dataSourceIsPotentialParent"`
	DataSourceState             struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"dataSourceState"`
	DataSourceType struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"dataSourceType"`
	Href string        `json:"href"`
	ID   int           `json:"id"`
	Name string        `json:"name"`
	Tags []interface{} `json:"tags"`
}

func (c *AuthenticatedClient) CreateConsumer(bUnitID int, nodename string) (consumer, error) {
	var cons string
	var newConsumer consumer

	endpoint := c.BaseURL + "/v1/bunits/" + strconv.Itoa(bUnitID) + "/consumers"
	name := map[string]string{
		"Name": nodename,
	}
	jsonBody, err := json.Marshal(name)
	if err != nil {
		return consumer{}, fmt.Errorf("error encoding json payload: %s", err)
	}

	cons, err = c.apiRequest(endpoint, http.MethodPost, jsonBody)
	if err != nil {
		return consumer{}, fmt.Errorf("error creating consumer: %s", err)
	}
	if err := json.Unmarshal([]byte(cons), &newConsumer); err != nil {
		return consumer{}, fmt.Errorf("error decoding consumer response: %s", err)
	}

	return newConsumer, nil
}

func (c *AuthenticatedClient) DeleteConsumer(bUnitID, consumerID int) (bool, error) {
	endpoint := c.BaseURL + "/v1/bunits/" + strconv.Itoa(bUnitID) + "/consumers/" + strconv.Itoa(consumerID)

	_, err := c.apiRequest(endpoint, http.MethodDelete, nil)
	if err != nil {
		return false, fmt.Errorf("error deleting consumer: %s", err)
	}

	return true, nil
}

func (c *AuthenticatedClient) GetConsumer(bUnitID int, nodename string) (consumer, error) {
	return consumer{}, nil
}
