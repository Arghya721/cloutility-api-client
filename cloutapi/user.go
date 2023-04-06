package cloutapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *AuthenticatedClient) GetUser() (me, error) {
	var user me
	endpoint := c.BaseURL + "/v1/me"

	body, err := c.apiRequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return me{}, fmt.Errorf("error requesting userdata: %s", err)
	}

	if err := json.Unmarshal([]byte(body), &user); err != nil {
		return me{}, fmt.Errorf("error decoding userdata: %s", err)
	}

	return user, nil
}
