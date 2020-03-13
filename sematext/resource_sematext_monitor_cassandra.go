package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorCassandra TODO Doc Comment
func resourceSematextMonitorCassandra() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateCassandra,
		Read:   resourceSematextMonitorReadCassandra,
		Update: resourceSematextMonitorUpdateCassandra,
		Delete: resourceSematextMonitorDeleteCassandra,
		Exists: resourceSematextMonitorExistsCassandra,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateCassandra TODO Doc Comment
func resourceSematextMonitorCreateCassandra(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Cassandra"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadCassandra TODO Doc Comment
func resourceSematextMonitorReadCassandra(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Cassandra"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateCassandra TODO Doc Comment
func resourceSematextMonitorUpdateCassandra(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Cassandra"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteCassandra TODO Doc Comment
func resourceSematextMonitorDeleteCassandra(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Cassandra"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsCassandra TODO Doc Comment
func resourceSematextMonitorExistsCassandra(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Cassandra"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportCassandra TODO Doc Comment
func resourceSematextMonitorImportCassandra(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Cassandra"))
	return SematextMonitorImport(d, meta)
}
