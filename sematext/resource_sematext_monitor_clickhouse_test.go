package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorClickhouseBasic tests resource creation.
func TestAccSematextMonitorClickhouseBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_clickhouse")
}

// TestAccSematextMonitorClickhouseUpdate tests for resource updates.
func TestAccSematextMonitorClickhouseUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_clickhouse")
}
