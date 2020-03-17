package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorSensei TODO Doc Comment
func resourceSematextMonitorSensei() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateSensei,
		Read:   resourceMonitorReadSensei,
		Update: resourceMonitorUpdateSensei,
		Delete: resourceMonitorDeleteSensei,
		Exists: resourceMonitorExistsSensei,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateSensei TODO Doc Comment
func resourceMonitorCreateSensei(d *schema.ResourceData, meta interface{}) error {
	apptype := "Sensei"
	return CommonMonitorCreate(d, meta, apptype)
}

// resourceMonitorReadSensei TODO Doc Comment
func resourceMonitorReadSensei(d *schema.ResourceData, meta interface{}) error {
	apptype := "Sensei"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateSensei TODO Doc Comment
func resourceMonitorUpdateSensei(d *schema.ResourceData, meta interface{}) error {
	apptype := "Sensei"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteSensei TODO Doc Comment
func resourceMonitorDeleteSensei(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Sensei"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsSensei TODO Doc Comment
func resourceMonitorExistsSensei(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Sensei"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportSensei TODO Doc Comment
func resourceSematextMonitorImportSensei(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Sensei"
	return CommonMonitorImport(d, meta, apptype)
}
