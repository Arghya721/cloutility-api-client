/*
 * Copyright (c) Blue Safespring AB - Jan Johansson <jj@safespring.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package cloutility

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type AuthenticatedClient struct {
	HttpClient   *http.Client
	BaseURL      string
	Origin       string
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Expires      int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type ErrorResponse struct {
	Status           int    `json:"status"`
	Code             string `json:"code"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage"`
}

// Initialize client and return an AuthenticatedClient
func Init(ctx context.Context, client_id, origin, username, password, baseURL string) (*AuthenticatedClient, error) {
	var c AuthenticatedClient

	// Initialize http.Client
	c.HttpClient = &http.Client{
		Timeout: time.Second * 10,
	}

	// Get baseurl from passed variable
	c.BaseURL = baseURL
	c.Origin = origin

	authurl := "/v1/oauth"
	requestURL, err := url.ParseRequestURI(c.BaseURL + authurl)
	if err != nil {
		return nil, fmt.Errorf("error parsing request url: %s", err)
	}

	// Construct body
	loginData := url.Values{}
	loginData.Add("client_id", client_id)
	loginData.Add("grant_type", "password")
	loginData.Add("username", username)
	loginData.Add("password", password)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL.String(), strings.NewReader(loginData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set header propertys
	req.Header.Set("User-Agent", "safespring-golang-client")
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", origin)

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to complete authentication request: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to retrieve authentication: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read authentication response: %s", err)
	}

	if err := json.Unmarshal([]byte(body), &c); err != nil {
		return nil, fmt.Errorf("failed to decode authentication response: %s", err)
	}

	return &c, nil
}

func (c *AuthenticatedClient) apiRequest(endpoint string, method string, payload []byte) (string, error) {
	ctx := context.Background()

	var reader io.Reader
	if payload != nil {
		reader = bytes.NewReader(payload)
	}

	requestURL, err := url.ParseRequestURI(c.BaseURL + endpoint)
	if err != nil {
		return "", fmt.Errorf("error parsing request url: %s", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, requestURL.String(), reader)
	if err != nil {
		return "", fmt.Errorf("failed to complete request: %s", err)
	}

	req.Header.Set("User-Agent", "safespring-golang-client")
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Origin", c.Origin)
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve response body: %s", err)
	}
	defer resp.Body.Close()

	// Check response code and return error if not 2xx
	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !statusOK {
		var reqErr ErrorResponse
		message, _ := io.ReadAll(resp.Body)
		if err := json.Unmarshal([]byte(message), &reqErr); err != nil {
			return "", fmt.Errorf("error response %v", resp.StatusCode)
		}
		return "", fmt.Errorf("%s", reqErr.DeveloperMessage)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %s", err)
	}

	return string(body), nil
}
