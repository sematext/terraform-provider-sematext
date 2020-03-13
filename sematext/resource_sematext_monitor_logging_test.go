package sematext

import (
	"strings"
)

// testAccSematextMonitorLoggingBasic tests resource creation.
func testAccSematextMonitorLoggingBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Logging"))
}

// testAccSematextMonitorLoggingUpdate tests for resource updates.
func testAccSematextMonitorLoggingUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Logging"))
}
