package sematext

import (
	"strings"
)

// testAccSematextMonitorSparkBasic tests resource creation.
func testAccSematextMonitorSparkBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Spark"))
}

// testAccSematextMonitorSparkUpdate tests for resource updates.
func testAccSematextMonitorSparkUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Spark"))
}
