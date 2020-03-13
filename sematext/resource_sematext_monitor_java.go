package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorJava TODO Doc Comment
func resourceSematextMonitorJava() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateJava,
		Read:   resourceSematextMonitorReadJava,
		Update: resourceSematextMonitorUpdateJava,
		Delete: resourceSematextMonitorDeleteJava,
		Exists: resourceSematextMonitorExistsJava,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateJava TODO Doc Comment
func resourceSematextMonitorCreateJava(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Java"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadJava TODO Doc Comment
func resourceSematextMonitorReadJava(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Java"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateJava TODO Doc Comment
func resourceSematextMonitorUpdateJava(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Java"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteJava TODO Doc Comment
func resourceSematextMonitorDeleteJava(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Java"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsJava TODO Doc Comment
func resourceSematextMonitorExistsJava(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Java"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportJava TODO Doc Comment
func resourceSematextMonitorImportJava(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Java"))
	return SematextMonitorImport(d, meta)
}
