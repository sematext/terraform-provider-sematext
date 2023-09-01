package solrcloud_test

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
	"github.com/sematext/terraform-provider-sematext/internal/testcase"
)

// TestAccResource tests resource lifecycle.
func TestAccResource(t *testing.T) {

	appType := "Solrcloud"

	switch appType {

	case "AWS EBS", "AWS EC2", "AWS ELB":

		testcase.TestAccResourceAWS(t, "sematext_app_solrcloud", "SolrCloud")

	default:

		testcase.TestAccResourceDefault(t, "sematext_app_solrcloud", "SolrCloud")
		
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

