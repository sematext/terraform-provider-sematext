package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorNodejsBasic tests resource creation.
func TestAccSematextMonitorNodejsBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_nodejs", "Node.js")
}

// TestAccSematextMonitorNodejsUpdate tests for resource updates.
func TestAccSematextMonitorNodejsUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_nodejs", "Node.js")
}
