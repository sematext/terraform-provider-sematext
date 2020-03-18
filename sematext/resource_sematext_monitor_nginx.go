package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorNginx TODO Doc Comment
func resourceSematextMonitorNginx() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateNginx,
		Read:   resourceMonitorReadNginx,
		Update: resourceMonitorUpdateNginx,
		Delete: resourceMonitorDeleteNginx,
		Exists: resourceMonitorExistsNginx,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateNginx TODO Doc Comment
func resourceMonitorCreateNginx(d *schema.ResourceData, meta interface{}) error {
	apptype := "Nginx"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadNginx TODO Doc Comment
func resourceMonitorReadNginx(d *schema.ResourceData, meta interface{}) error {
	apptype := "Nginx"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateNginx TODO Doc Comment
func resourceMonitorUpdateNginx(d *schema.ResourceData, meta interface{}) error {
	apptype := "Nginx"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteNginx TODO Doc Comment
func resourceMonitorDeleteNginx(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Nginx"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsNginx TODO Doc Comment
func resourceMonitorExistsNginx(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Nginx"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportNginx TODO Doc Comment
func resourceSematextMonitorImportNginx(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Nginx"
	return CommonMonitorImport(d, meta, apptype)
}
