package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorMongodb TODO Doc Comment
func resourceSematextMonitorMongodb() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

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
	d.Set("app_type", strings.ToLower("Mongodb"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadMongodb TODO Doc Comment
func resourceMonitorReadMongodb(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Mongodb"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateMongodb TODO Doc Comment
func resourceMonitorUpdateMongodb(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Mongodb"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteMongodb TODO Doc Comment
func resourceMonitorDeleteMongodb(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Mongodb"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsMongodb TODO Doc Comment
func resourceMonitorExistsMongodb(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Mongodb"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportMongodb TODO Doc Comment
func resourceSematextMonitorImportMongodb(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Mongodb"))
	return CommonMonitorImport(d, meta)
}
