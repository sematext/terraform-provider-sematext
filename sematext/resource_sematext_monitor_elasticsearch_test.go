package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorElasticsearchBasic tests resource creation.
func TestAccSematextMonitorElasticsearchBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_elasticsearch", "Elastic Search")
}

// TestAccSematextMonitorElasticsearchUpdate tests for resource updates.
func TestAccSematextMonitorElasticsearchUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_elasticsearch", "Elastic Search")
}
