package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorApacheBasic tests resource creation.
func TestAccSematextMonitorApacheBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Apache"))
}

// TestAccSematextMonitorApacheUpdate tests for resource updates.
func TestAccSematextMonitorApacheUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Apache"))
}
