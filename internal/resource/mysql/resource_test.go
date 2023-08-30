package mysql

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app_test.go.template
	Then run generate/generate.sh
*/


import (
	"testing"
	"github.com/sematext/terraform-provider-sematext/internal/common"
)

// TestAccResourceAppMysql tests resource lifecycle.
func TestAccAppResourceMysql(t *testing.T) {

	appType := "Mysql"

	switch appType {

	case "AWS EBS", "AWS EC2", "AWS ELB":

		common.TestAccResourceAWS(t, "sematext_app_mysql", "MySQL")

	default:

		common.TestAccResourceDefault(t, "sematext_app_mysql", "MySQL")
		
	}

}


