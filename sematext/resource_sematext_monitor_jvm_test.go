package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorJvmBasic tests resource creation.
func TestAccSematextMonitorJvmBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_jvm", "JVM")
}

// TestAccSematextMonitorJvmUpdate tests for resource updates.
func TestAccSematextMonitorJvmUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_jvm", "JVM")
}
