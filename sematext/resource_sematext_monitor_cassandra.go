package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorCassandra TODO Doc Comment
func resourceSematextMonitorCassandra() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateCassandra,
		Read:   resourceMonitorReadCassandra,
		Update: resourceMonitorUpdateCassandra,
		Delete: resourceMonitorDeleteCassandra,
		Exists: resourceMonitorExistsCassandra,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateCassandra TODO Doc Comment
func resourceMonitorCreateCassandra(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Cassandra"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadCassandra TODO Doc Comment
func resourceMonitorReadCassandra(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Cassandra"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateCassandra TODO Doc Comment
func resourceMonitorUpdateCassandra(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Cassandra"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteCassandra TODO Doc Comment
func resourceMonitorDeleteCassandra(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Cassandra"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsCassandra TODO Doc Comment
func resourceMonitorExistsCassandra(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Cassandra"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportCassandra TODO Doc Comment
func resourceSematextMonitorImportCassandra(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Cassandra"))
	return CommonMonitorImport(d, meta)
}
