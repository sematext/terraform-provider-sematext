package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorKafkaBasic tests resource creation.
func TestAccSematextMonitorKafkaBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_kafka", "Kafka")
}

// TestAccSematextMonitorKafkaUpdate tests for resource updates.
func TestAccSematextMonitorKafkaUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_kafka", "Kafka")
}
