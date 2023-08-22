package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
)

// TestAccResourceLifecycleAppKafka tests resource creation.
func TestAccResourceLifecycleAppKafka(t *testing.T) {
	ResourceTestLifecycleApp(t, "sematext_app_kafka", "Kafka")
}


// TestAccResourceUpdateAppKafka tests for resource updates.
func TestAccResourceUpdateAppKafka(t *testing.T) {
	ResourceTestUpdateApp(t, "sematext_app_kafka", "Kafka")
}
