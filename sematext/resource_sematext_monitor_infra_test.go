package sematext

import (
	"strings"
)

// testAccSematextMonitorInfraBasic tests resource creation.
func testAccSematextMonitorInfraBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Infra"))
}

// testAccSematextMonitorInfraUpdate tests for resource updates.
func testAccSematextMonitorInfraUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Infra"))
}
