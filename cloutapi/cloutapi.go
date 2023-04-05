package cloutapi

/* Copyright 2022-2023 (C) Blue Safespring AB
   Programmed by Jan Johansson
   Contributions by Daniel Oquiñena and Patrik Lundin
   All rights reserved for now, will have liberal
   license later */

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type AuthenticatedClient struct {
	HttpClient   *http.Client
	BaseURL      string
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Expires      int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// Temporary function for testing purposes
func RunClient() {
	var (
		user       me
		myConsumer consumer
		myNode     node
	)

	// Initialize client by passing username, password and client_id from config file
	c, err := Init(
		viper.GetString("client_id"),
		viper.GetString("client_origin"),
		viper.GetString("username"),
		viper.GetString("password"),
		viper.GetString("url"),
	)
	if err != nil {
		log.Printf("Error authenticating: %s", err)
		os.Exit(1)
	}

	if viper.GetBool("debug") {
		log.Println("Token type:", c.TokenType)
		log.Println("Expires:", c.Expires)
		log.Println("Refresh token:", c.RefreshToken)
		log.Println("Access token:", c.AccessToken)
	}

	user, err = c.GetUser()
	if err != nil {
		log.Println("Error retrieving userdata: ", err)
	}

	if viper.GetBool("debug") {
		log.Println(user.Name)
		log.Println(user.BusinessUnit.Name)
		log.Println(user.BusinessUnit.ID)
	}

	node, err := c.GetNode()
	if err != nil {
		log.Println("Error retrieving nodedata: ", err)
	}
	log.Println(node)

	if viper.GetBool("dry-run") {
		log.Println("Running in dry-run mode, exiting")
		os.Exit(0)
	}

	c.CreateConsumer(user.BusinessUnit.ID, viper.GetString("node_name"))
	log.Println("Created a Consumer")
	c.CreateNode(user.BusinessUnit.ID, myConsumer.ID)
	log.Println("Created a Node")
	log.Println(myNode)
}

func (c *AuthenticatedClient) apiRequest(contextPath string, method string, payload []byte) (string, error) {

	var reader io.Reader
	if payload != nil {
		reader = bytes.NewReader(payload)
	}

	req, err := http.NewRequest(method, c.BaseURL+contextPath, reader)
	if err != nil {
		return "", fmt.Errorf("failed to complete request: %s", err)
	}

	req.Header.Set("User-Agent", "safespring-golang-client")
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Origin", viper.GetString("client_origin"))
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	// XXX - needs conf file

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve response body: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %s", err)
	}

	return string(body), nil
}

// Initialize client and return an AuthenticatedClient
func Init(client_id, origin, username, password, baseURL string) (*AuthenticatedClient, error) {

	var c AuthenticatedClient

	// Initialize http.Client
	c.HttpClient = &http.Client{
		Timeout: time.Second * 10,
	}

	c.BaseURL = baseURL

	authurl := "/v1/oauth"

	loginData := url.Values{}
	loginData.Add("client_id", client_id)
	loginData.Add("grant_type", "password")
	loginData.Add("username", username)
	loginData.Add("password", password)

	if viper.GetBool("debug") {
		log.Println("data:\n", loginData)
		log.Println("enpoint:", viper.GetString("url")+authurl)
	}

	req, err := http.NewRequest(http.MethodPost, c.BaseURL+authurl, strings.NewReader(loginData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

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
		return &AuthenticatedClient{}, fmt.Errorf("failed to decode authentication response: %s", err)
	}

	return &c, nil

}
