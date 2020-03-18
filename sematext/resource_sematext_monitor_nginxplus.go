package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorNginxplus TODO Doc Comment
func resourceSematextMonitorNginxplus() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

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
	apptype := "Nginx-Plus"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadNginxplus TODO Doc Comment
func resourceMonitorReadNginxplus(d *schema.ResourceData, meta interface{}) error {
	apptype := "Nginx-Plus"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateNginxplus TODO Doc Comment
func resourceMonitorUpdateNginxplus(d *schema.ResourceData, meta interface{}) error {
	apptype := "Nginx-Plus"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteNginxplus TODO Doc Comment
func resourceMonitorDeleteNginxplus(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Nginx-Plus"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsNginxplus TODO Doc Comment
func resourceMonitorExistsNginxplus(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Nginx-Plus"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportNginxplus TODO Doc Comment
func resourceSematextMonitorImportNginxplus(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Nginx-Plus"
	return CommonMonitorImport(d, meta, apptype)
}
