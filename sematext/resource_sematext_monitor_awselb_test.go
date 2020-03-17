package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorAwselbBasic tests resource creation.
func TestAccSematextMonitorAwselbBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_awselb")
}

// TestAccSematextMonitorAwselbUpdate tests for resource updates.
func TestAccSematextMonitorAwselbUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_awselb")
}
