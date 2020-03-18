package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorJvm TODO Doc Comment
func resourceSematextMonitorJvm() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateJvm,
		Read:   resourceMonitorReadJvm,
		Update: resourceMonitorUpdateJvm,
		Delete: resourceMonitorDeleteJvm,
		Exists: resourceMonitorExistsJvm,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateJvm TODO Doc Comment
func resourceMonitorCreateJvm(d *schema.ResourceData, meta interface{}) error {
	apptype := "JVM"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadJvm TODO Doc Comment
func resourceMonitorReadJvm(d *schema.ResourceData, meta interface{}) error {
	apptype := "JVM"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateJvm TODO Doc Comment
func resourceMonitorUpdateJvm(d *schema.ResourceData, meta interface{}) error {
	apptype := "JVM"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteJvm TODO Doc Comment
func resourceMonitorDeleteJvm(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "JVM"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsJvm TODO Doc Comment
func resourceMonitorExistsJvm(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "JVM"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportJvm TODO Doc Comment
func resourceSematextMonitorImportJvm(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "JVM"
	return CommonMonitorImport(d, meta, apptype)
}
