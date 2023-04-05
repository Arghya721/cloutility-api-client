package cloutapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (c *AuthenticatedClient) CreateConsumer(myid int, nodename string) (consumer, error) {
	var cons string
	var newConsumer consumer

	createcons := "/v1/bunits/" + strconv.Itoa(myid) + "/consumers"
	name := map[string]string{
		"name": nodename,
	}
	jsonBody, err := json.Marshal(name)
	if err != nil {
		return consumer{}, fmt.Errorf("error encoding json payload: %s", err)
	}

	cons, err = c.apiRequest(createcons, http.MethodPost, jsonBody)
	if err != nil {
		return consumer{}, fmt.Errorf("error creating consumer: %s", err)
	}
	if err := json.Unmarshal([]byte(cons), &newConsumer); err != nil {
		return consumer{}, fmt.Errorf("error decoding consumer response: %s", err)
	}

	return newConsumer, nil
}

func (c *AuthenticatedClient) DeleteConsumer(myid int, nodename string) (consumer, error) {
	return consumer{}, nil
}

func (c *AuthenticatedClient) GetConsumer(myid int, nodename string) (consumer, error) {
	return consumer{}, nil
}
