package generated

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/sematext/terraform-provider-sematext/sematext"
	"testing"
)

// TestAccResourceLifecycleAppNginxplus tests resource creation.
func TestAccResourceLifecycleAppNginxplus(t *testing.T) {
	sematext.ResourceTestLifecycleApp(t, "sematext_app_nginxplus", "Nginx-Plus")
}

// TestAccResourceUpdateAppNginxplus tests for resource updates.
func TestAccResourceUpdateAppNginxplus(t *testing.T) {
	sematext.ResourceTestUpdateApp(t, "sematext_app_nginxplus", "Nginx-Plus")
}
