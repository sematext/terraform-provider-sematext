package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorCassandraBasic tests resource creation.
func TestAccSematextMonitorCassandraBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Cassandra"))
}

// TestAccSematextMonitorCassandraUpdate tests for resource updates.
func TestAccSematextMonitorCassandraUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Cassandra"))
}
