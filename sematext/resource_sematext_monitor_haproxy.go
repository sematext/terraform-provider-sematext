package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorHaproxy TODO Doc Comment
func resourceSematextMonitorHaproxy() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("HAProxy")

	return &schema.Resource{
		Create: resourceMonitorCreateHaproxy,
		Read:   resourceMonitorReadHaproxy,
		Update: resourceMonitorUpdateHaproxy,
		Delete: resourceMonitorDeleteHaproxy,
		Exists: resourceMonitorExistsHaproxy,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateHaproxy TODO Doc Comment
func resourceMonitorCreateHaproxy(d *schema.ResourceData, meta interface{}) error {
	apptype := "HAProxy"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadHaproxy TODO Doc Comment
func resourceMonitorReadHaproxy(d *schema.ResourceData, meta interface{}) error {
	apptype := "HAProxy"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateHaproxy TODO Doc Comment
func resourceMonitorUpdateHaproxy(d *schema.ResourceData, meta interface{}) error {
	apptype := "HAProxy"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteHaproxy TODO Doc Comment
func resourceMonitorDeleteHaproxy(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "HAProxy"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsHaproxy TODO Doc Comment
func resourceMonitorExistsHaproxy(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "HAProxy"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportHaproxy TODO Doc Comment
func resourceSematextMonitorImportHaproxy(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "HAProxy"
	return CommonMonitorImport(d, meta, apptype)
}
