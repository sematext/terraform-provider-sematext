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

// TestAccResourceAppAwselb tests resource lifecycle.
func TestAccAppResourceAwselb(t *testing.T) {

	appType := "Awselb"

	switch appType {

	case "AWS EBS", "AWS EC2", "AWS ELB":

		common.TestAccResourceAppAWS(t, "sematext_app_awselb", "AWS ELB")

	default:

		common.TestAccResourceAppDefault(t, "sematext_app_awselb", "AWS ELB")
		
	}

}


