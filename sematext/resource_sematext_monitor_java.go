package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorJava TODO Doc Comment
func resourceSematextMonitorJava() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateJava,
		Read:   resourceMonitorReadJava,
		Update: resourceMonitorUpdateJava,
		Delete: resourceMonitorDeleteJava,
		Exists: resourceMonitorExistsJava,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateJava TODO Doc Comment
func resourceMonitorCreateJava(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Java"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadJava TODO Doc Comment
func resourceMonitorReadJava(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Java"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateJava TODO Doc Comment
func resourceMonitorUpdateJava(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Java"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteJava TODO Doc Comment
func resourceMonitorDeleteJava(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Java"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsJava TODO Doc Comment
func resourceMonitorExistsJava(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Java"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportJava TODO Doc Comment
func resourceSematextMonitorImportJava(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Java"))
	return CommonMonitorImport(d, meta)
}
