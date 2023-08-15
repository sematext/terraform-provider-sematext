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

// TestAccResourceLifecycleAppAwselb tests resource creation.
func TestAccResourceLifecycleAppAwselb(t *testing.T) {
	sematext.ResourceTestLifecycleApp(t, "sematext_app_awselb", "AWS ELB")
}


// TestAccResourceUpdateAppAwselb tests for resource updates.
func TestAccResourceUpdateAppAwselb(t *testing.T) {
	sematext.ResourceTestUpdateApp(t, "sematext_app_awselb", "AWS ELB")
}
