package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorApache is the resource class that handles sematext_monitor_apache
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

// resourceMonitorCreateApache creates the sematext_monitor_apache resource.
func resourceMonitorCreateApache(d *schema.ResourceData, meta interface{}) error {
	apptype := "Apache"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadApache reads the sematext_monitor_apache resource from Sematext Cloud.
func resourceMonitorReadApache(d *schema.ResourceData, meta interface{}) error {
	apptype := "Apache"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateApache updates Sematext Cloud from the sematext_monitor_apache resource.
func resourceMonitorUpdateApache(d *schema.ResourceData, meta interface{}) error {
	apptype := "Apache"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteApache marks a sematext_monitor_apache resource as retired.
func resourceMonitorDeleteApache(d *schema.ResourceData, meta interface{}) error {
	apptype := "Apache"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsApache checks a sematext_monitor_apache resource exists in Sematext Cloud.
func resourceMonitorExistsApache(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Apache"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportApache checks a sematext_monitor_apache resource exists in Sematext Cloud.
func resourceSematextMonitorImportApache(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Apache"
	return CommonMonitorImport(d, meta, apptype)
}

*/
