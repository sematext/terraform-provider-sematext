package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppSpark tests resource creation.
func TestAccResourceLifecycleAppSpark(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_spark", "Spark")
}


// TestAccResourceUpdateAppSpark tests for resource updates.
func TestAccResourceUpdateAppSpark(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_spark", "Spark")
}
