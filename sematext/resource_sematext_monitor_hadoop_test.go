package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorHadoopBasic tests resource creation.
func TestAccSematextMonitorHadoopBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Hadoop"))
}

// TestAccSematextMonitorHadoopUpdate tests for resource updates.
func TestAccSematextMonitorHadoopUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Hadoop"))
}
