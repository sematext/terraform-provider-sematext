package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorSolrBasic tests resource creation.
func TestAccSematextMonitorSolrBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Solr"))
}

// TestAccSematextMonitorSolrUpdate tests for resource updates.
func TestAccSematextMonitorSolrUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Solr"))
}
