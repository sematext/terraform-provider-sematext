package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppSolrcloud tests resource creation.
func TestAccResourceLifecycleAppSolrcloud(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_solrcloud", "SolrCloud")
}


// TestAccResourceUpdateAppSolrcloud tests for resource updates.
func TestAccResourceUpdateAppSolrcloud(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_solrcloud", "SolrCloud")
}
