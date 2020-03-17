package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorSolrcloudBasic tests resource creation.
func TestAccSematextMonitorSolrcloudBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_solrcloud")
}

// TestAccSematextMonitorSolrcloudUpdate tests for resource updates.
func TestAccSematextMonitorSolrcloudUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_solrcloud")
}
