package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorJvm is the resource class that handles sematext_monitor_jvm
func resourceSematextMonitorJvm() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("JVM")

	return &schema.Resource{
		Create: resourceMonitorCreateJvm,
		Read:   resourceMonitorReadJvm,
		Update: resourceMonitorUpdateJvm,
		Delete: resourceMonitorDeleteJvm,
		Exists: resourceMonitorExistsJvm,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateJvm creates the sematext_monitor_jvm resource.
func resourceMonitorCreateJvm(d *schema.ResourceData, meta interface{}) error {
	apptype := "JVM"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadJvm reads the sematext_monitor_jvm resource from Sematext Cloud.
func resourceMonitorReadJvm(d *schema.ResourceData, meta interface{}) error {
	apptype := "JVM"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateJvm updates Sematext Cloud from the sematext_monitor_jvm resource.
func resourceMonitorUpdateJvm(d *schema.ResourceData, meta interface{}) error {
	apptype := "JVM"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteJvm marks a sematext_monitor_jvm resource as retired.
func resourceMonitorDeleteJvm(d *schema.ResourceData, meta interface{}) error {
	apptype := "JVM"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsJvm checks a sematext_monitor_jvm resource exists in Sematext Cloud.
func resourceMonitorExistsJvm(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "JVM"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportJvm checks a sematext_monitor_jvm resource exists in Sematext Cloud.
func resourceSematextMonitorImportJvm(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "JVM"
	return CommonMonitorImport(d, meta, apptype)
}

*/
