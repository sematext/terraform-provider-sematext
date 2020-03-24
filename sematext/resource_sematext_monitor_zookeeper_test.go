package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorZookeeperBasic tests resource creation.
func TestAccSematextMonitorZookeeperBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_zookeeper", "ZooKeeper")
}

// TestAccSematextMonitorZookeeperUpdate tests for resource updates.
func TestAccSematextMonitorZookeeperUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_zookeeper", "ZooKeeper")
}
