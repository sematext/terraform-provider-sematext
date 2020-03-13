package sematext

import (
	"strings"
)

// testAccSematextMonitorAkkaBasic tests resource creation.
func testAccSematextMonitorAkkaBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Akka"))
}

// testAccSematextMonitorAkkaUpdate tests for resource updates.
func testAccSematextMonitorAkkaUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Akka"))
}
