package sematext

import (
	"strings"
)

// testAccSematextMonitorAWSBasic tests resource creation.
func testAccSematextMonitorAWSBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_AWS"))
}

// testAccSematextMonitorAWSUpdate tests for resource updates.
func testAccSematextMonitorAWSUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_AWS"))
}
