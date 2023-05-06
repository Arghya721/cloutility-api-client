package cloutapi

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	endpoint := "/v1/bunits/" + fmt.Sprintf("%d", bUnitID) + "/consumers"

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

	endpoint := "/v1/bunits/" + fmt.Sprintf("%d", bUnitID) + "/consumers/" + fmt.Sprintf("%d", consumerID)
	_, err := c.apiRequest(endpoint, http.MethodDelete, nil)
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

	endpoint := "/v1/bunits/" + fmt.Sprintf("%d", bUnitID) + "/consumers"
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
