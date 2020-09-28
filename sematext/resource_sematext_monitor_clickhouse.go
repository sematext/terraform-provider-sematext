package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorClickhouse is the resource class that handles sematext_monitor_clickhouse
func resourceSematextMonitorClickhouse() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("ClickHouse")

	return &schema.Resource{
		Create: resourceMonitorCreateClickhouse,
		Read:   resourceMonitorReadClickhouse,
		Update: resourceMonitorUpdateClickhouse,
		Delete: resourceMonitorDeleteClickhouse,
		Exists: resourceMonitorExistsClickhouse,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateClickhouse creates the sematext_monitor_clickhouse resource.
func resourceMonitorCreateClickhouse(d *schema.ResourceData, meta interface{}) error {
	apptype := "ClickHouse"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadClickhouse reads the sematext_monitor_clickhouse resource from Sematext Cloud.
func resourceMonitorReadClickhouse(d *schema.ResourceData, meta interface{}) error {
	apptype := "ClickHouse"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateClickhouse updates Sematext Cloud from the sematext_monitor_clickhouse resource.
func resourceMonitorUpdateClickhouse(d *schema.ResourceData, meta interface{}) error {
	apptype := "ClickHouse"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteClickhouse marks a sematext_monitor_clickhouse resource as retired.
func resourceMonitorDeleteClickhouse(d *schema.ResourceData, meta interface{}) error {
	apptype := "ClickHouse"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsClickhouse checks a sematext_monitor_clickhouse resource exists in Sematext Cloud.
func resourceMonitorExistsClickhouse(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "ClickHouse"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportClickhouse checks a sematext_monitor_clickhouse resource exists in Sematext Cloud.
func resourceSematextMonitorImportClickhouse(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "ClickHouse"
	return CommonMonitorImport(d, meta, apptype)
}

*/
