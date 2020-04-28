package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorMongodbBasic tests resource creation.
func TestAccSematextMonitorMongodbBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_mongodb", "MongoDB")
}

// TestAccSematextMonitorMongodbUpdate tests for resource updates.
func TestAccSematextMonitorMongodbUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_mongodb", "MongoDB")
}
