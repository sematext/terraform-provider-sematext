package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorNginxplus TODO Doc Comment
func resourceSematextMonitorNginxplus() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateNginxplus,
		Read:   resourceMonitorReadNginxplus,
		Update: resourceMonitorUpdateNginxplus,
		Delete: resourceMonitorDeleteNginxplus,
		Exists: resourceMonitorExistsNginxplus,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateNginxplus TODO Doc Comment
func resourceMonitorCreateNginxplus(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nginxplus"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadNginxplus TODO Doc Comment
func resourceMonitorReadNginxplus(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nginxplus"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateNginxplus TODO Doc Comment
func resourceMonitorUpdateNginxplus(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Nginxplus"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteNginxplus TODO Doc Comment
func resourceMonitorDeleteNginxplus(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Nginxplus"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsNginxplus TODO Doc Comment
func resourceMonitorExistsNginxplus(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Nginxplus"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportNginxplus TODO Doc Comment
func resourceSematextMonitorImportNginxplus(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Nginxplus"))
	return CommonMonitorImport(d, meta)
}
