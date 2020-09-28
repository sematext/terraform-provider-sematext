package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorAwselb is the resource class that handles sematext_monitor_awselb
func resourceSematextMonitorAwselb() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("AWS ELB")

	return &schema.Resource{
		Create: resourceMonitorCreateAwselb,
		Read:   resourceMonitorReadAwselb,
		Update: resourceMonitorUpdateAwselb,
		Delete: resourceMonitorDeleteAwselb,
		Exists: resourceMonitorExistsAwselb,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateAwselb creates the sematext_monitor_awselb resource.
func resourceMonitorCreateAwselb(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS ELB"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadAwselb reads the sematext_monitor_awselb resource from Sematext Cloud.
func resourceMonitorReadAwselb(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS ELB"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateAwselb updates Sematext Cloud from the sematext_monitor_awselb resource.
func resourceMonitorUpdateAwselb(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS ELB"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteAwselb marks a sematext_monitor_awselb resource as retired.
func resourceMonitorDeleteAwselb(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS ELB"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsAwselb checks a sematext_monitor_awselb resource exists in Sematext Cloud.
func resourceMonitorExistsAwselb(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "AWS ELB"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportAwselb checks a sematext_monitor_awselb resource exists in Sematext Cloud.
func resourceSematextMonitorImportAwselb(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "AWS ELB"
	return CommonMonitorImport(d, meta, apptype)
}

*/
