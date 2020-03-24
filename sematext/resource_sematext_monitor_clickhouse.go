package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorClickhouse TODO Doc Comment
func resourceSematextMonitorClickhouse() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Clickhouse")

	return &schema.Resource{
		Create: resourceMonitorCreateClickhouse,
		Read:   resourceMonitorReadClickhouse,
		Update: resourceMonitorUpdateClickhouse,
		Delete: resourceMonitorDeleteClickhouse,
		Exists: resourceMonitorExistsClickhouse,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateClickhouse TODO Doc Comment
func resourceMonitorCreateClickhouse(d *schema.ResourceData, meta interface{}) error {
	apptype := "ClickHouse"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadClickhouse TODO Doc Comment
func resourceMonitorReadClickhouse(d *schema.ResourceData, meta interface{}) error {
	apptype := "ClickHouse"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateClickhouse TODO Doc Comment
func resourceMonitorUpdateClickhouse(d *schema.ResourceData, meta interface{}) error {
	apptype := "ClickHouse"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteClickhouse TODO Doc Comment
func resourceMonitorDeleteClickhouse(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "ClickHouse"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsClickhouse TODO Doc Comment
func resourceMonitorExistsClickhouse(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "ClickHouse"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportClickhouse TODO Doc Comment
func resourceSematextMonitorImportClickhouse(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "ClickHouse"
	return CommonMonitorImport(d, meta, apptype)
}
