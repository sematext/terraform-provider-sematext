package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorCassandraBasic tests resource creation.
func TestAccSematextMonitorCassandraBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_cassandra", "Cassandra")
}

// TestAccSematextMonitorCassandraUpdate tests for resource updates.
func TestAccSematextMonitorCassandraUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_cassandra", "Cassandra")
}
