package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorNodejsBasic tests resource creation.
func TestAccSematextMonitorNodejsBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Nodejs"))
}

// TestAccSematextMonitorNodejsUpdate tests for resource updates.
func TestAccSematextMonitorNodejsUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Nodejs"))
}
