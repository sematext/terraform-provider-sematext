package sematext

import (
	"strings"
)

// testAccSematextMonitorSolrcloudBasic tests resource creation.
func testAccSematextMonitorSolrcloudBasic(t *testing.T) {
	testAccSematextMonitorBasic(t, strings.ToLower("sematext_monitor_Solrcloud"))
}

// testAccSematextMonitorSolrcloudUpdate tests for resource updates.
func testAccSematextMonitorSolrcloudUpdate(t *testing.T) {
	testAccSematextMonitorUpdate(t, strings.ToLower("sematext_monitor_Solrcloud"))
}
