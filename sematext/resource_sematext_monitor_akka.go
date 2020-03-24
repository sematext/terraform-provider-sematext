package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorAkka TODO Doc Comment
func resourceSematextMonitorAkka() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Akka")

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
	apptype := "Akka"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadAkka TODO Doc Comment
func resourceMonitorReadAkka(d *schema.ResourceData, meta interface{}) error {
	apptype := "Akka"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateAkka TODO Doc Comment
func resourceMonitorUpdateAkka(d *schema.ResourceData, meta interface{}) error {
	apptype := "Akka"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteAkka TODO Doc Comment
func resourceMonitorDeleteAkka(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Akka"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsAkka TODO Doc Comment
func resourceMonitorExistsAkka(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Akka"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportAkka TODO Doc Comment
func resourceSematextMonitorImportAkka(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Akka"
	return CommonMonitorImport(d, meta, apptype)
}
