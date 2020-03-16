package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorInfra TODO Doc Comment
func resourceSematextMonitorInfra() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateInfra,
		Read:   resourceMonitorReadInfra,
		Update: resourceMonitorUpdateInfra,
		Delete: resourceMonitorDeleteInfra,
		Exists: resourceMonitorExistsInfra,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateInfra TODO Doc Comment
func resourceMonitorCreateInfra(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Infra"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadInfra TODO Doc Comment
func resourceMonitorReadInfra(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Infra"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateInfra TODO Doc Comment
func resourceMonitorUpdateInfra(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Infra"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteInfra TODO Doc Comment
func resourceMonitorDeleteInfra(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Infra"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsInfra TODO Doc Comment
func resourceMonitorExistsInfra(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Infra"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportInfra TODO Doc Comment
func resourceSematextMonitorImportInfra(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Infra"))
	return CommonMonitorImport(d, meta)
}
