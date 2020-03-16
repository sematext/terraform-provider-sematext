package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorAWS TODO Doc Comment
func resourceSematextMonitorAWS() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateAWS,
		Read:   resourceMonitorReadAWS,
		Update: resourceMonitorUpdateAWS,
		Delete: resourceMonitorDeleteAWS,
		Exists: resourceMonitorExistsAWS,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateAWS TODO Doc Comment
func resourceMonitorCreateAWS(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("AWS"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadAWS TODO Doc Comment
func resourceMonitorReadAWS(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("AWS"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateAWS TODO Doc Comment
func resourceMonitorUpdateAWS(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("AWS"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteAWS TODO Doc Comment
func resourceMonitorDeleteAWS(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("AWS"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsAWS TODO Doc Comment
func resourceMonitorExistsAWS(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("AWS"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportAWS TODO Doc Comment
func resourceSematextMonitorImportAWS(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("AWS"))
	return CommonMonitorImport(d, meta)
}
