package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorDocker is the resource class that handles sematext_monitor_docker
func resourceSematextMonitorDocker() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Docker")

	return &schema.Resource{
		Create: resourceMonitorCreateDocker,
		Read:   resourceMonitorReadDocker,
		Update: resourceMonitorUpdateDocker,
		Delete: resourceMonitorDeleteDocker,
		Exists: resourceMonitorExistsDocker,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateDocker creates the sematext_monitor_docker resource.
func resourceMonitorCreateDocker(d *schema.ResourceData, meta interface{}) error {
	apptype := "Docker"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadDocker reads the sematext_monitor_docker resource from Sematext Cloud.
func resourceMonitorReadDocker(d *schema.ResourceData, meta interface{}) error {
	apptype := "Docker"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateDocker updates Sematext Cloud from the sematext_monitor_docker resource.
func resourceMonitorUpdateDocker(d *schema.ResourceData, meta interface{}) error {
	apptype := "Docker"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteDocker marks a sematext_monitor_docker resource as retired.
func resourceMonitorDeleteDocker(d *schema.ResourceData, meta interface{}) error {
	apptype := "Docker"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsDocker checks a sematext_monitor_docker resource exists in Sematext Cloud.
func resourceMonitorExistsDocker(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Docker"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportDocker checks a sematext_monitor_docker resource exists in Sematext Cloud.
func resourceSematextMonitorImportDocker(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Docker"
	return CommonMonitorImport(d, meta, apptype)
}

*/
