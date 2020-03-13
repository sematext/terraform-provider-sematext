package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorDocker TODO Doc Comment
func resourceSematextMonitorDocker() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateDocker,
		Read:   resourceSematextMonitorReadDocker,
		Update: resourceSematextMonitorUpdateDocker,
		Delete: resourceSematextMonitorDeleteDocker,
		Exists: resourceSematextMonitorExistsDocker,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateDocker TODO Doc Comment
func resourceSematextMonitorCreateDocker(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Docker"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadDocker TODO Doc Comment
func resourceSematextMonitorReadDocker(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Docker"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateDocker TODO Doc Comment
func resourceSematextMonitorUpdateDocker(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Docker"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteDocker TODO Doc Comment
func resourceSematextMonitorDeleteDocker(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Docker"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsDocker TODO Doc Comment
func resourceSematextMonitorExistsDocker(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Docker"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportDocker TODO Doc Comment
func resourceSematextMonitorImportDocker(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Docker"))
	return SematextMonitorImport(d, meta)
}
