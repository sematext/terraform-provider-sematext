package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppStorm tests resource creation.
func TestAccResourceLifecycleAppStorm(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_storm", "Storm")
}


// TestAccResourceUpdateAppStorm tests for resource updates.
func TestAccResourceUpdateAppStorm(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_storm", "Storm")
}
