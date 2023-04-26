package cloutapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type consumers struct {
	Href     string     `json:"href"`
	Total    int        `json:"total"`
	Offset   int        `json:"offset"`
	First    string     `json:"first"`
	Consumer []Consumer `json:"items"`
}

type Consumer struct {
	CreatedDate time.Time `json:"createdDate"`
	Href        string    `json:"href"`
	ID          int       `json:"id"`
	Name        string    `json:"name"`
}

func (c *AuthenticatedClient) CreateConsumer(bUnitID int, nodename string) (Consumer, error) {
	var (
		cons     string
		consumer Consumer
	)

	// validate the base url to create the endpoint
	path := "/v1/bunits/" + strconv.Itoa(bUnitID) + "/consumers"
	baseURL, err := url.Parse(c.BaseURL)

	if err != nil {
		return Consumer{}, fmt.Errorf("error parsing base url: %s", err)
	}

	baseURL.Path = path
	endpoint := baseURL.String()

	name := map[string]string{
		"Name": nodename,
	}
	jsonBody, err := json.Marshal(name)
	if err != nil {
		return Consumer{}, fmt.Errorf("error encoding json payload: %s", err)
	}

	cons, err = c.apiRequest(endpoint, http.MethodPost, jsonBody)
	if err != nil {
		return Consumer{}, fmt.Errorf("error creating consumer: %s", err)
	}
	if err := json.Unmarshal([]byte(cons), &consumer); err != nil {
		return Consumer{}, fmt.Errorf("error decoding consumer response: %s", err)
	}

	return consumer, nil
}

func (c *AuthenticatedClient) DeleteConsumer(bUnitID, consumerID int) error {

	// validate the base url to create the endpoint
	path := "/v1/bunits/" + strconv.Itoa(bUnitID) + "/consumers/" + strconv.Itoa(consumerID)
	baseURL, err := url.Parse(c.BaseURL)

	if err != nil {
		return fmt.Errorf("error parsing base url: %s", err)
	}

	baseURL.Path = path
	endpoint := baseURL.String()

	_, err = c.apiRequest(endpoint, http.MethodDelete, nil)
	if err != nil {
		return fmt.Errorf("error deleting consumer: %s", err)
	}

	return nil
}

func (c *AuthenticatedClient) GetConsumers(bUnitID int) ([]Consumer, error) {
	var (
		list     consumers
		consumer []Consumer
	)

	// validate the base url to create the endpoint
	path := "/v1/bunits/" + strconv.Itoa(bUnitID) + "/consumers"
	baseURL, err := url.Parse(c.BaseURL)

	if err != nil {
		return nil, fmt.Errorf("error parsing base url: %s", err)
	}

	baseURL.Path = path
	endpoint := baseURL.String()

	resp, err := c.apiRequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return []Consumer{}, fmt.Errorf("error retrieving consumers: %s", err)
	}
	if err := json.Unmarshal([]byte(resp), &list); err != nil {
		return nil, fmt.Errorf("error decoding consumers response: %s", err)
	}

	consumer = append(consumer, list.Consumer...)

	return consumer, nil
}
