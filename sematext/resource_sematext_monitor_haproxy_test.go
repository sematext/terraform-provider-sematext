package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorHaproxyBasic tests resource creation.
func TestAccSematextMonitorHaproxyBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Haproxy"))
}

// TestAccSematextMonitorHaproxyUpdate tests for resource updates.
func TestAccSematextMonitorHaproxyUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Haproxy"))
}
