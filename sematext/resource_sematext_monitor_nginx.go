package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorNginx TODO Doc Comment
func resourceSematextMonitorNginx() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateNginx,
		Read:   resourceMonitorReadNginx,
		Update: resourceMonitorUpdateNginx,
		Delete: resourceMonitorDeleteNginx,
		Exists: resourceMonitorExistsNginx,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateNginx TODO Doc Comment
func resourceMonitorCreateNginx(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nginx"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadNginx TODO Doc Comment
func resourceMonitorReadNginx(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nginx"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateNginx TODO Doc Comment
func resourceMonitorUpdateNginx(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nginx"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteNginx TODO Doc Comment
func resourceMonitorDeleteNginx(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Nginx"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsNginx TODO Doc Comment
func resourceMonitorExistsNginx(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Nginx"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportNginx TODO Doc Comment
func resourceSematextMonitorImportNginx(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Nginx"))
	return CommonMonitorImport(d, meta)
}
