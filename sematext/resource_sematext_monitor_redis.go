package sematext

import (
    "strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// resourceSematextMonitorRedis TODO Doc Comment
func resourceSematextMonitorRedis() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateRedis,
		Read:   resourceSematextMonitorReadRedis,
		Update: resourceSematextMonitorUpdateRedis,
		Delete: resourceSematextMonitorDeleteRedis,
		Exists: resourceSematextMonitorExistsRedis,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateRedis TODO Doc Comment
func resourceSematextMonitorCreateRedis(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Redis"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadRedis TODO Doc Comment
func resourceSematextMonitorReadRedis(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Redis"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateRedis TODO Doc Comment
func resourceSematextMonitorUpdateRedis(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Redis"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteRedis TODO Doc Comment
func resourceSematextMonitorDeleteRedis(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Redis"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsRedis TODO Doc Comment
func resourceSematextMonitorExistsRedis(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Redis"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportRedis TODO Doc Comment
func resourceSematextMonitorImportRedis(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Redis"))
	return SematextMonitorImport(d, meta)
}
