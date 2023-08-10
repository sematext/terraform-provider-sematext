package generated

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"github.com/sematext/terraform-provider-sematext/sematext"
	"testing"
)

// TestAccResourceLifecycleAppCassandra tests resource creation.
func TestAccResourceLifecycleAppCassandra(t *testing.T) {
	sematext.ResourceTestLifecycleApp(t, "sematext_app_cassandra", "Cassandra")
}


// TestAccResourceUpdateAppCassandra tests for resource updates.
func TestAccResourceUpdateAppCassandra(t *testing.T) {
	sematext.ResourceTestUpdateApp(t, "sematext_app_cassandra", "Cassandra")
}
