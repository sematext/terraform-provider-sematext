package kafka

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
	"github.com/sematext/terraform-provider-sematext/internal/common"
)

// TestAccResourceAppKafka tests resource lifecycle.
func TestAccAppResourceKafka(t *testing.T) {

	appType := "Kafka"

	switch appType {

	case "AWS EBS", "AWS EC2", "AWS ELB":

		common.TestAccResourceAWS(t, "sematext_app_kafka", "Kafka")

	default:

		common.TestAccResourceDefault(t, "sematext_app_kafka", "Kafka")
		
	}

}


