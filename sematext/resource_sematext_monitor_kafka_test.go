package sematext

import (
	"strings"
)

// testAccSematextMonitorKafkaBasic tests resource creation.
func testAccSematextMonitorKafkaBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Kafka"))
}

// testAccSematextMonitorKafkaUpdate tests for resource updates.
func testAccSematextMonitorKafkaUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Kafka"))
}
