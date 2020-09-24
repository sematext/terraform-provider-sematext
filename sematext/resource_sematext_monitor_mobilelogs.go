package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorMobilelogs is the resource class that handles sematext_monitor_mobilelogs
func resourceSematextMonitorMobilelogs() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Mobile-Logs")

	return &schema.Resource{
		Create: resourceMonitorCreateMobilelogs,
		Read:   resourceMonitorReadMobilelogs,
		Update: resourceMonitorUpdateMobilelogs,
		Delete: resourceMonitorDeleteMobilelogs,
		Exists: resourceMonitorExistsMobilelogs,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateMobilelogs creates the sematext_monitor_mobilelogs resource.
func resourceMonitorCreateMobilelogs(d *schema.ResourceData, meta interface{}) error {
	apptype := "Mobile-Logs"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadMobilelogs reads the sematext_monitor_mobilelogs resource from Sematext Cloud.
func resourceMonitorReadMobilelogs(d *schema.ResourceData, meta interface{}) error {
	apptype := "Mobile-Logs"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateMobilelogs updates Sematext Cloud from the sematext_monitor_mobilelogs resource.
func resourceMonitorUpdateMobilelogs(d *schema.ResourceData, meta interface{}) error {
	apptype := "Mobile-Logs"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteMobilelogs marks a sematext_monitor_mobilelogs resource as retired.
func resourceMonitorDeleteMobilelogs(d *schema.ResourceData, meta interface{}) error {
	apptype := "Mobile-Logs"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsMobilelogs checks a sematext_monitor_mobilelogs resource exists in Sematext Cloud.
func resourceMonitorExistsMobilelogs(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Mobile-Logs"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportMobilelogs checks a sematext_monitor_mobilelogs resource exists in Sematext Cloud.
func resourceSematextMonitorImportMobilelogs(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Mobile-Logs"
	return CommonMonitorImport(d, meta, apptype)
}

*/
