package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorLogseneBasic tests resource creation.
func TestAccSematextMonitorLogseneBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_logsene", "Logsene")
}

// TestAccSematextMonitorLogseneUpdate tests for resource updates.
func TestAccSematextMonitorLogseneUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_logsene", "Logsene")
}
