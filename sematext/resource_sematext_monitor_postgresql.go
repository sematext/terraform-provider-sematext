package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorPostgresql is the resource class that handles sematext_monitor_postgresql
func resourceSematextMonitorPostgresql() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Postgresql")

	return &schema.Resource{
		Create: resourceMonitorCreatePostgresql,
		Read:   resourceMonitorReadPostgresql,
		Update: resourceMonitorUpdatePostgresql,
		Delete: resourceMonitorDeletePostgresql,
		Exists: resourceMonitorExistsPostgresql,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreatePostgresql creates the sematext_monitor_postgresql resource.
func resourceMonitorCreatePostgresql(d *schema.ResourceData, meta interface{}) error {
	apptype := "Postgresql"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadPostgresql reads the sematext_monitor_postgresql resource from Sematext Cloud.
func resourceMonitorReadPostgresql(d *schema.ResourceData, meta interface{}) error {
	apptype := "Postgresql"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdatePostgresql updates Sematext Cloud from the sematext_monitor_postgresql resource.
func resourceMonitorUpdatePostgresql(d *schema.ResourceData, meta interface{}) error {
	apptype := "Postgresql"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeletePostgresql marks a sematext_monitor_postgresql resource as retired.
func resourceMonitorDeletePostgresql(d *schema.ResourceData, meta interface{}) error {
	apptype := "Postgresql"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsPostgresql checks a sematext_monitor_postgresql resource exists in Sematext Cloud.
func resourceMonitorExistsPostgresql(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Postgresql"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportPostgresql checks a sematext_monitor_postgresql resource exists in Sematext Cloud.
func resourceSematextMonitorImportPostgresql(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Postgresql"
	return CommonMonitorImport(d, meta, apptype)
}

*/
