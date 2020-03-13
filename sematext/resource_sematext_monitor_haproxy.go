package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorHaproxy TODO Doc Comment
func resourceSematextMonitorHaproxy() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateHaproxy,
		Read:   resourceSematextMonitorReadHaproxy,
		Update: resourceSematextMonitorUpdateHaproxy,
		Delete: resourceSematextMonitorDeleteHaproxy,
		Exists: resourceSematextMonitorExistsHaproxy,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateHaproxy TODO Doc Comment
func resourceSematextMonitorCreateHaproxy(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Haproxy"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadHaproxy TODO Doc Comment
func resourceSematextMonitorReadHaproxy(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Haproxy"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateHaproxy TODO Doc Comment
func resourceSematextMonitorUpdateHaproxy(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Haproxy"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteHaproxy TODO Doc Comment
func resourceSematextMonitorDeleteHaproxy(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Haproxy"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsHaproxy TODO Doc Comment
func resourceSematextMonitorExistsHaproxy(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Haproxy"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportHaproxy TODO Doc Comment
func resourceSematextMonitorImportHaproxy(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Haproxy"))
	return SematextMonitorImport(d, meta)
}
