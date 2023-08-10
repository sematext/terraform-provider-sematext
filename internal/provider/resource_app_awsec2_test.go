package generated

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"github.com/sematext/terraform-provider-sematext/sematext"
	"testing"
)

// TestAccResourceLifecycleAppAwsec2 tests resource creation.
func TestAccResourceLifecycleAppAwsec2(t *testing.T) {
	sematext.ResourceTestLifecycleApp(t, "sematext_app_awsec2", "AWS EC2")
}


// TestAccResourceUpdateAppAwsec2 tests for resource updates.
func TestAccResourceUpdateAppAwsec2(t *testing.T) {
	sematext.ResourceTestUpdateApp(t, "sematext_app_awsec2", "AWS EC2")
}
