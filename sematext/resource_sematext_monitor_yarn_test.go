package sematext

import (
	"strings"
)

// testAccSematextMonitorYarnBasic tests resource creation.
func testAccSematextMonitorYarnBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Yarn"))
}

// testAccSematextMonitorYarnUpdate tests for resource updates.
func testAccSematextMonitorYarnUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Yarn"))
}
