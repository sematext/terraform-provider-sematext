package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorLoggingBasic tests resource creation.
func TestAccSematextMonitorLoggingBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Logging"))
}

// TestAccSematextMonitorLoggingUpdate tests for resource updates.
func TestAccSematextMonitorLoggingUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Logging"))
}
