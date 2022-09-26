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

// TestAccResourceLifecycleAppLogsene tests resource creation.
func TestAccResourceLifecycleAppLogsene(t *testing.T) {
	sematext.ResourceTestLifecycleApp(t, "sematext_app_logsene", "Logsene")
}

// TestAccResourceUpdateAppLogsene tests for resource updates.
func TestAccResourceUpdateAppLogsene(t *testing.T) {
	sematext.ResourceTestUpdateApp(t, "sematext_app_logsene", "Logsene")
}
