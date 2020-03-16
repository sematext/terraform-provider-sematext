package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorHbaseBasic tests resource creation.
func TestAccSematextMonitorHbaseBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Hbase"))
}

// TestAccSematextMonitorHbaseUpdate tests for resource updates.
func TestAccSematextMonitorHbaseUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Hbase"))
}
