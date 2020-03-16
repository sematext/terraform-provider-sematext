package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorInfraBasic tests resource creation.
func TestAccSematextMonitorInfraBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Infra"))
}

// TestAccSematextMonitorInfraUpdate tests for resource updates.
func TestAccSematextMonitorInfraUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Infra"))
}
