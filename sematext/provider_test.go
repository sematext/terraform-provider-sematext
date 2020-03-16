package sematext

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider
var testAccProviderFunc func() *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"sematext": testAccProvider,
	}
	testAccProviderFunc = func() *schema.Provider { return testAccProvider }
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func TestProvider_Configure(t *testing.T) {
	provider := Provider().(*schema.Provider)

	raw := map[string]interface{}{
		"sematext_region": "US",
	}

	// TODO rest of config into client

	config := terraform.NewResourceConfigRaw(raw)

	err := provider.Configure(config)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	/*
		if v := os.Getenv("CONSUL_HTTP_ADDR"); v != "" {
			return
		}
		if v := os.Getenv("CONSUL_ADDRESS"); v != "" {
			return
		}
		t.Fatal("Either CONSUL_ADDRESS or CONSUL_HTTP_ADDR must be set for acceptance tests")
	*/
}
