package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorNginxplusBasic tests resource creation.
func TestAccSematextMonitorNginxplusBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Nginxplus"))
}

// TestAccSematextMonitorNginxplusUpdate tests for resource updates.
func TestAccSematextMonitorNginxplusUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Nginxplus"))
}
