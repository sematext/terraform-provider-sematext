package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorRedis TODO Doc Comment
func resourceSematextMonitorRedis() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateRedis,
		Read:   resourceMonitorReadRedis,
		Update: resourceMonitorUpdateRedis,
		Delete: resourceMonitorDeleteRedis,
		Exists: resourceMonitorExistsRedis,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateRedis TODO Doc Comment
func resourceMonitorCreateRedis(d *schema.ResourceData, meta interface{}) error {
	apptype := "Redis"
	return CommonMonitorCreate(d, meta, apptype)
}

// resourceMonitorReadRedis TODO Doc Comment
func resourceMonitorReadRedis(d *schema.ResourceData, meta interface{}) error {
	apptype := "Redis"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateRedis TODO Doc Comment
func resourceMonitorUpdateRedis(d *schema.ResourceData, meta interface{}) error {
	apptype := "Redis"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteRedis TODO Doc Comment
func resourceMonitorDeleteRedis(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Redis"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsRedis TODO Doc Comment
func resourceMonitorExistsRedis(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Redis"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportRedis TODO Doc Comment
func resourceSematextMonitorImportRedis(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Redis"
	return CommonMonitorImport(d, meta, apptype)
}
