package sematext

import (
	"strings"
)

// testAccSematextMonitorDockerBasic tests resource creation.
func testAccSematextMonitorDockerBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Docker"))
}

// testAccSematextMonitorDockerUpdate tests for resource updates.
func testAccSematextMonitorDockerUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Docker"))
}
