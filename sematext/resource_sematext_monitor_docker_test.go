package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorDockerBasic tests resource creation.
func TestAccSematextMonitorDockerBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Docker"))
}

// TestAccSematextMonitorDockerUpdate tests for resource updates.
func TestAccSematextMonitorDockerUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Docker"))
}
