package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorNginx is the resource class that handles sematext_monitor_nginx
func resourceSematextMonitorNginx() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Nginx")

	return &schema.Resource{
		Create: resourceMonitorCreateNginx,
		Read:   resourceMonitorReadNginx,
		Update: resourceMonitorUpdateNginx,
		Delete: resourceMonitorDeleteNginx,
		Exists: resourceMonitorExistsNginx,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateNginx creates the sematext_monitor_nginx resource.
func resourceMonitorCreateNginx(d *schema.ResourceData, meta interface{}) error {
	apptype := "Nginx"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadNginx reads the sematext_monitor_nginx resource from Sematext Cloud.
func resourceMonitorReadNginx(d *schema.ResourceData, meta interface{}) error {
	apptype := "Nginx"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateNginx updates Sematext Cloud from the sematext_monitor_nginx resource.
func resourceMonitorUpdateNginx(d *schema.ResourceData, meta interface{}) error {
	apptype := "Nginx"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteNginx marks a sematext_monitor_nginx resource as retired.
func resourceMonitorDeleteNginx(d *schema.ResourceData, meta interface{}) error {
	apptype := "Nginx"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsNginx checks a sematext_monitor_nginx resource exists in Sematext Cloud.
func resourceMonitorExistsNginx(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Nginx"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportNginx checks a sematext_monitor_nginx resource exists in Sematext Cloud.
func resourceSematextMonitorImportNginx(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Nginx"
	return CommonMonitorImport(d, meta, apptype)
}

*/
