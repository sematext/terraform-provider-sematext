package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorNginxplusBasic tests resource creation.
func TestAccSematextMonitorNginxplusBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_nginxplus")
}

// TestAccSematextMonitorNginxplusUpdate tests for resource updates.
func TestAccSematextMonitorNginxplusUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_nginxplus")
}
