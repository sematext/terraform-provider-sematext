package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorHaproxy TODO Doc Comment
func resourceSematextMonitorHaproxy() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateHaproxy,
		Read:   resourceMonitorReadHaproxy,
		Update: resourceMonitorUpdateHaproxy,
		Delete: resourceMonitorDeleteHaproxy,
		Exists: resourceMonitorExistsHaproxy,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateHaproxy TODO Doc Comment
func resourceMonitorCreateHaproxy(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Haproxy"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadHaproxy TODO Doc Comment
func resourceMonitorReadHaproxy(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Haproxy"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateHaproxy TODO Doc Comment
func resourceMonitorUpdateHaproxy(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Haproxy"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteHaproxy TODO Doc Comment
func resourceMonitorDeleteHaproxy(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Haproxy"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsHaproxy TODO Doc Comment
func resourceMonitorExistsHaproxy(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Haproxy"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportHaproxy TODO Doc Comment
func resourceSematextMonitorImportHaproxy(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Haproxy"))
	return CommonMonitorImport(d, meta)
}
