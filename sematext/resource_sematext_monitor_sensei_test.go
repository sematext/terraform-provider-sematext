package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorSenseiBasic tests resource creation.
func TestAccSematextMonitorSenseiBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_sensei", "Sensei")
}

// TestAccSematextMonitorSenseiUpdate tests for resource updates.
func TestAccSematextMonitorSenseiUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_sensei", "Sensei")
}
