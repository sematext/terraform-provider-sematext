package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorInfra TODO Doc Comment
func resourceSematextMonitorInfra() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

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
	apptype := "Infra"
	return CommonMonitorCreate(d, meta, apptype)
}

// resourceMonitorReadInfra TODO Doc Comment
func resourceMonitorReadInfra(d *schema.ResourceData, meta interface{}) error {
	apptype := "Infra"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateInfra TODO Doc Comment
func resourceMonitorUpdateInfra(d *schema.ResourceData, meta interface{}) error {
	apptype := "Infra"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteInfra TODO Doc Comment
func resourceMonitorDeleteInfra(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Infra"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsInfra TODO Doc Comment
func resourceMonitorExistsInfra(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Infra"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportInfra TODO Doc Comment
func resourceSematextMonitorImportInfra(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Infra"
	return CommonMonitorImport(d, meta, apptype)
}
