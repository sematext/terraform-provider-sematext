package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorAWSBasic tests resource creation.
func TestAccSematextMonitorAWSBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_AWS"))
}

// TestAccSematextMonitorAWSUpdate tests for resource updates.
func TestAccSematextMonitorAWSUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_AWS"))
}
