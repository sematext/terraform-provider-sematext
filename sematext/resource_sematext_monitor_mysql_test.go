package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorMysqlBasic tests resource creation.
func TestAccSematextMonitorMysqlBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_mysql", "MySQL")
}

// TestAccSematextMonitorMysqlUpdate tests for resource updates.
func TestAccSematextMonitorMysqlUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_mysql", "MySQL")
}
