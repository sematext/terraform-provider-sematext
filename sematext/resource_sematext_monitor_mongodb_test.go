package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorMongodbBasic tests resource creation.
func TestAccSematextMonitorMongodbBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Mongodb"))
}

// TestAccSematextMonitorMongodbUpdate tests for resource updates.
func TestAccSematextMonitorMongodbUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Mongodb"))
}
