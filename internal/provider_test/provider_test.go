package provider_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/sematext/sematext-api-client-go/stcloud"

	"github.com/sematext/terraform-provider-sematext/internal/provider"
	"github.com/sematext/terraform-provider-sematext/internal/util"
)

// / testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var ProviderProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"sematext_cloud": providerserver.NewProtocol6WithError(provider.New("test")()),
}

func ProviderPreCheckTestAWS(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.

	token := os.Getenv("SEMATEXT_API_KEY")
	if !util.IsValidUUID(token) {
		t.Fatal("ERROR : SEMATEXT_API_KEY environment not set correctly")
	}

	region := os.Getenv("SEMATEXT_REGION")
	if !stcloud.IsValidSematextRegion(region) {
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

}

func ProviderPreCheckTestDefault(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.

	token := os.Getenv("SEMATEXT_API_KEY")
	if !util.IsValidUUID(token) {
		t.Fatal("ERROR : SEMATEXT_API_KEY environment not set correctly")
	}

	region := os.Getenv("SEMATEXT_REGION")
	if !stcloud.IsValidSematextRegion(region) {
		t.Fatal("ERROR : SEMATEXT_REGION environment not set correctly")
	}

}

/*

	@TODO - check provider config

	p := schema.Provider{}

	config := map[string]interface{}{
		"sematext_region": "US",
	}
	err := p.Configure(context.Background(), terraform.NewResourceConfigRaw(config))
	if err != nil {
		t.Fatal(err)
	}
*/
