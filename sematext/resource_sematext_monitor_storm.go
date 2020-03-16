package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorStorm TODO Doc Comment
func resourceSematextMonitorStorm() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateStorm,
		Read:   resourceMonitorReadStorm,
		Update: resourceMonitorUpdateStorm,
		Delete: resourceMonitorDeleteStorm,
		Exists: resourceMonitorExistsStorm,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateStorm TODO Doc Comment
func resourceMonitorCreateStorm(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Storm"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadStorm TODO Doc Comment
func resourceMonitorReadStorm(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Storm"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateStorm TODO Doc Comment
func resourceMonitorUpdateStorm(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Storm"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteStorm TODO Doc Comment
func resourceMonitorDeleteStorm(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Storm"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsStorm TODO Doc Comment
func resourceMonitorExistsStorm(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Storm"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportStorm TODO Doc Comment
func resourceSematextMonitorImportStorm(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Storm"))
	return CommonMonitorImport(d, meta)
}
