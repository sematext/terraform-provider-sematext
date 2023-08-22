package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppMongodb tests resource creation.
func TestAccResourceLifecycleAppMongodb(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_mongodb", "MongoDB")
}


// TestAccResourceUpdateAppMongodb tests for resource updates.
func TestAccResourceUpdateAppMongodb(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_mongodb", "MongoDB")
}
