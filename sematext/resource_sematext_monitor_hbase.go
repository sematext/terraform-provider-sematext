package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorHbase TODO Doc Comment
func resourceSematextMonitorHbase() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Hbase")

	return &schema.Resource{
		Create: resourceMonitorCreateHbase,
		Read:   resourceMonitorReadHbase,
		Update: resourceMonitorUpdateHbase,
		Delete: resourceMonitorDeleteHbase,
		Exists: resourceMonitorExistsHbase,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateHbase TODO Doc Comment
func resourceMonitorCreateHbase(d *schema.ResourceData, meta interface{}) error {
	apptype := "HBase"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadHbase TODO Doc Comment
func resourceMonitorReadHbase(d *schema.ResourceData, meta interface{}) error {
	apptype := "HBase"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateHbase TODO Doc Comment
func resourceMonitorUpdateHbase(d *schema.ResourceData, meta interface{}) error {
	apptype := "HBase"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteHbase TODO Doc Comment
func resourceMonitorDeleteHbase(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "HBase"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsHbase TODO Doc Comment
func resourceMonitorExistsHbase(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "HBase"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportHbase TODO Doc Comment
func resourceSematextMonitorImportHbase(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "HBase"
	return CommonMonitorImport(d, meta, apptype)
}
