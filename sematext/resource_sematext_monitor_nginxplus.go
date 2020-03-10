package sematext

import (
    "strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// resourceSematextMonitorNginxplus TODO Doc Comment
func resourceSematextMonitorNginxplus() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateNginxplus,
		Read:   resourceSematextMonitorReadNginxplus,
		Update: resourceSematextMonitorUpdateNginxplus,
		Delete: resourceSematextMonitorDeleteNginxplus,
		Exists: resourceSematextMonitorExistsNginxplus,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateNginxplus TODO Doc Comment
func resourceSematextMonitorCreateNginxplus(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nginxplus"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadNginxplus TODO Doc Comment
func resourceSematextMonitorReadNginxplus(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nginxplus"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateNginxplus TODO Doc Comment
func resourceSematextMonitorUpdateNginxplus(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nginxplus"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteNginxplus TODO Doc Comment
func resourceSematextMonitorDeleteNginxplus(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Nginxplus"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsNginxplus TODO Doc Comment
func resourceSematextMonitorExistsNginxplus(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Nginxplus"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportNginxplus TODO Doc Comment
func resourceSematextMonitorImportNginxplus(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Nginxplus"))
	return SematextMonitorImport(d, meta)
}
