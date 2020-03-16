package sematext

import (
	"strings"
	"testing"
)

// TestAccSematextMonitorElasticsearchBasic tests resource creation.
func TestAccSematextMonitorElasticsearchBasic(t *testing.T) {
	CommonMonitorBasicTest(t, strings.ToLower("sematext_monitor_Elasticsearch"))
}

// TestAccSematextMonitorElasticsearchUpdate tests for resource updates.
func TestAccSematextMonitorElasticsearchUpdate(t *testing.T) {
	CommonMonitorUpdateText(t, strings.ToLower("sematext_monitor_Elasticsearch"))
}
