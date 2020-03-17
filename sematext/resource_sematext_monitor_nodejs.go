package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorNodejs TODO Doc Comment
func resourceSematextMonitorNodejs() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

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
	apptype := "Node.js"
	return CommonMonitorCreate(d, meta, apptype)
}

// resourceMonitorReadNodejs TODO Doc Comment
func resourceMonitorReadNodejs(d *schema.ResourceData, meta interface{}) error {
	apptype := "Node.js"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateNodejs TODO Doc Comment
func resourceMonitorUpdateNodejs(d *schema.ResourceData, meta interface{}) error {
	apptype := "Node.js"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteNodejs TODO Doc Comment
func resourceMonitorDeleteNodejs(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Node.js"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsNodejs TODO Doc Comment
func resourceMonitorExistsNodejs(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Node.js"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportNodejs TODO Doc Comment
func resourceSematextMonitorImportNodejs(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Node.js"
	return CommonMonitorImport(d, meta, apptype)
}
