package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorDocker TODO Doc Comment
func resourceSematextMonitorDocker() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

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
	d.Set("app_type", strings.ToLower("Docker"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadDocker TODO Doc Comment
func resourceMonitorReadDocker(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Docker"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateDocker TODO Doc Comment
func resourceMonitorUpdateDocker(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Docker"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteDocker TODO Doc Comment
func resourceMonitorDeleteDocker(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Docker"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsDocker TODO Doc Comment
func resourceMonitorExistsDocker(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Docker"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportDocker TODO Doc Comment
func resourceSematextMonitorImportDocker(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Docker"))
	return CommonMonitorImport(d, meta)
}
