package sematext

import (
	"strings"
)

// testAccSematextMonitorClickhouseBasic tests resource creation.
func testAccSematextMonitorClickhouseBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Clickhouse"))
}

// testAccSematextMonitorClickhouseUpdate tests for resource updates.
func testAccSematextMonitorClickhouseUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Clickhouse"))
}
