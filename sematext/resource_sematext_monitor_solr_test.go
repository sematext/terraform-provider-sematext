package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorSolrBasic tests resource creation.
func TestAccSematextMonitorSolrBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_solr")
}

// TestAccSematextMonitorSolrUpdate tests for resource updates.
func TestAccSematextMonitorSolrUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_solr")
}
