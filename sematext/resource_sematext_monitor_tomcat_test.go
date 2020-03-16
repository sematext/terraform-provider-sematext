package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorTomcatBasic tests resource creation.
func TestAccSematextMonitorTomcatBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Tomcat"))
}

// TestAccSematextMonitorTomcatUpdate tests for resource updates.
func TestAccSematextMonitorTomcatUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Tomcat"))
}
