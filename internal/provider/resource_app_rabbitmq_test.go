package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppRabbitmq tests resource creation.
func TestAccResourceLifecycleAppRabbitmq(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_rabbitmq", "rabbitmq")
}


// TestAccResourceUpdateAppRabbitmq tests for resource updates.
func TestAccResourceUpdateAppRabbitmq(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_rabbitmq", "rabbitmq")
}
