package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorDockerBasic tests resource creation.
func TestAccSematextMonitorDockerBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_docker")
}

// TestAccSematextMonitorDockerUpdate tests for resource updates.
func TestAccSematextMonitorDockerUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_docker")
}
