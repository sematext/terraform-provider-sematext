package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorMemcachedBasic tests resource creation.
func TestAccSematextMonitorMemcachedBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_memcached")
}

// TestAccSematextMonitorMemcachedUpdate tests for resource updates.
func TestAccSematextMonitorMemcachedUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_memcached")
}
