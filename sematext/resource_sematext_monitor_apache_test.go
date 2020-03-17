package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorApacheBasic tests resource creation.
func TestAccSematextMonitorApacheBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_apache")
}

// TestAccSematextMonitorApacheUpdate tests for resource updates.
func TestAccSematextMonitorApacheUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_apache")
}
