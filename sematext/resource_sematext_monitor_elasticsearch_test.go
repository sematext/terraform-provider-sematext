package sematext

import (
	"strings"
)

// testAccSematextMonitorElasticsearchBasic tests resource creation.
func testAccSematextMonitorElasticsearchBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Elasticsearch"))
}

// testAccSematextMonitorElasticsearchUpdate tests for resource updates.
func testAccSematextMonitorElasticsearchUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Elasticsearch"))
}
