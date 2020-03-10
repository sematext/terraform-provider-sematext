package sematext

import (
    "strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// resourceSematextMonitorInfra TODO Doc Comment
func resourceSematextMonitorInfra() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateInfra,
		Read:   resourceSematextMonitorReadInfra,
		Update: resourceSematextMonitorUpdateInfra,
		Delete: resourceSematextMonitorDeleteInfra,
		Exists: resourceSematextMonitorExistsInfra,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateInfra TODO Doc Comment
func resourceSematextMonitorCreateInfra(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Infra"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadInfra TODO Doc Comment
func resourceSematextMonitorReadInfra(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Infra"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateInfra TODO Doc Comment
func resourceSematextMonitorUpdateInfra(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Infra"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteInfra TODO Doc Comment
func resourceSematextMonitorDeleteInfra(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Infra"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsInfra TODO Doc Comment
func resourceSematextMonitorExistsInfra(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Infra"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportInfra TODO Doc Comment
func resourceSematextMonitorImportInfra(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Infra"))
	return SematextMonitorImport(d, meta)
}
