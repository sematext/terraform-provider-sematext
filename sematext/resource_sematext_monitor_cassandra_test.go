package sematext

import (
	"strings"
)

// testAccSematextMonitorCassandraBasic tests resource creation.
func testAccSematextMonitorCassandraBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Cassandra"))
}

// testAccSematextMonitorCassandraUpdate tests for resource updates.
func testAccSematextMonitorCassandraUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Cassandra"))
}
