package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorHaproxy is the resource class that handles sematext_monitor_haproxy
func resourceSematextMonitorHaproxy() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("HAProxy")

	return &schema.Resource{
		Create: resourceMonitorCreateHaproxy,
		Read:   resourceMonitorReadHaproxy,
		Update: resourceMonitorUpdateHaproxy,
		Delete: resourceMonitorDeleteHaproxy,
		Exists: resourceMonitorExistsHaproxy,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateHaproxy creates the sematext_monitor_haproxy resource.
func resourceMonitorCreateHaproxy(d *schema.ResourceData, meta interface{}) error {
	apptype := "HAProxy"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadHaproxy reads the sematext_monitor_haproxy resource from Sematext Cloud.
func resourceMonitorReadHaproxy(d *schema.ResourceData, meta interface{}) error {
	apptype := "HAProxy"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateHaproxy updates Sematext Cloud from the sematext_monitor_haproxy resource.
func resourceMonitorUpdateHaproxy(d *schema.ResourceData, meta interface{}) error {
	apptype := "HAProxy"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteHaproxy marks a sematext_monitor_haproxy resource as retired.
func resourceMonitorDeleteHaproxy(d *schema.ResourceData, meta interface{}) error {
	apptype := "HAProxy"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsHaproxy checks a sematext_monitor_haproxy resource exists in Sematext Cloud.
func resourceMonitorExistsHaproxy(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "HAProxy"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportHaproxy checks a sematext_monitor_haproxy resource exists in Sematext Cloud.
func resourceSematextMonitorImportHaproxy(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "HAProxy"
	return CommonMonitorImport(d, meta, apptype)
}

*/
