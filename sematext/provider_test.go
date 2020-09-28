package sematext

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider schema.Provider
var testAccProviderFunc func() *schema.Provider

func init() {
	testAccProvider = *Provider()
	testAccProviders = map[string]*schema.Provider{
		"sematext": &testAccProvider,
	}
	testAccProviderFunc = func() *schema.Provider { return &testAccProvider }
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ schema.Provider = *Provider()
}

func TestProviderConfigure(t *testing.T) {

	p := Provider()

	config := map[string]interface{}{
		"sematext_region": "US",
	}
	err := p.Configure(nil, terraform.NewResourceConfigRaw(config))
	if err != nil {
		t.Fatal(err)
	}
}

func testAccPreCheck(t *testing.T) {

	token := os.Getenv("SEMATEXT_API_KEY")
	if !IsValidUUID(token) {
		t.Fatal("ERROR : SEMATEXT_API_KEY environment not set correctly")
		return
	}

	region := os.Getenv("SEMATEXT_REGION")
	if !IsValidSematextRegion(region) {
		t.Fatal("ERROR : SEMATEXT_REGION environment not set correctly")
	}

	if os.Getenv("AWS_ACCESS_KEY_ID") == "" {
		t.Fatal("ERROR : AWS_ACCESS_KEY_ID must be set for acceptance tests")
	}

	if os.Getenv("AWS_ACCESS_KEY_ID") != "" && os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
		t.Fatal("ERROR : AWS_SECRET_ACCESS_KEY must be set for acceptance tests")
	}

	if os.Getenv("AWS_REGION") == "" {
		t.Fatal("ERROR : AWS_REGION must be set for acceptance tests")
	}

	return
}
