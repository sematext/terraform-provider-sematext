package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"github.com/sematext/terraform-provider-sematext/sematext"
	"testing"
)

// TestAccResourceLifecycleAppStorm tests resource creation.
func TestAccResourceLifecycleAppStorm(t *testing.T) {
	sematext.ResourceTestLifecycleApp(t, "sematext_app_storm", "Storm")
}


// TestAccResourceUpdateAppStorm tests for resource updates.
func TestAccResourceUpdateAppStorm(t *testing.T) {
	sematext.ResourceTestUpdateApp(t, "sematext_app_storm", "Storm")
}
