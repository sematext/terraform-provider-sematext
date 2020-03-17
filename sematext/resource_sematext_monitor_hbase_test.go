package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorHbaseBasic tests resource creation.
func TestAccSematextMonitorHbaseBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_hbase")
}

// TestAccSematextMonitorHbaseUpdate tests for resource updates.
func TestAccSematextMonitorHbaseUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_hbase")
}
