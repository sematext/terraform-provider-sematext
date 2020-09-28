package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorMysql is the resource class that handles sematext_monitor_mysql
func resourceSematextMonitorMysql() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("MySQL")

	return &schema.Resource{
		Create: resourceMonitorCreateMysql,
		Read:   resourceMonitorReadMysql,
		Update: resourceMonitorUpdateMysql,
		Delete: resourceMonitorDeleteMysql,
		Exists: resourceMonitorExistsMysql,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateMysql creates the sematext_monitor_mysql resource.
func resourceMonitorCreateMysql(d *schema.ResourceData, meta interface{}) error {
	apptype := "MySQL"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadMysql reads the sematext_monitor_mysql resource from Sematext Cloud.
func resourceMonitorReadMysql(d *schema.ResourceData, meta interface{}) error {
	apptype := "MySQL"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateMysql updates Sematext Cloud from the sematext_monitor_mysql resource.
func resourceMonitorUpdateMysql(d *schema.ResourceData, meta interface{}) error {
	apptype := "MySQL"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteMysql marks a sematext_monitor_mysql resource as retired.
func resourceMonitorDeleteMysql(d *schema.ResourceData, meta interface{}) error {
	apptype := "MySQL"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsMysql checks a sematext_monitor_mysql resource exists in Sematext Cloud.
func resourceMonitorExistsMysql(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "MySQL"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportMysql checks a sematext_monitor_mysql resource exists in Sematext Cloud.
func resourceSematextMonitorImportMysql(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "MySQL"
	return CommonMonitorImport(d, meta, apptype)
}

*/
