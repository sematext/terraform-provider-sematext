package sematext

import (
    "strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// resourceSematextMonitorHbase TODO Doc Comment
func resourceSematextMonitorHbase() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateHbase,
		Read:   resourceSematextMonitorReadHbase,
		Update: resourceSematextMonitorUpdateHbase,
		Delete: resourceSematextMonitorDeleteHbase,
		Exists: resourceSematextMonitorExistsHbase,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateHbase TODO Doc Comment
func resourceSematextMonitorCreateHbase(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Hbase"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadHbase TODO Doc Comment
func resourceSematextMonitorReadHbase(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Hbase"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateHbase TODO Doc Comment
func resourceSematextMonitorUpdateHbase(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Hbase"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteHbase TODO Doc Comment
func resourceSematextMonitorDeleteHbase(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Hbase"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsHbase TODO Doc Comment
func resourceSematextMonitorExistsHbase(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Hbase"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportHbase TODO Doc Comment
func resourceSematextMonitorImportHbase(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Hbase"))
	return SematextMonitorImport(d, meta)
}
