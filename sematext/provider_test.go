package sematext

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"sematext": testAccProvider,
	}
}

var testProviders = map[string]terraform.ResourceProvider{
	"sematext": Provider(),
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	err := testAccProvider.Configure(terraform.NewResourceConfigRaw(nil)) // TODO - what does this do?
	if err != nil {
		t.Fatal(err)
	}
}

// TODO provider fixture and check, pulling API token from environment
