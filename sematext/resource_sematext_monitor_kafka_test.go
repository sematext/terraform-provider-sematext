package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorKafkaBasic tests resource creation.
func TestAccSematextMonitorKafkaBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Kafka"))
}

// TestAccSematextMonitorKafkaUpdate tests for resource updates.
func TestAccSematextMonitorKafkaUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Kafka"))
}
