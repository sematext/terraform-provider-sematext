package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorNodejs TODO Doc Comment
func resourceSematextMonitorNodejs() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateNodejs,
		Read:   resourceMonitorReadNodejs,
		Update: resourceMonitorUpdateNodejs,
		Delete: resourceMonitorDeleteNodejs,
		Exists: resourceMonitorExistsNodejs,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateNodejs TODO Doc Comment
func resourceMonitorCreateNodejs(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nodejs"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadNodejs TODO Doc Comment
func resourceMonitorReadNodejs(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nodejs"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateNodejs TODO Doc Comment
func resourceMonitorUpdateNodejs(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nodejs"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteNodejs TODO Doc Comment
func resourceMonitorDeleteNodejs(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Nodejs"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsNodejs TODO Doc Comment
func resourceMonitorExistsNodejs(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Nodejs"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportNodejs TODO Doc Comment
func resourceSematextMonitorImportNodejs(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Nodejs"))
	return CommonMonitorImport(d, meta)
}
