package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorMemcached TODO Doc Comment
func resourceSematextMonitorMemcached() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateMemcached,
		Read:   resourceMonitorReadMemcached,
		Update: resourceMonitorUpdateMemcached,
		Delete: resourceMonitorDeleteMemcached,
		Exists: resourceMonitorExistsMemcached,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateMemcached TODO Doc Comment
func resourceMonitorCreateMemcached(d *schema.ResourceData, meta interface{}) error {
	apptype := "Memcached"
	return CommonMonitorCreate(d, meta, apptype)
}

// resourceMonitorReadMemcached TODO Doc Comment
func resourceMonitorReadMemcached(d *schema.ResourceData, meta interface{}) error {
	apptype := "Memcached"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateMemcached TODO Doc Comment
func resourceMonitorUpdateMemcached(d *schema.ResourceData, meta interface{}) error {
	apptype := "Memcached"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteMemcached TODO Doc Comment
func resourceMonitorDeleteMemcached(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Memcached"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsMemcached TODO Doc Comment
func resourceMonitorExistsMemcached(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Memcached"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportMemcached TODO Doc Comment
func resourceSematextMonitorImportMemcached(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Memcached"
	return CommonMonitorImport(d, meta, apptype)
}
