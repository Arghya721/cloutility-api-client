package cloutility

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type actions struct {
	Href   string   `json:"href"`
	First  string   `json:"first"`
	Items  []Action `json:"items"`
	Total  int      `json:"total"`
	Offset int      `json:"offset"`
}

type Action struct {
	CreatedDate time.Time `json:"createdDate"`
	Href        string    `json:"href"`
	Name        string    `json:"name"`
	ID          int       `json:"id"`
}

func (c *AuthenticatedClient) GetActions() ([]Action, error) {
	var (
		list   actions
		action []Action
	)

	endpoint := "/v1/actions"
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

	endpoint := "/v1/actions/" + fmt.Sprintf("%d", actionID)
	resp, err := c.apiRequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return Action{}, fmt.Errorf("error requesting action: %s", err)
	}
	if err := json.Unmarshal([]byte(resp), &action); err != nil {
		return Action{}, fmt.Errorf("failed to decode action: %s", err)
	}

	return action, nil
}
