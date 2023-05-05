package cloutapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type actions struct {
	Href   string   `json:"href"`
	Total  int      `json:"total"`
	Offset int      `json:"offset"`
	First  string   `json:"first"`
	Items  []Action `json:"items"`
}

type Action struct {
	Href        string    `json:"href"`
	Name        string    `json:"name"`
	ID          int       `json:"id"`
	CreatedDate time.Time `json:"createdDate"`
}

func (c *AuthenticatedClient) GetActions() ([]Action, error) {
	var (
		list   actions
		action []Action
	)

	path := "/v1/actions"
	baseURL, err := url.Parse(c.BaseURL)

	if err != nil {
		return nil, fmt.Errorf("error parsing base url: %s", err)
	}

	baseURL.Path = path
	endpoint := baseURL.String()

	resp, err := c.apiRequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, fmt.Errorf("error requesting actions: %s", err)
	}
	if err := json.Unmarshal([]byte(resp), &list); err != nil {
		return nil, fmt.Errorf("failed to decode actions: %s", err)
	}

	action = append(action, list.Items...)

	return action, nil
}

func (c *AuthenticatedClient) GetAction(actionID int) (Action, error) {
	var action Action

	path := "/v1/actions/" + fmt.Sprintf("%d", actionID)
	baseURL, err := url.Parse(c.BaseURL)

	if err != nil {
		return Action{}, fmt.Errorf("error parsing base url: %s", err)
	}

	baseURL.Path = path
	endpoint := baseURL.String()

	resp, err := c.apiRequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return Action{}, fmt.Errorf("error requesting action: %s", err)
	}
	if err := json.Unmarshal([]byte(resp), &action); err != nil {
		return Action{}, fmt.Errorf("failed to decode action: %s", err)
	}

	return action, nil
}
