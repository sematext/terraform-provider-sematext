package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorHbase TODO Doc Comment
func resourceSematextMonitorHbase() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateHbase,
		Read:   resourceMonitorReadHbase,
		Update: resourceMonitorUpdateHbase,
		Delete: resourceMonitorDeleteHbase,
		Exists: resourceMonitorExistsHbase,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateHbase TODO Doc Comment
func resourceMonitorCreateHbase(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Hbase"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadHbase TODO Doc Comment
func resourceMonitorReadHbase(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Hbase"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateHbase TODO Doc Comment
func resourceMonitorUpdateHbase(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Hbase"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteHbase TODO Doc Comment
func resourceMonitorDeleteHbase(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Hbase"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsHbase TODO Doc Comment
func resourceMonitorExistsHbase(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Hbase"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportHbase TODO Doc Comment
func resourceSematextMonitorImportHbase(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Hbase"))
	return CommonMonitorImport(d, meta)
}
