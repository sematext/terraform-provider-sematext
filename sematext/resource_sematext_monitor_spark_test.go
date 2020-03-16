package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorSparkBasic tests resource creation.
func TestAccSematextMonitorSparkBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Spark"))
}

// TestAccSematextMonitorSparkUpdate tests for resource updates.
func TestAccSematextMonitorSparkUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Spark"))
}
