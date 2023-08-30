package mongodb

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
	"github.com/sematext/terraform-provider-sematext/internal/common"
)

// TestAccResourceAppMongodb tests resource lifecycle.
func TestAccAppResourceMongodb(t *testing.T) {

	appType := "Mongodb"

	switch appType {

	case "AWS EBS", "AWS EC2", "AWS ELB":

		common.TestAccResourceAWS(t, "sematext_app_mongodb", "MongoDB")

	default:

		common.TestAccResourceDefault(t, "sematext_app_mongodb", "MongoDB")
		
	}

}


