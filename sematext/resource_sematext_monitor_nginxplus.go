package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorNginxplus is the resource class that handles sematext_monitor_nginxplus
func resourceSematextMonitorNginxplus() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Nginx-Plus")

	return &schema.Resource{
		Create: resourceMonitorCreateNginxplus,
		Read:   resourceMonitorReadNginxplus,
		Update: resourceMonitorUpdateNginxplus,
		Delete: resourceMonitorDeleteNginxplus,
		Exists: resourceMonitorExistsNginxplus,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateNginxplus creates the sematext_monitor_nginxplus resource.
func resourceMonitorCreateNginxplus(d *schema.ResourceData, meta interface{}) error {
	apptype := "Nginx-Plus"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadNginxplus reads the sematext_monitor_nginxplus resource from Sematext Cloud.
func resourceMonitorReadNginxplus(d *schema.ResourceData, meta interface{}) error {
	apptype := "Nginx-Plus"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateNginxplus updates Sematext Cloud from the sematext_monitor_nginxplus resource.
func resourceMonitorUpdateNginxplus(d *schema.ResourceData, meta interface{}) error {
	apptype := "Nginx-Plus"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteNginxplus marks a sematext_monitor_nginxplus resource as retired.
func resourceMonitorDeleteNginxplus(d *schema.ResourceData, meta interface{}) error {
	apptype := "Nginx-Plus"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsNginxplus checks a sematext_monitor_nginxplus resource exists in Sematext Cloud.
func resourceMonitorExistsNginxplus(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Nginx-Plus"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportNginxplus checks a sematext_monitor_nginxplus resource exists in Sematext Cloud.
func resourceSematextMonitorImportNginxplus(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Nginx-Plus"
	return CommonMonitorImport(d, meta, apptype)
}

*/
