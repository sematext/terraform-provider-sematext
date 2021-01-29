package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorInfra is the resource class that handles sematext_monitor_infra
func resourceSematextMonitorInfra() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Infra")

	return &schema.Resource{
		Create: resourceMonitorCreateInfra,
		Read:   resourceMonitorReadInfra,
		Update: resourceMonitorUpdateInfra,
		Delete: resourceMonitorDeleteInfra,
		Exists: resourceMonitorExistsInfra,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateInfra creates the sematext_monitor_infra resource.
func resourceMonitorCreateInfra(d *schema.ResourceData, meta interface{}) error {
	apptype := "Infra"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadInfra reads the sematext_monitor_infra resource from Sematext Cloud.
func resourceMonitorReadInfra(d *schema.ResourceData, meta interface{}) error {
	apptype := "Infra"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateInfra updates Sematext Cloud from the sematext_monitor_infra resource.
func resourceMonitorUpdateInfra(d *schema.ResourceData, meta interface{}) error {
	apptype := "Infra"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteInfra marks a sematext_monitor_infra resource as retired.
func resourceMonitorDeleteInfra(d *schema.ResourceData, meta interface{}) error {
	apptype := "Infra"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsInfra checks a sematext_monitor_infra resource exists in Sematext Cloud.
func resourceMonitorExistsInfra(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Infra"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportInfra checks a sematext_monitor_infra resource exists in Sematext Cloud.
func resourceSematextMonitorImportInfra(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Infra"
	return CommonMonitorImport(d, meta, apptype)
}

*/
