package sematext

import (
	"strings"
)

// testAccSematextMonitorZookeeperBasic tests resource creation.
func testAccSematextMonitorZookeeperBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Zookeeper"))
}

// testAccSematextMonitorZookeeperUpdate tests for resource updates.
func testAccSematextMonitorZookeeperUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Zookeeper"))
}
