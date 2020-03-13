package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorApache TODO Doc Comment
func resourceSematextMonitorApache() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateApache,
		Read:   resourceSematextMonitorReadApache,
		Update: resourceSematextMonitorUpdateApache,
		Delete: resourceSematextMonitorDeleteApache,
		Exists: resourceSematextMonitorExistsApache,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateApache TODO Doc Comment
func resourceSematextMonitorCreateApache(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Apache"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadApache TODO Doc Comment
func resourceSematextMonitorReadApache(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Apache"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateApache TODO Doc Comment
func resourceSematextMonitorUpdateApache(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Apache"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteApache TODO Doc Comment
func resourceSematextMonitorDeleteApache(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Apache"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsApache TODO Doc Comment
func resourceSematextMonitorExistsApache(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Apache"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportApache TODO Doc Comment
func resourceSematextMonitorImportApache(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Apache"))
	return SematextMonitorImport(d, meta)
}
