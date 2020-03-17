package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorStorm TODO Doc Comment
func resourceSematextMonitorStorm() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateStorm,
		Read:   resourceMonitorReadStorm,
		Update: resourceMonitorUpdateStorm,
		Delete: resourceMonitorDeleteStorm,
		Exists: resourceMonitorExistsStorm,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateStorm TODO Doc Comment
func resourceMonitorCreateStorm(d *schema.ResourceData, meta interface{}) error {
	apptype := "Storm"
	return CommonMonitorCreate(d, meta, apptype)
}

// resourceMonitorReadStorm TODO Doc Comment
func resourceMonitorReadStorm(d *schema.ResourceData, meta interface{}) error {
	apptype := "Storm"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateStorm TODO Doc Comment
func resourceMonitorUpdateStorm(d *schema.ResourceData, meta interface{}) error {
	apptype := "Storm"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteStorm TODO Doc Comment
func resourceMonitorDeleteStorm(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Storm"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsStorm TODO Doc Comment
func resourceMonitorExistsStorm(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Storm"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportStorm TODO Doc Comment
func resourceSematextMonitorImportStorm(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Storm"
	return CommonMonitorImport(d, meta, apptype)
}
