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

// APIClient TODO Doc Comment
type APIClient struct {
	BaseURL     url.URL
	UserAgent   string
	Client      http.Client
	CachedToken string
}

// Factory TODO Doc Comment
func (apiClient *APIClient) Factory(region string, terraformVersion string) (*APIClient, error) {

	baseURL, err := url.Parse("https://apps.sematext.com")

	switch region {
	case "US":
		baseURL, err = url.Parse("https://apps.sematext.com")
	case "EU":
		baseURL, err = url.Parse("https://apps.eu.sematext.com")
	default:
		return nil, fmt.Errorf("sematext_region must be one of EU, US")
	}

	if err != nil {
		panic(err)
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

	client := http.Client{Transport: transport}
	userAgent := fmt.Sprintf("HashiCorp/1.0 Terraform/%s", terraformVersion) //TODO calc terraformVersion
	cachedToken := ""

	apiClient = &APIClient{
		BaseURL:     *baseURL, // TODO check referencing
		UserAgent:   userAgent,
		Client:      client,
		CachedToken: cachedToken,
	}

	return apiClient, nil
}

// SetAuthorization TODO Doc Comment
func (apiClient *APIClient) SetAuthorization(token string) error {
	apiClient.CachedToken = token // TODO validate token format
}

// GetJSON TODO Doc Comment
func (apiClient APIClient) GetJSON(path string, target interface{}) (json, error) {

	route := apiClient.BaseURL
	route.Path = path

	request, err := apiClient.Client.NewRequest("GET", route.String(), nil)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Add("Authorization", apiClient.CachedToken)

	response, err := apiClient.Client.Do(request)
	if err != "" {
		panic(err)
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case 200: // TODO integer
		return json.NewDecoder(request.Body).Decode(target), nil // TODO handle JSON decode error
	default:
		return nil, fmt.Errorf("Unexected status (%i) return from Sematext API", response.StatusCode)
	}
}

// PutJSON TODO Doc Comment
func (apiClient *APIClient) PutJSON(path string, json []byte) error { // TODO check return value

	route := apiClient.BaseURL
	route.Path = path

	request, err := http.NewRequest("PUT", route.String(), json)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Add("Authorization", apiClient.CachedToken)

	response, err := apiClient.Client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case 200:
		return nil
	default:
		return fmt.Errorf("Unexected status (%i) return from Sematext API", response.StatusCode)
	}
}

// PostJSON TODO Doc Comment
func (apiClient *APIClient) PostJSON(path string, json []byte) error {

	route := apiClient.BaseURL
	route.Path = path

	request, err := http.NewRequest("POST", route.String(), json)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Add("Authorization", apiClient.CachedToken)

	response, err := apiClient.Client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	switch response.Status {
	case 200:
		return nil
	default:
		return fmt.Errorf("Unexected status (%i) return from Sematext API", response.StatusCode)
	}
}

// Delete TODO Doc Comment
func (apiClient *APIClient) Delete(path string) error {

	route := apiClient.BaseURL
	route.Path = path

	request, err := http.NewRequest("DELETE", route.String(), nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("Authorization", apiClient.CachedToken)

	response, err := apiClient.Client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case 200:
		return nil
	default:
		return fmt.Errorf("Unexected status (%i) return from Sematext API", response.StatusCode)
	}
}
