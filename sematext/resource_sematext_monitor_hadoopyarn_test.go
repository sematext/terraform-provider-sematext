package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorHadoopyarnBasic tests resource creation.
func TestAccSematextMonitorHadoopyarnBasic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_hadoopyarn")
}

// TestAccSematextMonitorHadoopyarnUpdate tests for resource updates.
func TestAccSematextMonitorHadoopyarnUpdate(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_hadoopyarn")
}
