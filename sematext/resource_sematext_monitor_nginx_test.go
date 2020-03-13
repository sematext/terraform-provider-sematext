package sematext

import (
	"strings"
)

// testAccSematextMonitorNginxBasic tests resource creation.
func testAccSematextMonitorNginxBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Nginx"))
}

// testAccSematextMonitorNginxUpdate tests for resource updates.
func testAccSematextMonitorNginxUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Nginx"))
}
