package sematext

import (
	"strings"
)

// testAccSematextMonitorMongodbBasic tests resource creation.
func testAccSematextMonitorMongodbBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Mongodb"))
}

// testAccSematextMonitorMongodbUpdate tests for resource updates.
func testAccSematextMonitorMongodbUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Mongodb"))
}
