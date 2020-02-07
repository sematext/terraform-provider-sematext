package sematext

// TODO - The API Key needs to be passed as Header parameter with name Authorization and value should be in the format apiKey <Value>. e.g. apiKey e5f18450-205a-48eb-8589-7d49edaea813

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"time"
)

type SematextAPIClient struct {
	BaseUrl     url.URL
	UserAgent   string
	Client      http.Client
	CachedToken string
}

func (apiClient SematextAPIClient) Factory(region string, terraformVersion string) (*SematextAPIClient, error) {

	switch region {
	case "US":
		apiClient.BaseUrl, err := url.Parse("https://apps.sematext.com")
	case "EU":
		apiClient.BaseUrl, err := url.Parse("https://apps.eu.sematext.com")
	default:
		return nil, fmt.Errorf("sematext_region must be one of EU, US")
	}

	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           (&net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		DisableKeepAlives:     true,
		TLSClientConfig:       &tls.Config{},
	}

	apiClient.Client = &http.Client{Transport: transport}
	apiClient.UserAgent = fmt.Sprintf("HashiCorp/1.0 Terraform/%s", terraformVersion) //TODO calc terraformVersion
	apiClient.CachedToken = nil

	return apiClient, nil
}

func (apiClient SematextAPIClient) SetAuthorization(token string) error {
	apiClient.CachedToken = token // TODO validate token format
}

func (apiClient SematextAPIClient) GetJSON(path string, target interface{}) error {

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
	case 200:
		return json.NewDecoder(request.Body).Decode(target), nil // TODO handle JSON decode error
	default:
		return nil, fmt.Errorf("Unexected status (%i) return from Sematext API.", response.Staus)
	}
}

func (c *SematextAPIClient) PutJSON(path string, json []byte) error {

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
	case 200:
		return nil, nil
	default:
		return nil, fmt.Errorf("Unexected status (%i) return from Sematext API.", response.Staus)
	}
}

func (c *SematextAPIClient) PostJSON(path string, json []byte) error {

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
	case 200:
		return nil, nil
	default:
		return nil, fmt.Errorf("Unexected status (%i) return from Sematext API.", response.Staus)
	}
}

func (c *SematextAPIClient) Delete(path string) error {

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
	case 200:
		return nil, nil
	default:
		return nil, fmt.Errorf("Unexected status (%i) return from Sematext API.", response.Staus)
	}
}
