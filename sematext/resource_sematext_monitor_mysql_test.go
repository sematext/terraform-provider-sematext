package sematext

import (
	"strings"
)

// testAccSematextMonitorMysqlBasic tests resource creation.
func testAccSematextMonitorMysqlBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Mysql"))
}

// testAccSematextMonitorMysqlUpdate tests for resource updates.
func testAccSematextMonitorMysqlUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Mysql"))
}
