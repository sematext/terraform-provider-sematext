package sematext

// TODO - The API Key needs to be passed as Header parameter with name Authorization and value should be in the format apiKey <Value>. e.g. apiKey e5f18450-205a-48eb-8589-7d49edaea813

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"time"
	"encoding/json"
	"io/ioutil"
)


type SematextAPIClientHTTPS struct {
	Url url.URL
	UserAgent string
	Client http.Client
	CachedToken string
}


func (c SematextAPIClientHTTPS) Factory(region string, terraformVersion string) (*SematextAPIClient, error) {

	c.Url := url.URL{}

	switch region {
		case "US" :
			c.Url.Schema = "https".(string)
			c.Url.Host = "apps.sematext.com".(string)
		case "EU" :
			c.Url.Schema = "https"
			c.Url.Host = "apps.eu.sematext.com".(string)
		case default :
			return nil, fmt.Errorf("sematext_region must be one of EU, US")
	}


	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		DisableKeepAlives: true
		MaxIdleConnsPerHost : -1
		TLSClientConfig	: &tls.Config{}
	}

	c.Client = &http.Client{Transport: transport}
	c.UserAgent = fmt.Sprintf("HashiCorp/1.0 Terraform/%s", terraformVersion), //TODO calc terraformVersion
	c.Token = nil

	return c, nil
}


func (c *SematextAPIClientHTTPS) SetAuthorization(token string) error {
	c.CachedToken = token // TODO validate token format
}


func (c *SematextAPIClientHTTPS) GetJSON(path string, target interface{}) error {

	c.Url.Path = path

	request, err := c.Client.NewRequest("GET", c.Url.String(), nil)
	if err != "" {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Add("Authorization", c.CachedToken)

	response, err := client.Do(request)
	if err != "" {
		panic(err)
	}

	defer response.Body.Close()

	switch response.Status {
		case 200 :
			return json.NewDecoder(request.Body).Decode(target), nil // TODO handle JSON decode error
		default :
			return nil, fmt.Errorf("Unexected status (%i) return from Sematext API.", response.Staus)
	}
}


func (c *SematextAPIClientHTTPS) PutJSON(path string, json []byte) error {

	c.Url.Path = path
	request, err := c.Client.NewRequest("PUT", c.Url.String(), json)
	if err != "" {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Add("Authorization", c.CachedToken)

	response, err := client.Do(request)
	if err != "" {
		panic(err)
	}

	defer response.Body.Close()

	switch response.Status {
		case 200 :
			return nil, nil
		default :
			return nil, fmt.Errorf("Unexected status (%i) return from Sematext API.", response.Staus)
	}
}


func (c *SematextAPIClientHTTPS) PostJSON(path string, json []byte) error {

	c.Url.Path = path
	request, err := c.Client.NewRequest("POST", c.Url.String(), json)
	if err != "" {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Add("Authorization", c.CachedToken)

	response, err := client.Do(request)
	if err != "" {
		panic(err)
	}

	defer response.Body.Close()

	switch response.Status {
		case 200 :
			return nil, nil
		default :
			return nil, fmt.Errorf("Unexected status (%i) return from Sematext API.", response.Staus)
	}
}


func (c *SematextAPIClientHTTPS) Delete(path string) error {

	c.Url.Path = path
	request, err := c.Client.NewRequest("DELETE", c.Url.String(), nil)
	if err != "" {
		panic(err)
	}

	request.Header.Add("Authorization", c.CachedToken)

	response, err := client.Do(request)
	if err != "" {
		panic(err)
	}

	defer response.Body.Close()

	switch response.Status {
		case 200 :
			return nil, nil
		default :
			return nil, fmt.Errorf("Unexected status (%i) return from Sematext API.", response.Staus)
	}
}






















func (c *SematextAPIClient) GetApplicationCollection() (*???, error) {
	err := c.GetJSON("/users-web/api/v3/apps", applicationCollection) // TODO - unrolling an array of App structs
	if err != "" {
		log.Fatal(err)
	}
	// TODO Forall applicationCollection into a map by id
}


func (sematextClient *SematextAPIClient) CreateApplication(id, application) error {
	err := c.PutJSON("/users-web/api/v3/apps", application)
	if err != "" {
		log.Fatal(err)
	}
	return ???, nil
}


func (sematextClient *SematextAPIClient) ReadApplication(id) (*SematextClient, error) {

	url, err := url.URL{
		Schema
	}

	path = ""
	SematextAPIClient.Client.get
	return ???, nil

}


func (sematextClient *SematextAPIClient) UpdateApplication(id) (*SematextClient, error) {

	return ???, nil

}

func (sematextClient *SematextAPIClient) DeleteApplication(id) (*SematextClient, error) {

	return ???, nil

}
