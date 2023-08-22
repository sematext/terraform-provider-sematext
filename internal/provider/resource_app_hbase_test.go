package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppHbase tests resource creation.
func TestAccResourceLifecycleAppHbase(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_hbase", "HBase")
}


// TestAccResourceUpdateAppHbase tests for resource updates.
func TestAccResourceUpdateAppHbase(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_hbase", "HBase")
}
