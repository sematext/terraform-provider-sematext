package sematext

import (
	"strings"
)

// testAccSematextMonitorNodejsBasic tests resource creation.
func testAccSematextMonitorNodejsBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Nodejs"))
}

// testAccSematextMonitorNodejsUpdate tests for resource updates.
func testAccSematextMonitorNodejsUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Nodejs"))
}
