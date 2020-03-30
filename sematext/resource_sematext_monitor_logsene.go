package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorLogsene TODO Doc Comment
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

// resourceMonitorCreateLogsene TODO Doc Comment
func resourceMonitorCreateLogsene(d *schema.ResourceData, meta interface{}) error {
	apptype := "Logsene"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadLogsene TODO Doc Comment
func resourceMonitorReadLogsene(d *schema.ResourceData, meta interface{}) error {
	apptype := "Logsene"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateLogsene TODO Doc Comment
func resourceMonitorUpdateLogsene(d *schema.ResourceData, meta interface{}) error {
	apptype := "Logsene"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteLogsene TODO Doc Comment
func resourceMonitorDeleteLogsene(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Logsene"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsLogsene TODO Doc Comment
func resourceMonitorExistsLogsene(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Logsene"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportLogsene TODO Doc Comment
func resourceSematextMonitorImportLogsene(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Logsene"
	return CommonMonitorImport(d, meta, apptype)
}
