package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorAwsec2Basic tests resource creation.
func TestAccSematextMonitorAwsec2Basic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_awsec2")
}

// TestAccSematextMonitorAwsec2Update tests for resource updates.
func TestAccSematextMonitorAwsec2Update(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_awsec2")
}
