package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorStorm is the resource class that handles sematext_monitor_storm
func resourceSematextMonitorStorm() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Storm")

	return &schema.Resource{
		Create: resourceMonitorCreateStorm,
		Read:   resourceMonitorReadStorm,
		Update: resourceMonitorUpdateStorm,
		Delete: resourceMonitorDeleteStorm,
		Exists: resourceMonitorExistsStorm,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateStorm creates the sematext_monitor_storm resource.
func resourceMonitorCreateStorm(d *schema.ResourceData, meta interface{}) error {
	apptype := "Storm"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadStorm reads the sematext_monitor_storm resource from Sematext Cloud.
func resourceMonitorReadStorm(d *schema.ResourceData, meta interface{}) error {
	apptype := "Storm"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateStorm updates Sematext Cloud from the sematext_monitor_storm resource.
func resourceMonitorUpdateStorm(d *schema.ResourceData, meta interface{}) error {
	apptype := "Storm"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteStorm marks a sematext_monitor_storm resource as retired.
func resourceMonitorDeleteStorm(d *schema.ResourceData, meta interface{}) error {
	apptype := "Storm"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsStorm checks a sematext_monitor_storm resource exists in Sematext Cloud.
func resourceMonitorExistsStorm(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Storm"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportStorm checks a sematext_monitor_storm resource exists in Sematext Cloud.
func resourceSematextMonitorImportStorm(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Storm"
	return CommonMonitorImport(d, meta, apptype)
}

*/
