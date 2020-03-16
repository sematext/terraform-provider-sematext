package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorZookeeperBasic tests resource creation.
func TestAccSematextMonitorZookeeperBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Zookeeper"))
}

// TestAccSematextMonitorZookeeperUpdate tests for resource updates.
func TestAccSematextMonitorZookeeperUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Zookeeper"))
}
