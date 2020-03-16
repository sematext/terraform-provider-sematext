package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorAkkaBasic tests resource creation.
func TestAccSematextMonitorAkkaBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Akka"))
}

// TestAccSematextMonitorAkkaUpdate tests for resource updates.
func TestAccSematextMonitorAkkaUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Akka"))
}
