package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorAkkaBasic tests resource creation.
func TestAccSematextMonitorAkkaBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_akka")
}

// TestAccSematextMonitorAkkaUpdate tests for resource updates.
func TestAccSematextMonitorAkkaUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_akka")
}
