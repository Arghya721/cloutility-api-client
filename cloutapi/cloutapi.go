package cloutapi

/* Copyright 2022-2023 (C) Blue Safespring AB
   Programmed by Jan Johansson
   Contributions by Daniel Oquiñena and Patrik Lundin
   All rights reserved for now, will have liberal
   license later */

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Client struct {
	httpClient   *http.Client
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Expires      int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func RunClient() {
	var (
		c          Client
		user       me
		myConsumer consumer
		myNode     node
	)

	// Initialize http.Client
	c.httpClient = &http.Client{
		Timeout: time.Second * 10,
	}

	// Get Auth token by passing username, password and client_id from config file
	err := c.authenticate(
		viper.GetString("client_id"),
		viper.GetString("username"),
		viper.GetString("password"),
	)
	if err != nil {
		log.Printf("Error authenticating: %s", err)
	}

	if viper.GetBool("debug") {
		log.Println("Token type:", c.TokenType)
		log.Println("Expires:", c.Expires)
		log.Println("Refresh token:", c.RefreshToken)
		log.Println("Access token:", c.AccessToken)
	}

	user, err = c.getUser()
	if err != nil {
		log.Println("Error retrieving userdata: ", err)
	}

	if viper.GetBool("debug") {
		log.Println(user.Name)
		log.Println(user.BusinessUnit.Name)
		log.Println(user.BusinessUnit.ID)
	}

	node, err := c.getNode()
	if err != nil {
		log.Println("Error retrieving nodedata: ", err)
	}
	log.Println(node)

	if viper.GetBool("dry-run") {
		log.Println("Running in dry-run mode, exiting")
		os.Exit(0)
	}

	c.createConsumer(user.BusinessUnit.ID)
	log.Println("Created a Consumer")
	c.createNode(user.BusinessUnit.ID, myConsumer.ID)
	log.Println("Created a Node")
	log.Println(myNode)
}

func (c *Client) getUser() (me, error) {

	var result me

	body, err := c.makeRequest("/v1/me", "GET", nil)
	if err != nil {
		return me{}, err
	}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return me{}, err
	}

	return result, nil

}

func (c *Client) getNode() (string, error) {
	node, err := c.makeRequest("/v1/bunits/17/consumers/31/node", "GET", nil)
	if err != nil {
		return "", err
	}
	return node, nil
	// XXX needs conf or code to use your bUnit/node instead
}

func (c *Client) createConsumer(myid int) (consumer, error) {
	var cons string
	var newConsumer consumer

	createcons := "/v1/bunits/" + strconv.Itoa(myid) + "/consumers"
	name := map[string]string{
		"name": viper.GetString("node_name"),
	}
	jsonBody, err := json.Marshal(name)
	if err != nil {
		return consumer{}, err
	}

	cons, err = c.makeRequest(createcons, http.MethodPost, jsonBody)
	if err != nil {
		return consumer{}, err
	}
	if err := json.Unmarshal([]byte(cons), &newConsumer); err != nil {
		log.Fatal(err)
	}

	return newConsumer, nil
}

func (c *Client) createNode(myid int, myConsumer int) (node, error) {

	var newNode node
	var nodestr string
	createstr := "/v1/bunits/" + strconv.Itoa(myid) + "/consumers/" +
		strconv.Itoa(myConsumer) + "/node"
	data := map[string]interface{}{
		"operatingSystem": map[string]string{
			"name": "Linux",
		},
		"type": map[string]string{
			"name": "File server",
		},
		"server": map[string]string{
			"name": "tsm12.backup.sto2.safedc.net",
		},
		"clientOptionSet": map[string]string{
			"name": "STANDARD",
		},
		"contact":  "Someone",
		"cpuCount": 1,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return node{}, err
	}

	nodestr, err = c.makeRequest(createstr, http.MethodPost, payload)
	if err != nil {
		return node{}, err
	}
	if err := json.Unmarshal([]byte(nodestr), &newNode); err != nil {
		return node{}, err
	}

	return newNode, nil

}

func (c *Client) makeRequest(contextPath string, method string, payload []byte) (string, error) {

	var reader io.Reader
	if payload != nil {
		reader = bytes.NewReader(payload)
	}

	req, err := http.NewRequest(method, viper.GetString("url")+contextPath, reader)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "safespring-golang-client")
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Origin", viper.GetString("client_origin"))
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	// XXX - needs conf file

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (c *Client) authenticate(client_id, username, password string) error {

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

	req, err := http.NewRequest(http.MethodPost, viper.GetString("url")+authurl,
		strings.NewReader(loginData.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "safespring-golang-client")
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Origin", viper.GetString("client_origin"))
	// XXX - needs conf file

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(body), &c); err != nil {
		return err
	}

	return nil

}
