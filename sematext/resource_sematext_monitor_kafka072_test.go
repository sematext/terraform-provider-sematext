package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorKafka072Basic tests resource creation.
func TestAccSematextMonitorKafka072Basic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_kafka072")
}

// TestAccSematextMonitorKafka072Update tests for resource updates.
func TestAccSematextMonitorKafka072Update(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_kafka072")
}
