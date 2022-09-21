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

// TestAccResourceLifecycleAppKafka tests resource creation.
func TestAccResourceLifecycleAppKafka(t *testing.T) {
	sematext.ResourceTestLifecycleApp(t, "sematext_app_kafka", "Kafka")
}

// TestAccResourceUpdateAppKafka tests for resource updates.
func TestAccResourceUpdateAppKafka(t *testing.T) {
	sematext.ResourceTestUpdateApp(t, "sematext_app_kafka", "Kafka")
}
