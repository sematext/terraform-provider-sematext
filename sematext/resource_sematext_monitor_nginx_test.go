package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorNginxBasic tests resource creation.
func TestAccSematextMonitorNginxBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Nginx"))
}

// TestAccSematextMonitorNginxUpdate tests for resource updates.
func TestAccSematextMonitorNginxUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Nginx"))
}
