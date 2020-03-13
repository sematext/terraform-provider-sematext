package sematext

import (
	"strings"
)

// testAccSematextMonitorStormBasic tests resource creation.
func testAccSematextMonitorStormBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Storm"))
}

// testAccSematextMonitorStormUpdate tests for resource updates.
func testAccSematextMonitorStormUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Storm"))
}
