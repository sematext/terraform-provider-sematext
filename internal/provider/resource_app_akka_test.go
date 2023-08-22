package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppAkka tests resource creation.
func TestAccResourceLifecycleAppAkka(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_akka", "Akka")
}


// TestAccResourceUpdateAppAkka tests for resource updates.
func TestAccResourceUpdateAppAkka(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_akka", "Akka")
}
