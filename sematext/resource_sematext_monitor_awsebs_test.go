package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorAwsebsBasic tests resource creation.
func TestAccSematextMonitorAwsebsBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_awsebs")
}

// TestAccSematextMonitorAwsebsUpdate tests for resource updates.
func TestAccSematextMonitorAwsebsUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_awsebs")
}
