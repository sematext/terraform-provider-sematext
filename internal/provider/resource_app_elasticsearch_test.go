package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppElasticsearch tests resource creation.
func TestAccResourceLifecycleAppElasticsearch(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_elasticsearch", "Elastic Search")
}


// TestAccResourceUpdateAppElasticsearch tests for resource updates.
func TestAccResourceUpdateAppElasticsearch(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_elasticsearch", "Elastic Search")
}
