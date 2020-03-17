package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorCassandra TODO Doc Comment
func resourceSematextMonitorCassandra() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

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
	apptype := "Cassandra"
	return CommonMonitorCreate(d, meta, apptype)
}

// resourceMonitorReadCassandra TODO Doc Comment
func resourceMonitorReadCassandra(d *schema.ResourceData, meta interface{}) error {
	apptype := "Cassandra"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateCassandra TODO Doc Comment
func resourceMonitorUpdateCassandra(d *schema.ResourceData, meta interface{}) error {
	apptype := "Cassandra"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteCassandra TODO Doc Comment
func resourceMonitorDeleteCassandra(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Cassandra"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsCassandra TODO Doc Comment
func resourceMonitorExistsCassandra(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Cassandra"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportCassandra TODO Doc Comment
func resourceSematextMonitorImportCassandra(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Cassandra"
	return CommonMonitorImport(d, meta, apptype)
}
