package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorNodejs is the resource class that handles sematext_monitor_nodejs
func resourceSematextMonitorNodejs() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Node.js")

	return &schema.Resource{
		Create: resourceMonitorCreateNodejs,
		Read:   resourceMonitorReadNodejs,
		Update: resourceMonitorUpdateNodejs,
		Delete: resourceMonitorDeleteNodejs,
		Exists: resourceMonitorExistsNodejs,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateNodejs creates the sematext_monitor_nodejs resource.
func resourceMonitorCreateNodejs(d *schema.ResourceData, meta interface{}) error {
	apptype := "Node.js"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadNodejs reads the sematext_monitor_nodejs resource from Sematext Cloud.
func resourceMonitorReadNodejs(d *schema.ResourceData, meta interface{}) error {
	apptype := "Node.js"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateNodejs updates Sematext Cloud from the sematext_monitor_nodejs resource.
func resourceMonitorUpdateNodejs(d *schema.ResourceData, meta interface{}) error {
	apptype := "Node.js"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteNodejs marks a sematext_monitor_nodejs resource as retired.
func resourceMonitorDeleteNodejs(d *schema.ResourceData, meta interface{}) error {
	apptype := "Node.js"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsNodejs checks a sematext_monitor_nodejs resource exists in Sematext Cloud.
func resourceMonitorExistsNodejs(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Node.js"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportNodejs checks a sematext_monitor_nodejs resource exists in Sematext Cloud.
func resourceSematextMonitorImportNodejs(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Node.js"
	return CommonMonitorImport(d, meta, apptype)
}

*/
