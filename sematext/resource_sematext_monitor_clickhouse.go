package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorClickhouse TODO Doc Comment
func resourceSematextMonitorClickhouse() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateClickhouse,
		Read:   resourceSematextMonitorReadClickhouse,
		Update: resourceSematextMonitorUpdateClickhouse,
		Delete: resourceSematextMonitorDeleteClickhouse,
		Exists: resourceSematextMonitorExistsClickhouse,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateClickhouse TODO Doc Comment
func resourceSematextMonitorCreateClickhouse(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Clickhouse"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadClickhouse TODO Doc Comment
func resourceSematextMonitorReadClickhouse(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Clickhouse"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateClickhouse TODO Doc Comment
func resourceSematextMonitorUpdateClickhouse(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Clickhouse"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteClickhouse TODO Doc Comment
func resourceSematextMonitorDeleteClickhouse(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Clickhouse"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsClickhouse TODO Doc Comment
func resourceSematextMonitorExistsClickhouse(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Clickhouse"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportClickhouse TODO Doc Comment
func resourceSematextMonitorImportClickhouse(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Clickhouse"))
	return SematextMonitorImport(d, meta)
}
