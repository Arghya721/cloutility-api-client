package cloutapi

import (
	"encoding/json"
	"fmt"
)

func (c *AuthenticatedClient) GetUser() (me, error) {

	var result me

	body, err := c.apiRequest("/v1/me", "GET", nil)
	if err != nil {
		return me{}, fmt.Errorf("error requesting userdata: %s", err)
	}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return me{}, fmt.Errorf("error unmarshalling userdata: %s", err)
	}

	return result, nil

}
