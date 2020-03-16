package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorSolrcloudBasic tests resource creation.
func TestAccSematextMonitorSolrcloudBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Solrcloud"))
}

// TestAccSematextMonitorSolrcloudUpdate tests for resource updates.
func TestAccSematextMonitorSolrcloudUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Solrcloud"))
}
