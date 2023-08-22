package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppAwsec2 tests resource creation.
func TestAccResourceLifecycleAppAwsec2(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_awsec2", "AWS EC2")
}


// TestAccResourceUpdateAppAwsec2 tests for resource updates.
func TestAccResourceUpdateAppAwsec2(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_awsec2", "AWS EC2")
}
