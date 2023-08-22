package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppJvm tests resource creation.
func TestAccResourceLifecycleAppJvm(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_jvm", "JVM")
}


// TestAccResourceUpdateAppJvm tests for resource updates.
func TestAccResourceUpdateAppJvm(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_jvm", "JVM")
}
