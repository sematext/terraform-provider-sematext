package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorAWS TODO Doc Comment
func resourceSematextMonitorAWS() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateAWS,
		Read:   resourceSematextMonitorReadAWS,
		Update: resourceSematextMonitorUpdateAWS,
		Delete: resourceSematextMonitorDeleteAWS,
		Exists: resourceSematextMonitorExistsAWS,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateAWS TODO Doc Comment
func resourceSematextMonitorCreateAWS(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("AWS"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadAWS TODO Doc Comment
func resourceSematextMonitorReadAWS(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("AWS"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateAWS TODO Doc Comment
func resourceSematextMonitorUpdateAWS(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("AWS"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteAWS TODO Doc Comment
func resourceSematextMonitorDeleteAWS(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("AWS"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsAWS TODO Doc Comment
func resourceSematextMonitorExistsAWS(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("AWS"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportAWS TODO Doc Comment
func resourceSematextMonitorImportAWS(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("AWS"))
	return SematextMonitorImport(d, meta)
}
