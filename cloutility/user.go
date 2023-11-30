package cloutility

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	UserBUnit UserBUnit `json:"businessUnit"`
	Href      string    `json:"href"`
	Name      string    `json:"name"`
	ID        int       `json:"id"`
	Locked    bool      `json:"locked"`
}

type UserBUnit struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
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
