package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor_test.go.template
	Then run generate/generate.sh
*/

import (
	"testing"
)

// TestAccSematextMonitorHadoopmrv1Basic tests resource creation.
func TestAccSematextMonitorHadoopmrv1Basic(t *testing.T) {
	CommonMonitorBasicTest(t, "sematext_monitor_hadoopmrv1", "Hadoop-MRv1")
}

// TestAccSematextMonitorHadoopmrv1Update tests for resource updates.
func TestAccSematextMonitorHadoopmrv1Update(t *testing.T) {
	CommonMonitorUpdateTest(t, "sematext_monitor_hadoopmrv1", "Hadoop-MRv1")
}
