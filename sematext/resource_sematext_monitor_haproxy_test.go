package sematext

import (
	"strings"
)

// testAccSematextMonitorHaproxyBasic tests resource creation.
func testAccSematextMonitorHaproxyBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Haproxy"))
}

// testAccSematextMonitorHaproxyUpdate tests for resource updates.
func testAccSematextMonitorHaproxyUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Haproxy"))
}
