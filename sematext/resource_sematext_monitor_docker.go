package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorDocker TODO Doc Comment
func resourceSematextMonitorDocker() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateDocker,
		Read:   resourceMonitorReadDocker,
		Update: resourceMonitorUpdateDocker,
		Delete: resourceMonitorDeleteDocker,
		Exists: resourceMonitorExistsDocker,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateDocker TODO Doc Comment
func resourceMonitorCreateDocker(d *schema.ResourceData, meta interface{}) error {
	apptype := "Docker"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadDocker TODO Doc Comment
func resourceMonitorReadDocker(d *schema.ResourceData, meta interface{}) error {
	apptype := "Docker"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateDocker TODO Doc Comment
func resourceMonitorUpdateDocker(d *schema.ResourceData, meta interface{}) error {
	apptype := "Docker"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteDocker TODO Doc Comment
func resourceMonitorDeleteDocker(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Docker"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsDocker TODO Doc Comment
func resourceMonitorExistsDocker(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Docker"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportDocker TODO Doc Comment
func resourceSematextMonitorImportDocker(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Docker"
	return CommonMonitorImport(d, meta, apptype)
}
