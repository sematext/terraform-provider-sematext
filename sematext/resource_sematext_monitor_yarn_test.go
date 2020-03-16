package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorYarnBasic tests resource creation.
func TestAccSematextMonitorYarnBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Yarn"))
}

// TestAccSematextMonitorYarnUpdate tests for resource updates.
func TestAccSematextMonitorYarnUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Yarn"))
}
