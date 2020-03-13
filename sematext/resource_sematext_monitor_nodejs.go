package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorNodejs TODO Doc Comment
func resourceSematextMonitorNodejs() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateNodejs,
		Read:   resourceSematextMonitorReadNodejs,
		Update: resourceSematextMonitorUpdateNodejs,
		Delete: resourceSematextMonitorDeleteNodejs,
		Exists: resourceSematextMonitorExistsNodejs,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateNodejs TODO Doc Comment
func resourceSematextMonitorCreateNodejs(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nodejs"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadNodejs TODO Doc Comment
func resourceSematextMonitorReadNodejs(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nodejs"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateNodejs TODO Doc Comment
func resourceSematextMonitorUpdateNodejs(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nodejs"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteNodejs TODO Doc Comment
func resourceSematextMonitorDeleteNodejs(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Nodejs"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsNodejs TODO Doc Comment
func resourceSematextMonitorExistsNodejs(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Nodejs"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportNodejs TODO Doc Comment
func resourceSematextMonitorImportNodejs(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Nodejs"))
	return SematextMonitorImport(d, meta)
}
