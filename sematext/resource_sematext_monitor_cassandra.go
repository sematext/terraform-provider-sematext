package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorCassandra is the resource class that handles sematext_monitor_cassandra
func resourceSematextMonitorCassandra() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Cassandra")

	return &schema.Resource{
		Create: resourceMonitorCreateCassandra,
		Read:   resourceMonitorReadCassandra,
		Update: resourceMonitorUpdateCassandra,
		Delete: resourceMonitorDeleteCassandra,
		Exists: resourceMonitorExistsCassandra,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateCassandra creates the sematext_monitor_cassandra resource.
func resourceMonitorCreateCassandra(d *schema.ResourceData, meta interface{}) error {
	apptype := "Cassandra"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadCassandra reads the sematext_monitor_cassandra resource from Sematext Cloud.
func resourceMonitorReadCassandra(d *schema.ResourceData, meta interface{}) error {
	apptype := "Cassandra"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateCassandra updates Sematext Cloud from the sematext_monitor_cassandra resource.
func resourceMonitorUpdateCassandra(d *schema.ResourceData, meta interface{}) error {
	apptype := "Cassandra"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteCassandra marks a sematext_monitor_cassandra resource as retired.
func resourceMonitorDeleteCassandra(d *schema.ResourceData, meta interface{}) error {
	apptype := "Cassandra"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsCassandra checks a sematext_monitor_cassandra resource exists in Sematext Cloud.
func resourceMonitorExistsCassandra(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Cassandra"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportCassandra checks a sematext_monitor_cassandra resource exists in Sematext Cloud.
func resourceSematextMonitorImportCassandra(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Cassandra"
	return CommonMonitorImport(d, meta, apptype)
}

*/
