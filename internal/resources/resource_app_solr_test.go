package resources

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
	"github.com/sematext/terraform-provider-sematext/internal/common"
)

// TestAccResourceAppSolr tests resource lifecycle.
func TestAccAppResourceSolr(t *testing.T) {

	appType := "Solr"

	switch appType {

	case "AWS EBS", "AWS EC2", "AWS ELB":

		common.TestAccResourceAppAWS(t, "sematext_app_solr", "Solr")

	default:

		common.TestAccResourceAppDefault(t, "sematext_app_solr", "Solr")
		
	}

}


