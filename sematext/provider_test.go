package sematext

import (
	"os"
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

func TestProviderConfigure(t *testing.T) {
	p := Provider().(*schema.Provider)
	config := map[string]interface{}{
		"sematext_region": "US",
	}
	err := p.Configure(terraform.NewResourceConfigRaw(config))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {

	token := os.Getenv("SEMATEXT_API_TOKEN")
	if !IsValidUUID(token) {
		t.Fatal("SEMATEXT_API_TOKEN environment not set correctly")
		return
	}

	region := os.Getenv("SEMATEXT_REGION")
	if !IsValidSematextRegion(region) {
		t.Fatal("SEMATEXT_REGION environment not set correctly")
	}

	if os.Getenv("AWS_ACCESS_KEY_ID") == "" {
		t.Fatal("AWS_ACCESS_KEY_ID must be set for acceptance tests")
	}

	if os.Getenv("AWS_ACCESS_KEY_ID") != "" && os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
		t.Fatal("AWS_SECRET_ACCESS_KEY must be set for acceptance tests")
	}

	if os.Getenv("AWS_REGION") == "" {
		t.Fatal("AWS_REGION must be set for acceptance tests")
	}

	return
}
