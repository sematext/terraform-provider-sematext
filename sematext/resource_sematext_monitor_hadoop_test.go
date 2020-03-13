package sematext

import (
	"strings"
)

// testAccSematextMonitorHadoopBasic tests resource creation.
func testAccSematextMonitorHadoopBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Hadoop"))
}

// testAccSematextMonitorHadoopUpdate tests for resource updates.
func testAccSematextMonitorHadoopUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Hadoop"))
}
