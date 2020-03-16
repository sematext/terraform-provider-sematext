package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorStormBasic tests resource creation.
func TestAccSematextMonitorStormBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Storm"))
}

// TestAccSematextMonitorStormUpdate tests for resource updates.
func TestAccSematextMonitorStormUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Storm"))
}
