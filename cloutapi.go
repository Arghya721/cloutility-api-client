package main

/* Copyright 2022-2023 (C) Blue Safespring AB
   Programmed by Jan Johansson
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
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func main() {

	var myauth auth
	var user me
	var myConsumer consumer
	var myNode node

	viper.SetConfigName("config")
	viper.SetConfigType("properties")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config file: %w", err))
	}

	myauth = doLogin(
		viper.GetString("client_id"),
		viper.GetString("username"),
		viper.GetString("password"),
	)

	if viper.GetBool("debug") {
		fmt.Println("Token type:", myauth.TokenType)
		fmt.Println("Expires:", myauth.Expires)
		fmt.Println("Refresh token:", myauth.RefreshToken)
		fmt.Println("Access token:", myauth.AccessToken)
	}

	user = getUser(myauth)

	if viper.GetBool("debug") {
		fmt.Println(user.Name)
		fmt.Println(user.BusinessUnit.Name)
		fmt.Println(user.BusinessUnit.ID)
	}

	fmt.Println(getNode(myauth))

	if viper.GetBool("create") {
		myConsumer = createConsumer(myauth, user.BusinessUnit.ID)
		fmt.Println("Created a Consumer")
		myNode = createNode(myauth, user.BusinessUnit.ID, myConsumer.ID)
		fmt.Println("Created a Node")
		fmt.Println(myNode)
	}
}

func getUser(myauth auth) me {

	var result me

	will_print := 0

	body := getRequest(myauth, "/v1/me", will_print)

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatal(err)
	}

	return result

}

func getNode(myauth auth) string {
	return getRequest(myauth, "/v1/bunits/17/consumers/31/node", 0)
	// XXX needs conf or code to use your bUnit/node instead
}

func createConsumer(myauth auth, myid int) consumer {
	var cons string
	var newConsumer consumer

	createcons := "/v1/bunits/" + strconv.Itoa(myid) + "/consumers"
	name := map[string]string{
		"name": viper.GetString("client_name"),
	}
	jsonBody, err := json.Marshal(name)
	if err != nil {
		fmt.Printf("Could not marshal data: %s", err)
	}
	// jsonBody := []byte(`{"name": "test-host-name-goes-here"}`)
	// XXX name should come from input  ^^

	cons = postRequest(myauth, createcons, jsonBody)
	if err := json.Unmarshal([]byte(cons), &newConsumer); err != nil {
		log.Fatal(err)
	}

	return newConsumer
}

func createNode(myauth auth, myid int, myConsumer int) node {

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
	jsonBody, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error marshaling data %s", err)
	}

	// jsonBody := []byte(
	// 	`{
	//        "operatingSystem": {
	//          "name": "Linux",
	//        },
	//        "type": {
	//          "name": "File server",
	//        },
	//        "server": {
	//           "name": "tsm12.backup.sto2.safedc.net",
	//        },
	//        "clientOptionSet": {
	//          "name": "STANDARD",
	//        },
	//      "contact": "Someone",
	//      "cpuCount": 1
	//      }`)
	// XXX hardcoded platform, needs conf

	nodestr = postRequest(myauth, createstr, jsonBody)
	if err := json.Unmarshal([]byte(nodestr), &newNode); err != nil {
		log.Fatal(err)
	}

	return newNode

}

func postRequest(myauth auth, posturl string, jsonBody []byte) string {

	postClient := http.Client{
		Timeout: time.Second * 10,
	}

	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, viper.GetString("url")+posturl, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "safespring-golang-client")
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Origin", viper.GetString("client_origin"))
	req.Header.Set("Authorization", "Bearer "+myauth.AccessToken)
	// XXX - needs conf file

	resp, getErr := postClient.Do(req)
	if getErr != nil {
		fmt.Printf("HTTP: %s\n", resp.Status)
		log.Fatal(getErr)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if viper.GetBool("debug") {
		fmt.Println("Body: ")
		fmt.Println(resp.Body)
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Printf("HTTP: %s\n", resp.Status)
		log.Fatal(readErr)
	}

	return string(body)
}

func getRequest(myauth auth, geturl string, print int) string {

	getClient := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, viper.GetString("url")+geturl, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "safespring-golang-client")
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Origin", viper.GetString("client_origin"))
	req.Header.Set("Authorization", "Bearer "+myauth.AccessToken)
	// XXX - needs conf file

	resp, getErr := getClient.Do(req)
	if getErr != nil {
		fmt.Printf("HTTP: %s\n", resp.Status)
		log.Fatal(getErr)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Printf("HTTP: %s\n", resp.Status)
		log.Fatal(readErr)
	}

	if print > 0 {
		fmt.Println(string(body))
	}

	return string(body)
}

func doLogin(client_id, username, password string) auth {

	authurl := "/v1/oauth"

	loginClient := http.Client{
		Timeout: time.Second * 10,
	}

	loginData := url.Values{}
	loginData.Add("client_id", client_id)
	loginData.Add("grant_type", "password")
	loginData.Add("username", username)
	loginData.Add("password", password)
	// XXX - needs conf file

	if viper.GetBool("debug") {
		fmt.Println("data:\n", loginData)
		fmt.Println("enpoint:", viper.GetString("url")+authurl)
	}

	req, err := http.NewRequest(http.MethodPost, viper.GetString("url")+authurl,
		strings.NewReader(loginData.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "safespring-golang-client")
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Origin", viper.GetString("client_origin"))
	// XXX - needs conf file

	res, getErr := loginClient.Do(req)
	if getErr != nil {
		fmt.Printf("HTTP: %s\n", res.Status)
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Printf("HTTP: %s\n", res.Status)
		log.Fatal(readErr)
	}

	// fmt.Println("Body1: ", string(body))

	//	var result map[string]interface{}
	var result auth
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		log.Fatal(err)
	}

	return result
}
