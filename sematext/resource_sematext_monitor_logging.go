package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorLogging TODO Doc Comment
func resourceSematextMonitorLogging() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateLogging,
		Read:   resourceSematextMonitorReadLogging,
		Update: resourceSematextMonitorUpdateLogging,
		Delete: resourceSematextMonitorDeleteLogging,
		Exists: resourceSematextMonitorExistsLogging,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateLogging TODO Doc Comment
func resourceSematextMonitorCreateLogging(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Logging"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadLogging TODO Doc Comment
func resourceSematextMonitorReadLogging(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Logging"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateLogging TODO Doc Comment
func resourceSematextMonitorUpdateLogging(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Logging"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteLogging TODO Doc Comment
func resourceSematextMonitorDeleteLogging(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Logging"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsLogging TODO Doc Comment
func resourceSematextMonitorExistsLogging(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Logging"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportLogging TODO Doc Comment
func resourceSematextMonitorImportLogging(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Logging"))
	return SematextMonitorImport(d, meta)
}
