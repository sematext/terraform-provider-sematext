package sematext

// TODO - The API Key needs to be passed as Header parameter with name Authorization and value should be in the format apiKey <Value>. e.g. apiKey e5f18450-205a-48eb-8589-7d49edaea813

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"runtime"
	"time"
)

// Config is the structure that stores the configuration to talk to a sematext endpoint
type Config struct {
	Domain string
	Client http.Client
	UserAgent string
}


// defaultPooledTransport returns a new http.Transport with similar default
// values to http.DefaultTransport.
func defaultPooledTransport() *http.Transport {
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
	}
	return transport
}


// defaultTransport returns a new http.Transport with similar default values to
// http.DefaultTransport, but with idle connections and keepalives disabled.
func defaultTransport() *http.Transport {
	transport := defaultPooledTransport()
	transport.DisableKeepAlives = true
	transport.MaxIdleConnsPerHost = -1
	return transport
}


// buildHTTPSClient sets up SSL
func buildHTTPSClient() (*http.Client, error) {
	tlsConfig := &tls.Config{}
	transport := defaultTransport()
	transport.TLSClientConfig = tlsConfig
	return &http.Client{Transport: transport}, nil
}


// NewClient returns a new configured HTTPS client.
func (c *Config) NewClient() (*http.Client, error) {
	return buildHTTPSClient()
}


// NewConfig returns a new Config from a supplied ResourceData.
func NewConfig(d *schema.ResourceData) (*Config, error) {

	region := d.Get("sematext_region").(string)

	if region == "US" {
		domain := "apps.sematext.com".(string)
	} else if region == "EU" {
		domain := "apps.eu.sematext.com".(string)
	} else {
		return nil, fmt.Errorf("sematext_region must be one of EU, US")
	}

	config := &Config {
		AcceptableTerraformVersion: "12" //TODO version acceptance
		Schema : "https"
		Domain: domain
		Client: config.NewClient()
		UserAgent fmt.Sprintf("HashiCorp/1.0 Terraform/%s", terraformVersion)
	}

	return config, nil
}
