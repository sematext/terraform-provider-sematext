package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorRedisBasic tests resource creation.
func TestAccSematextMonitorRedisBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_redis")
}

// TestAccSematextMonitorRedisUpdate tests for resource updates.
func TestAccSematextMonitorRedisUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_redis")
}
