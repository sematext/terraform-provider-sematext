package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorPostgresqlBasic tests resource creation.
func TestAccSematextMonitorPostgresqlBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_postgresql", "Postgresql")
}

// TestAccSematextMonitorPostgresqlUpdate tests for resource updates.
func TestAccSematextMonitorPostgresqlUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_postgresql", "Postgresql")
}
