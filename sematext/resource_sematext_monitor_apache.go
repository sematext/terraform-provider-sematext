package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorApache TODO Doc Comment
func resourceSematextMonitorApache() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateApache,
		Read:   resourceMonitorReadApache,
		Update: resourceMonitorUpdateApache,
		Delete: resourceMonitorDeleteApache,
		Exists: resourceMonitorExistsApache,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateApache TODO Doc Comment
func resourceMonitorCreateApache(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Apache"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadApache TODO Doc Comment
func resourceMonitorReadApache(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Apache"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateApache TODO Doc Comment
func resourceMonitorUpdateApache(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Apache"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteApache TODO Doc Comment
func resourceMonitorDeleteApache(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Apache"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsApache TODO Doc Comment
func resourceMonitorExistsApache(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Apache"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportApache TODO Doc Comment
func resourceSematextMonitorImportApache(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Apache"))
	return CommonMonitorImport(d, meta)
}
