package sematext

// TODO - The API Key needs to be passed as Header parameter with name Authorization and value should be in the format apiKey <Value>. e.g. apiKey e5f18450-205a-48eb-8589-7d49edaea813

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"time"
)

// APIClient TODO Doc Comment
type APIClient struct {
	BaseURL     url.URL
	UserAgent   string
	httpClient  *http.Client
	CachedToken string
}

// Init TODO Doc Comment
func (apiClient *APIClient) Init(region string, terraformVersion string) error {

	baseURL, err := url.Parse("https://apps.sematext.com")
	if err != nil {
		return err
	}

	switch region {
	case "US":
		baseURL, err = url.Parse("https://apps.sematext.com")
	case "EU":
		baseURL, err = url.Parse("https://apps.eu.sematext.com")
	default:
		err = errors.New("sematext_region must be one of EU, US")
	}

	if err != nil {
		return err
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

	apiClient.BaseURL = *baseURL
	apiClient.UserAgent = fmt.Sprintf("HashiCorp/1.0 Terraform/%s", terraformVersion)
	apiClient.httpClient = &http.Client{Transport: transport}
	apiClient.CachedToken = ""

	return nil
}

// SetAuthorization TODO Doc Comment
func (apiClient *APIClient) SetAuthorization(token string) error {
	re := regexp.MustCompile(`[0-9]{8}-[0-9]{4}-[0-9]{4}-[0-9]{4}-[0-9]{8}`)
	if re.Match([]byte(token)) {
		apiClient.CachedToken = fmt.Sprintf("apiKey %s", token)
		return nil
	}
	apiClient.CachedToken = ""
	return errors.New("Bad or missing Token")
}

// GetJSON TODO Doc Comment
func (apiClient *APIClient) GetJSON(path string, object interface{}) (*GenericAPIResponse, error) {

	if apiClient.CachedToken == "" {
		panic("Code error : method called without setting token")
	}

	route := apiClient.BaseURL
	route.Path = path

	req, err := http.NewRequest("GET", route.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", apiClient.CachedToken)

	res, err := apiClient.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		genericAppResponse := &GenericAppResponse{}
		err = json.NewDecoder(res.Body).Decode(genericAppResponse)
		if err != nil {
			return nil, err

		}
		return genericAppResponse, nil
	}

	return nil, fmt.Errorf("Unexected status (%i) return from Sematext API", res.StatusCode)
}

// PutJSON TODO Doc Comment
func (apiClient *APIClient) PutJSON(path string, object interface{}) (*GenericAPIResponse, error) {

	if apiClient.CachedToken == "" {
		panic("Code error : method called without setting token")
	}

	route := apiClient.BaseURL
	route.Path = path

	jsn, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", route.String(), bytes.NewBuffer(jsn))
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", apiClient.CachedToken)

	res, err := apiClient.httpClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		genericAppResponse := new(GenericAppResponse)
		err = json.NewDecoder(res.Body).Decode(genericAppResponse)
		if err != nil {
			return nil, err

		}
		return genericAppResponse, nil
	}

	return nil, fmt.Errorf("Unexected status (%i) return from Sematext API", res.StatusCode)
}

// PostJSON TODO Doc Comment
func (apiClient *APIClient) PostJSON(path string, object interface{}) (GenericAppResponse, error) {

	if apiClient.CachedToken == "" {
		panic("Code error : method called without setting token")
	}

	route := apiClient.BaseURL
	route.Path = path

	jsn, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", route.String(), bytes.NewBuffer(jsn))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", apiClient.CachedToken)

	res, err := apiClient.httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		genericAppResponse := &GenericAppResponse{}
		err = json.NewDecoder(res.Body).Decode(genericAppResponse)
		if err != nil {
			return nil, err

		}
		return genericAppResponse, nil
	}

	return nil, fmt.Errorf("Unexected status (%i) return from Sematext API", res.StatusCode)
}

// Delete TODO Doc Comment
func (apiClient *APIClient) Delete(path string) error {

	if apiClient.CachedToken == "" {
		panic("Code error : method called without setting token")
	}

	route := apiClient.BaseURL
	route.Path = path

	req, err := http.NewRequest("DELETE", route.String(), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", apiClient.CachedToken)

	res, err := apiClient.httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return res.StatusCode, fmt.Errorf("Unexected status (%i) return from Sematext API on Delete", res.StatusCode)
	}
	return nil
	nil
}
