package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorRedis TODO Doc Comment
func resourceSematextMonitorRedis() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

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
	d.Set("app_type", strings.ToLower("Redis"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadRedis TODO Doc Comment
func resourceMonitorReadRedis(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Redis"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateRedis TODO Doc Comment
func resourceMonitorUpdateRedis(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Redis"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteRedis TODO Doc Comment
func resourceMonitorDeleteRedis(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Redis"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsRedis TODO Doc Comment
func resourceMonitorExistsRedis(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Redis"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportRedis TODO Doc Comment
func resourceSematextMonitorImportRedis(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Redis"))
	return CommonMonitorImport(d, meta)
}
