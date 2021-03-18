package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorCassandra is the resource class that handles sematext_monitor_cassandra
func resourceSematextMonitorCassandra() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Cassandra")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateCassandra,
		ReadContext:   resourceMonitorReadCassandra,
		UpdateContext: resourceMonitorUpdateCassandra,
		DeleteContext: resourceMonitorDeleteCassandra,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateCassandra creates the sematext_monitor_cassandra resource.
func resourceMonitorCreateCassandra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Cassandra"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadCassandra reads the sematext_monitor_cassandra resource from Sematext Cloud.
func resourceMonitorReadCassandra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Cassandra"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateCassandra updates Sematext Cloud from the sematext_monitor_cassandra resource.
func resourceMonitorUpdateCassandra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Cassandra"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteCassandra marks a sematext_monitor_cassandra resource as retired.
func resourceMonitorDeleteCassandra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Cassandra"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportCassandra checks a sematext_monitor_cassandra resource exists in Sematext Cloud.
func resourceSematextMonitorImportCassandra(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Cassandra"
	return CommonMonitorImport(d, meta, apptype)
}

*/
