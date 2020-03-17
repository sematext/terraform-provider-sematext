package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorMongodb TODO Doc Comment
func resourceSematextMonitorMongodb() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateMongodb,
		Read:   resourceMonitorReadMongodb,
		Update: resourceMonitorUpdateMongodb,
		Delete: resourceMonitorDeleteMongodb,
		Exists: resourceMonitorExistsMongodb,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateMongodb TODO Doc Comment
func resourceMonitorCreateMongodb(d *schema.ResourceData, meta interface{}) error {
	apptype := "MongoDB"
	return CommonMonitorCreate(d, meta, apptype)
}

// resourceMonitorReadMongodb TODO Doc Comment
func resourceMonitorReadMongodb(d *schema.ResourceData, meta interface{}) error {
	apptype := "MongoDB"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateMongodb TODO Doc Comment
func resourceMonitorUpdateMongodb(d *schema.ResourceData, meta interface{}) error {
	apptype := "MongoDB"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteMongodb TODO Doc Comment
func resourceMonitorDeleteMongodb(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "MongoDB"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsMongodb TODO Doc Comment
func resourceMonitorExistsMongodb(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "MongoDB"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportMongodb TODO Doc Comment
func resourceSematextMonitorImportMongodb(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "MongoDB"
	return CommonMonitorImport(d, meta, apptype)
}
