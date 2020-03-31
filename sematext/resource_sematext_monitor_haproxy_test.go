package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorHaproxyBasic tests resource creation.
func TestAccSematextMonitorHaproxyBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_haproxy", "HAProxy")
}

// TestAccSematextMonitorHaproxyUpdate tests for resource updates.
func TestAccSematextMonitorHaproxyUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_haproxy", "HAProxy")
}
