package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorSpark TODO Doc Comment
func resourceSematextMonitorSpark() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Spark")

	return &schema.Resource{
		Create: resourceMonitorCreateSpark,
		Read:   resourceMonitorReadSpark,
		Update: resourceMonitorUpdateSpark,
		Delete: resourceMonitorDeleteSpark,
		Exists: resourceMonitorExistsSpark,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateSpark TODO Doc Comment
func resourceMonitorCreateSpark(d *schema.ResourceData, meta interface{}) error {
	apptype := "Spark"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadSpark TODO Doc Comment
func resourceMonitorReadSpark(d *schema.ResourceData, meta interface{}) error {
	apptype := "Spark"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateSpark TODO Doc Comment
func resourceMonitorUpdateSpark(d *schema.ResourceData, meta interface{}) error {
	apptype := "Spark"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteSpark TODO Doc Comment
func resourceMonitorDeleteSpark(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Spark"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsSpark TODO Doc Comment
func resourceMonitorExistsSpark(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Spark"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportSpark TODO Doc Comment
func resourceSematextMonitorImportSpark(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Spark"
	return CommonMonitorImport(d, meta, apptype)
}
