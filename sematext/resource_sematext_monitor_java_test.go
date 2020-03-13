package sematext

import (
	"strings"
)

// testAccSematextMonitorJavaBasic tests resource creation.
func testAccSematextMonitorJavaBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Java"))
}

// testAccSematextMonitorJavaUpdate tests for resource updates.
func testAccSematextMonitorJavaUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Java"))
}
