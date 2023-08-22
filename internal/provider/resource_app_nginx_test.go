package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppNginx tests resource creation.
func TestAccResourceLifecycleAppNginx(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_nginx", "Nginx")
}


// TestAccResourceUpdateAppNginx tests for resource updates.
func TestAccResourceUpdateAppNginx(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_nginx", "Nginx")
}
