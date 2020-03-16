package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorClickhouseBasic tests resource creation.
func TestAccSematextMonitorClickhouseBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Clickhouse"))
}

// TestAccSematextMonitorClickhouseUpdate tests for resource updates.
func TestAccSematextMonitorClickhouseUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Clickhouse"))
}
