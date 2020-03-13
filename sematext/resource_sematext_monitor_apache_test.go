package sematext

import (
	"strings"
)

// testAccSematextMonitorApacheBasic tests resource creation.
func testAccSematextMonitorApacheBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Apache"))
}

// testAccSematextMonitorApacheUpdate tests for resource updates.
func testAccSematextMonitorApacheUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Apache"))
}
