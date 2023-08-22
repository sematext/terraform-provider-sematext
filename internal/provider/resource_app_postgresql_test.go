package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppPostgresql tests resource creation.
func TestAccResourceLifecycleAppPostgresql(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_postgresql", "postgresql")
}


// TestAccResourceUpdateAppPostgresql tests for resource updates.
func TestAccResourceUpdateAppPostgresql(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_postgresql", "postgresql")
}
