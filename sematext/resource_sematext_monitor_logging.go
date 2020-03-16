package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorLogging TODO Doc Comment
func resourceSematextMonitorLogging() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateLogging,
		Read:   resourceMonitorReadLogging,
		Update: resourceMonitorUpdateLogging,
		Delete: resourceMonitorDeleteLogging,
		Exists: resourceMonitorExistsLogging,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateLogging TODO Doc Comment
func resourceMonitorCreateLogging(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Logging"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadLogging TODO Doc Comment
func resourceMonitorReadLogging(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Logging"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateLogging TODO Doc Comment
func resourceMonitorUpdateLogging(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Logging"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteLogging TODO Doc Comment
func resourceMonitorDeleteLogging(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Logging"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsLogging TODO Doc Comment
func resourceMonitorExistsLogging(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Logging"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportLogging TODO Doc Comment
func resourceSematextMonitorImportLogging(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Logging"))
	return CommonMonitorImport(d, meta)
}
