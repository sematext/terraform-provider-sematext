package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorJavaBasic tests resource creation.
func TestAccSematextMonitorJavaBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Java"))
}

// TestAccSematextMonitorJavaUpdate tests for resource updates.
func TestAccSematextMonitorJavaUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Java"))
}
