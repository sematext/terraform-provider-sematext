package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppZookeeper tests resource creation.
func TestAccResourceLifecycleAppZookeeper(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_zookeeper", "ZooKeeper")
}


// TestAccResourceUpdateAppZookeeper tests for resource updates.
func TestAccResourceUpdateAppZookeeper(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_zookeeper", "ZooKeeper")
}
