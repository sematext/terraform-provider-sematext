package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorClickhouse TODO Doc Comment
func resourceSematextMonitorClickhouse() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

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
	d.Set("app_type", strings.ToLower("Clickhouse"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadClickhouse TODO Doc Comment
func resourceMonitorReadClickhouse(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Clickhouse"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateClickhouse TODO Doc Comment
func resourceMonitorUpdateClickhouse(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Clickhouse"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteClickhouse TODO Doc Comment
func resourceMonitorDeleteClickhouse(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Clickhouse"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsClickhouse TODO Doc Comment
func resourceMonitorExistsClickhouse(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Clickhouse"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportClickhouse TODO Doc Comment
func resourceSematextMonitorImportClickhouse(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Clickhouse"))
	return CommonMonitorImport(d, meta)
}
