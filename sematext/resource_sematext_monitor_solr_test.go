package sematext

import (
	"strings"
)

// testAccSematextMonitorSolrBasic tests resource creation.
func testAccSematextMonitorSolrBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Solr"))
}

// testAccSematextMonitorSolrUpdate tests for resource updates.
func testAccSematextMonitorSolrUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Solr"))
}
