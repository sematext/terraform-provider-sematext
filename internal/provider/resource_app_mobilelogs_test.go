package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppMobilelogs tests resource creation.
func TestAccResourceLifecycleAppMobilelogs(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_mobilelogs", "mobile-logs")
}


// TestAccResourceUpdateAppMobilelogs tests for resource updates.
func TestAccResourceUpdateAppMobilelogs(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_mobilelogs", "mobile-logs")
}
