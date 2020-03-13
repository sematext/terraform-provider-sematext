package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorNginx TODO Doc Comment
func resourceSematextMonitorNginx() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateNginx,
		Read:   resourceSematextMonitorReadNginx,
		Update: resourceSematextMonitorUpdateNginx,
		Delete: resourceSematextMonitorDeleteNginx,
		Exists: resourceSematextMonitorExistsNginx,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateNginx TODO Doc Comment
func resourceSematextMonitorCreateNginx(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nginx"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadNginx TODO Doc Comment
func resourceSematextMonitorReadNginx(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nginx"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateNginx TODO Doc Comment
func resourceSematextMonitorUpdateNginx(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nginx"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteNginx TODO Doc Comment
func resourceSematextMonitorDeleteNginx(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Nginx"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsNginx TODO Doc Comment
func resourceSematextMonitorExistsNginx(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Nginx"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportNginx TODO Doc Comment
func resourceSematextMonitorImportNginx(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Nginx"))
	return SematextMonitorImport(d, meta)
}
