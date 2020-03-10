package sematext

import (
    "strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// resourceSematextMonitorMongodb TODO Doc Comment
func resourceSematextMonitorMongodb() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateMongodb,
		Read:   resourceSematextMonitorReadMongodb,
		Update: resourceSematextMonitorUpdateMongodb,
		Delete: resourceSematextMonitorDeleteMongodb,
		Exists: resourceSematextMonitorExistsMongodb,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateMongodb TODO Doc Comment
func resourceSematextMonitorCreateMongodb(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Mongodb"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadMongodb TODO Doc Comment
func resourceSematextMonitorReadMongodb(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Mongodb"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateMongodb TODO Doc Comment
func resourceSematextMonitorUpdateMongodb(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Mongodb"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteMongodb TODO Doc Comment
func resourceSematextMonitorDeleteMongodb(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Mongodb"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsMongodb TODO Doc Comment
func resourceSematextMonitorExistsMongodb(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Mongodb"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportMongodb TODO Doc Comment
func resourceSematextMonitorImportMongodb(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Mongodb"))
	return SematextMonitorImport(d, meta)
}
