package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppSolr tests resource creation.
func TestAccResourceLifecycleAppSolr(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_solr", "Solr")
}


// TestAccResourceUpdateAppSolr tests for resource updates.
func TestAccResourceUpdateAppSolr(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_solr", "Solr")
}
