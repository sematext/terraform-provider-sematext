package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorNginxBasic tests resource creation.
func TestAccSematextMonitorNginxBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_nginx", "Nginx")
}

// TestAccSematextMonitorNginxUpdate tests for resource updates.
func TestAccSematextMonitorNginxUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_nginx", "Nginx")
}
