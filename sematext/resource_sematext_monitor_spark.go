package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorSpark TODO Doc Comment
func resourceSematextMonitorSpark() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateSpark,
		Read:   resourceMonitorReadSpark,
		Update: resourceMonitorUpdateSpark,
		Delete: resourceMonitorDeleteSpark,
		Exists: resourceMonitorExistsSpark,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateSpark TODO Doc Comment
func resourceMonitorCreateSpark(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Spark"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadSpark TODO Doc Comment
func resourceMonitorReadSpark(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Spark"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateSpark TODO Doc Comment
func resourceMonitorUpdateSpark(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Spark"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteSpark TODO Doc Comment
func resourceMonitorDeleteSpark(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Spark"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsSpark TODO Doc Comment
func resourceMonitorExistsSpark(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Spark"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportSpark TODO Doc Comment
func resourceSematextMonitorImportSpark(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Spark"))
	return CommonMonitorImport(d, meta)
}
