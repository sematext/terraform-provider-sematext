package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorAkka TODO Doc Comment
func resourceSematextMonitorAkka() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateAkka,
		Read:   resourceMonitorReadAkka,
		Update: resourceMonitorUpdateAkka,
		Delete: resourceMonitorDeleteAkka,
		Exists: resourceMonitorExistsAkka,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateAkka TODO Doc Comment
func resourceMonitorCreateAkka(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Akka"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadAkka TODO Doc Comment
func resourceMonitorReadAkka(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Akka"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateAkka TODO Doc Comment
func resourceMonitorUpdateAkka(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Akka"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteAkka TODO Doc Comment
func resourceMonitorDeleteAkka(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Akka"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsAkka TODO Doc Comment
func resourceMonitorExistsAkka(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Akka"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportAkka TODO Doc Comment
func resourceSematextMonitorImportAkka(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Akka"))
	return CommonMonitorImport(d, meta)
}
