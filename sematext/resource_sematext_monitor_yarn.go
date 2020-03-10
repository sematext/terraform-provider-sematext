package sematext

import (
    "strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// resourceSematextMonitorYarn TODO Doc Comment
func resourceSematextMonitorYarn() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateYarn,
		Read:   resourceSematextMonitorReadYarn,
		Update: resourceSematextMonitorUpdateYarn,
		Delete: resourceSematextMonitorDeleteYarn,
		Exists: resourceSematextMonitorExistsYarn,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateYarn TODO Doc Comment
func resourceSematextMonitorCreateYarn(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Yarn"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadYarn TODO Doc Comment
func resourceSematextMonitorReadYarn(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Yarn"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateYarn TODO Doc Comment
func resourceSematextMonitorUpdateYarn(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Yarn"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteYarn TODO Doc Comment
func resourceSematextMonitorDeleteYarn(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Yarn"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsYarn TODO Doc Comment
func resourceSematextMonitorExistsYarn(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Yarn"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportYarn TODO Doc Comment
func resourceSematextMonitorImportYarn(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Yarn"))
	return SematextMonitorImport(d, meta)
}
