package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorStorm TODO Doc Comment
func resourceSematextMonitorStorm() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateStorm,
		Read:   resourceSematextMonitorReadStorm,
		Update: resourceSematextMonitorUpdateStorm,
		Delete: resourceSematextMonitorDeleteStorm,
		Exists: resourceSematextMonitorExistsStorm,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateStorm TODO Doc Comment
func resourceSematextMonitorCreateStorm(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Storm"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadStorm TODO Doc Comment
func resourceSematextMonitorReadStorm(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Storm"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateStorm TODO Doc Comment
func resourceSematextMonitorUpdateStorm(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Storm"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteStorm TODO Doc Comment
func resourceSematextMonitorDeleteStorm(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Storm"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsStorm TODO Doc Comment
func resourceSematextMonitorExistsStorm(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Storm"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportStorm TODO Doc Comment
func resourceSematextMonitorImportStorm(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Storm"))
	return SematextMonitorImport(d, meta)
}
