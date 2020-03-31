package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorApache TODO Doc Comment
func resourceSematextMonitorApache() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Apache")

	return &schema.Resource{
		Create: resourceMonitorCreateApache,
		Read:   resourceMonitorReadApache,
		Update: resourceMonitorUpdateApache,
		Delete: resourceMonitorDeleteApache,
		Exists: resourceMonitorExistsApache,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateApache TODO Doc Comment
func resourceMonitorCreateApache(d *schema.ResourceData, meta interface{}) error {
	apptype := "Apache"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadApache TODO Doc Comment
func resourceMonitorReadApache(d *schema.ResourceData, meta interface{}) error {
	apptype := "Apache"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateApache TODO Doc Comment
func resourceMonitorUpdateApache(d *schema.ResourceData, meta interface{}) error {
	apptype := "Apache"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteApache TODO Doc Comment
func resourceMonitorDeleteApache(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Apache"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsApache TODO Doc Comment
func resourceMonitorExistsApache(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Apache"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportApache TODO Doc Comment
func resourceSematextMonitorImportApache(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Apache"
	return CommonMonitorImport(d, meta, apptype)
}
