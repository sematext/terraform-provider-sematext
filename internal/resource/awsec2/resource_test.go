package awsec2

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
	"github.com/sematext/terraform-provider-sematext/internal/common"
)

// TestAccResourceAppAwsec2 tests resource lifecycle.
func TestAccAppResourceAwsec2(t *testing.T) {

	appType := "Awsec2"

	switch appType {

	case "AWS EBS", "AWS EC2", "AWS ELB":

		common.TestAccResourceAWS(t, "sematext_app_awsec2", "AWS EC2")

	default:

		common.TestAccResourceDefault(t, "sematext_app_awsec2", "AWS EC2")
		
	}

}


