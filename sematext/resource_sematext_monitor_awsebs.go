package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorAwsebs is the resource class that handles sematext_monitor_awsebs
func resourceSematextMonitorAwsebs() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("AWS EBS")

	return &schema.Resource{
		Create: resourceMonitorCreateAwsebs,
		Read:   resourceMonitorReadAwsebs,
		Update: resourceMonitorUpdateAwsebs,
		Delete: resourceMonitorDeleteAwsebs,
		Exists: resourceMonitorExistsAwsebs,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateAwsebs creates the sematext_monitor_awsebs resource.
func resourceMonitorCreateAwsebs(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS EBS"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadAwsebs reads the sematext_monitor_awsebs resource from Sematext Cloud.
func resourceMonitorReadAwsebs(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS EBS"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateAwsebs updates Sematext Cloud from the sematext_monitor_awsebs resource.
func resourceMonitorUpdateAwsebs(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS EBS"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteAwsebs marks a sematext_monitor_awsebs resource as retired.
func resourceMonitorDeleteAwsebs(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS EBS"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsAwsebs checks a sematext_monitor_awsebs resource exists in Sematext Cloud.
func resourceMonitorExistsAwsebs(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "AWS EBS"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportAwsebs checks a sematext_monitor_awsebs resource exists in Sematext Cloud.
func resourceSematextMonitorImportAwsebs(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "AWS EBS"
	return CommonMonitorImport(d, meta, apptype)
}

*/
