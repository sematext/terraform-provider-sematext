package sematext

import (
	"strings"
)

// testAccSematextMonitorTomcatBasic tests resource creation.
func testAccSematextMonitorTomcatBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Tomcat"))
}

// testAccSematextMonitorTomcatUpdate tests for resource updates.
func testAccSematextMonitorTomcatUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Tomcat"))
}
