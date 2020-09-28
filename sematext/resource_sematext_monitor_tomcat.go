package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorTomcat is the resource class that handles sematext_monitor_tomcat
func resourceSematextMonitorTomcat() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Tomcat")

	return &schema.Resource{
		Create: resourceMonitorCreateTomcat,
		Read:   resourceMonitorReadTomcat,
		Update: resourceMonitorUpdateTomcat,
		Delete: resourceMonitorDeleteTomcat,
		Exists: resourceMonitorExistsTomcat,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateTomcat creates the sematext_monitor_tomcat resource.
func resourceMonitorCreateTomcat(d *schema.ResourceData, meta interface{}) error {
	apptype := "Tomcat"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadTomcat reads the sematext_monitor_tomcat resource from Sematext Cloud.
func resourceMonitorReadTomcat(d *schema.ResourceData, meta interface{}) error {
	apptype := "Tomcat"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateTomcat updates Sematext Cloud from the sematext_monitor_tomcat resource.
func resourceMonitorUpdateTomcat(d *schema.ResourceData, meta interface{}) error {
	apptype := "Tomcat"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteTomcat marks a sematext_monitor_tomcat resource as retired.
func resourceMonitorDeleteTomcat(d *schema.ResourceData, meta interface{}) error {
	apptype := "Tomcat"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsTomcat checks a sematext_monitor_tomcat resource exists in Sematext Cloud.
func resourceMonitorExistsTomcat(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Tomcat"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportTomcat checks a sematext_monitor_tomcat resource exists in Sematext Cloud.
func resourceSematextMonitorImportTomcat(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Tomcat"
	return CommonMonitorImport(d, meta, apptype)
}

*/
