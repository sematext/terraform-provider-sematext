package sematext

import (
	"strings"
)

// testAccSematextMonitorHbaseBasic tests resource creation.
func testAccSematextMonitorHbaseBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Hbase"))
}

// testAccSematextMonitorHbaseUpdate tests for resource updates.
func testAccSematextMonitorHbaseUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Hbase"))
}
