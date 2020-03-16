package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorMysqlBasic tests resource creation.
func TestAccSematextMonitorMysqlBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Mysql"))
}

// TestAccSematextMonitorMysqlUpdate tests for resource updates.
func TestAccSematextMonitorMysqlUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Mysql"))
}
