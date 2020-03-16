package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorYarn TODO Doc Comment
func resourceSematextMonitorYarn() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateYarn,
		Read:   resourceMonitorReadYarn,
		Update: resourceMonitorUpdateYarn,
		Delete: resourceMonitorDeleteYarn,
		Exists: resourceMonitorExistsYarn,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateYarn TODO Doc Comment
func resourceMonitorCreateYarn(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Yarn"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadYarn TODO Doc Comment
func resourceMonitorReadYarn(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Yarn"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateYarn TODO Doc Comment
func resourceMonitorUpdateYarn(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Yarn"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteYarn TODO Doc Comment
func resourceMonitorDeleteYarn(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Yarn"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsYarn TODO Doc Comment
func resourceMonitorExistsYarn(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Yarn"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportYarn TODO Doc Comment
func resourceSematextMonitorImportYarn(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Yarn"))
	return CommonMonitorImport(d, meta)
}
