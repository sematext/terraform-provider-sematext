package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppNodejs tests resource creation.
func TestAccResourceLifecycleAppNodejs(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_nodejs", "Node.js")
}


// TestAccResourceUpdateAppNodejs tests for resource updates.
func TestAccResourceUpdateAppNodejs(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_nodejs", "Node.js")
}
