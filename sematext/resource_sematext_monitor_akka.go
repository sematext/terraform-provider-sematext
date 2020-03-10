package sematext

import (
    "strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// resourceSematextMonitorAkka TODO Doc Comment
func resourceSematextMonitorAkka() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateAkka,
		Read:   resourceSematextMonitorReadAkka,
		Update: resourceSematextMonitorUpdateAkka,
		Delete: resourceSematextMonitorDeleteAkka,
		Exists: resourceSematextMonitorExistsAkka,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateAkka TODO Doc Comment
func resourceSematextMonitorCreateAkka(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Akka"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadAkka TODO Doc Comment
func resourceSematextMonitorReadAkka(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Akka"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateAkka TODO Doc Comment
func resourceSematextMonitorUpdateAkka(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Akka"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteAkka TODO Doc Comment
func resourceSematextMonitorDeleteAkka(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Akka"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsAkka TODO Doc Comment
func resourceSematextMonitorExistsAkka(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Akka"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportAkka TODO Doc Comment
func resourceSematextMonitorImportAkka(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Akka"))
	return SematextMonitorImport(d, meta)
}
