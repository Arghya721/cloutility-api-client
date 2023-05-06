package cloutility

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Href      string    `json:"href"`
	ID        int       `json:"id"`
	Locked    bool      `json:"locked"`
	Name      string    `json:"name"`
	UserBUnit UserBUnit `json:"businessUnit"`
}

type UserBUnit struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *AuthenticatedClient) GetUser() (*User, error) {
	var user User

	endpoint := "/v1/me"
	body, err := c.apiRequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, fmt.Errorf("error requesting userdata: %s", err)
	}

	if err := json.Unmarshal([]byte(body), &user); err != nil {
		return nil, fmt.Errorf("error decoding userdata: %s", err)
	}

	return &user, nil
}
