package sematext

import (
    "strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// resourceSematextMonitorSpark TODO Doc Comment
func resourceSematextMonitorSpark() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateSpark,
		Read:   resourceSematextMonitorReadSpark,
		Update: resourceSematextMonitorUpdateSpark,
		Delete: resourceSematextMonitorDeleteSpark,
		Exists: resourceSematextMonitorExistsSpark,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateSpark TODO Doc Comment
func resourceSematextMonitorCreateSpark(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Spark"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadSpark TODO Doc Comment
func resourceSematextMonitorReadSpark(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Spark"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateSpark TODO Doc Comment
func resourceSematextMonitorUpdateSpark(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Spark"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteSpark TODO Doc Comment
func resourceSematextMonitorDeleteSpark(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Spark"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsSpark TODO Doc Comment
func resourceSematextMonitorExistsSpark(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Spark"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportSpark TODO Doc Comment
func resourceSematextMonitorImportSpark(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Spark"))
	return SematextMonitorImport(d, meta)
}
