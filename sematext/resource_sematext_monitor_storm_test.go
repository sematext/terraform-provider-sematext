package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorStormBasic tests resource creation.
func TestAccSematextMonitorStormBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_storm")
}

// TestAccSematextMonitorStormUpdate tests for resource updates.
func TestAccSematextMonitorStormUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_storm")
}
