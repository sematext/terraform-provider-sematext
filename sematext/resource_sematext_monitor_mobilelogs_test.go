package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorMobilelogsBasic tests resource creation.
func TestAccSematextMonitorMobilelogsBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_mobilelogs", "mobile-logs")
}

// TestAccSematextMonitorMobilelogsUpdate tests for resource updates.
func TestAccSematextMonitorMobilelogsUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_mobilelogs", "mobile-logs")
}
