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

// TestAccResourceAppInfra tests resource lifecycle.
func TestAccAppResourceInfra(t *testing.T) {

	appType := "Infra"

	switch appType {

	case "AWS EBS", "AWS EC2", "AWS ELB":

		common.TestAccResourceAppAWS(t, "sematext_app_infra", "Infra")

	default:

		common.TestAccResourceAppDefault(t, "sematext_app_infra", "Infra")
		
	}

}


