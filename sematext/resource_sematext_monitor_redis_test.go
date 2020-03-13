package sematext

import (
	"strings"
)

// testAccSematextMonitorRedisBasic tests resource creation.
func testAccSematextMonitorRedisBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Redis"))
}

// testAccSematextMonitorRedisUpdate tests for resource updates.
func testAccSematextMonitorRedisUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Redis"))
}
