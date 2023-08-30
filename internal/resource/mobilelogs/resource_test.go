package mobilelogs

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
	"github.com/sematext/terraform-provider-sematext/internal/test"
)

// TestAccResource tests resource lifecycle.
func TestAccResource(t *testing.T) {

	appType := "Mobilelogs"

	switch appType {

	case "AWS EBS", "AWS EC2", "AWS ELB":

		test.TestAccResourceAWS(t, "sematext_app_mobilelogs", "mobile-logs")

	default:

		test.TestAccResourceDefault(t, "sematext_app_mobilelogs", "mobile-logs")
		
	}

}


