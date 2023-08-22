package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppRedis tests resource creation.
func TestAccResourceLifecycleAppRedis(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_redis", "Redis")
}


// TestAccResourceUpdateAppRedis tests for resource updates.
func TestAccResourceUpdateAppRedis(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_redis", "Redis")
}
