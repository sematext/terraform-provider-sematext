package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorSpark is the resource class that handles sematext_monitor_spark
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

// resourceMonitorCreateSpark creates the sematext_monitor_spark resource.
func resourceMonitorCreateSpark(d *schema.ResourceData, meta interface{}) error {
	apptype := "Spark"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadSpark reads the sematext_monitor_spark resource from Sematext Cloud.
func resourceMonitorReadSpark(d *schema.ResourceData, meta interface{}) error {
	apptype := "Spark"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateSpark updates Sematext Cloud from the sematext_monitor_spark resource.
func resourceMonitorUpdateSpark(d *schema.ResourceData, meta interface{}) error {
	apptype := "Spark"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteSpark marks a sematext_monitor_spark resource as retired.
func resourceMonitorDeleteSpark(d *schema.ResourceData, meta interface{}) error {
	apptype := "Spark"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsSpark checks a sematext_monitor_spark resource exists in Sematext Cloud.
func resourceMonitorExistsSpark(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Spark"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportSpark checks a sematext_monitor_spark resource exists in Sematext Cloud.
func resourceSematextMonitorImportSpark(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Spark"
	return CommonMonitorImport(d, meta, apptype)
}

*/
