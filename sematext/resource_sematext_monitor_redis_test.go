package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorRedisBasic tests resource creation.
func TestAccSematextMonitorRedisBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Redis"))
}

// TestAccSematextMonitorRedisUpdate tests for resource updates.
func TestAccSematextMonitorRedisUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Redis"))
}
