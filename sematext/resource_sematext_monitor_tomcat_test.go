package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorTomcatBasic tests resource creation.
func TestAccSematextMonitorTomcatBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_tomcat", "Tomcat")
}

// TestAccSematextMonitorTomcatUpdate tests for resource updates.
func TestAccSematextMonitorTomcatUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_tomcat", "Tomcat")
}
