package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorLogsene is the resource class that handles sematext_monitor_logsene
func resourceSematextMonitorLogsene() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Logsene")

	return &schema.Resource{
		Create: resourceMonitorCreateLogsene,
		Read:   resourceMonitorReadLogsene,
		Update: resourceMonitorUpdateLogsene,
		Delete: resourceMonitorDeleteLogsene,
		Exists: resourceMonitorExistsLogsene,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateLogsene creates the sematext_monitor_logsene resource.
func resourceMonitorCreateLogsene(d *schema.ResourceData, meta interface{}) error {
	apptype := "Logsene"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadLogsene reads the sematext_monitor_logsene resource from Sematext Cloud.
func resourceMonitorReadLogsene(d *schema.ResourceData, meta interface{}) error {
	apptype := "Logsene"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateLogsene updates Sematext Cloud from the sematext_monitor_logsene resource.
func resourceMonitorUpdateLogsene(d *schema.ResourceData, meta interface{}) error {
	apptype := "Logsene"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteLogsene marks a sematext_monitor_logsene resource as retired.
func resourceMonitorDeleteLogsene(d *schema.ResourceData, meta interface{}) error {
	apptype := "Logsene"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsLogsene checks a sematext_monitor_logsene resource exists in Sematext Cloud.
func resourceMonitorExistsLogsene(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Logsene"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportLogsene checks a sematext_monitor_logsene resource exists in Sematext Cloud.
func resourceSematextMonitorImportLogsene(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Logsene"
	return CommonMonitorImport(d, meta, apptype)
}

*/
