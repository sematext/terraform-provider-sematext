package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorSearchanalyticsBasic tests resource creation.
func TestAccSematextMonitorSearchanalyticsBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_searchanalytics")
}

// TestAccSematextMonitorSearchanalyticsUpdate tests for resource updates.
func TestAccSematextMonitorSearchanalyticsUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_searchanalytics")
}
