package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorMongodb is the resource class that handles sematext_monitor_mongodb
func resourceSematextMonitorMongodb() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("MongoDB")

	return &schema.Resource{
		Create: resourceMonitorCreateMongodb,
		Read:   resourceMonitorReadMongodb,
		Update: resourceMonitorUpdateMongodb,
		Delete: resourceMonitorDeleteMongodb,
		Exists: resourceMonitorExistsMongodb,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateMongodb creates the sematext_monitor_mongodb resource.
func resourceMonitorCreateMongodb(d *schema.ResourceData, meta interface{}) error {
	apptype := "MongoDB"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadMongodb reads the sematext_monitor_mongodb resource from Sematext Cloud.
func resourceMonitorReadMongodb(d *schema.ResourceData, meta interface{}) error {
	apptype := "MongoDB"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateMongodb updates Sematext Cloud from the sematext_monitor_mongodb resource.
func resourceMonitorUpdateMongodb(d *schema.ResourceData, meta interface{}) error {
	apptype := "MongoDB"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteMongodb marks a sematext_monitor_mongodb resource as retired.
func resourceMonitorDeleteMongodb(d *schema.ResourceData, meta interface{}) error {
	apptype := "MongoDB"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsMongodb checks a sematext_monitor_mongodb resource exists in Sematext Cloud.
func resourceMonitorExistsMongodb(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "MongoDB"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportMongodb checks a sematext_monitor_mongodb resource exists in Sematext Cloud.
func resourceSematextMonitorImportMongodb(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "MongoDB"
	return CommonMonitorImport(d, meta, apptype)
}

*/
