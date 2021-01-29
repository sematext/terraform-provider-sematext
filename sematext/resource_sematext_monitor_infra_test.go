package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorInfraBasic tests resource creation.
func TestAccSematextMonitorInfraBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_infra", "Infra")
}

// TestAccSematextMonitorInfraUpdate tests for resource updates.
func TestAccSematextMonitorInfraUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_infra", "Infra")
}
