package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorSparkBasic tests resource creation.
func TestAccSematextMonitorSparkBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_spark", "Spark")
}

// TestAccSematextMonitorSparkUpdate tests for resource updates.
func TestAccSematextMonitorSparkUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_spark", "Spark")
}
