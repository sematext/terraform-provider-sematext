package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorHbase is the resource class that handles sematext_monitor_hbase
func resourceSematextMonitorHbase() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("HBase")

	return &schema.Resource{
		Create: resourceMonitorCreateHbase,
		Read:   resourceMonitorReadHbase,
		Update: resourceMonitorUpdateHbase,
		Delete: resourceMonitorDeleteHbase,
		Exists: resourceMonitorExistsHbase,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateHbase creates the sematext_monitor_hbase resource.
func resourceMonitorCreateHbase(d *schema.ResourceData, meta interface{}) error {
	apptype := "HBase"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadHbase reads the sematext_monitor_hbase resource from Sematext Cloud.
func resourceMonitorReadHbase(d *schema.ResourceData, meta interface{}) error {
	apptype := "HBase"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateHbase updates Sematext Cloud from the sematext_monitor_hbase resource.
func resourceMonitorUpdateHbase(d *schema.ResourceData, meta interface{}) error {
	apptype := "HBase"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteHbase marks a sematext_monitor_hbase resource as retired.
func resourceMonitorDeleteHbase(d *schema.ResourceData, meta interface{}) error {
	apptype := "HBase"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsHbase checks a sematext_monitor_hbase resource exists in Sematext Cloud.
func resourceMonitorExistsHbase(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "HBase"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportHbase checks a sematext_monitor_hbase resource exists in Sematext Cloud.
func resourceSematextMonitorImportHbase(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "HBase"
	return CommonMonitorImport(d, meta, apptype)
}

*/
