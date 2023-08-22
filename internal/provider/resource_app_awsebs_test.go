package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppAwsebs tests resource creation.
func TestAccResourceLifecycleAppAwsebs(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_awsebs", "AWS EBS")
}


// TestAccResourceUpdateAppAwsebs tests for resource updates.
func TestAccResourceUpdateAppAwsebs(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_awsebs", "AWS EBS")
}
