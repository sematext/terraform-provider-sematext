package sematext

import (
	"strings"
)

// testAccSematextMonitorNginxplusBasic tests resource creation.
func testAccSematextMonitorNginxplusBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Nginxplus"))
}

// testAccSematextMonitorNginxplusUpdate tests for resource updates.
func testAccSematextMonitorNginxplusUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Nginxplus"))
}
