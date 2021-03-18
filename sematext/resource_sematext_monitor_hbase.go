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

// resourceSematextMonitorHbase is the resource class that handles sematext_monitor_hbase
func resourceSematextMonitorHbase() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("HBase")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateHbase,
		ReadContext:   resourceMonitorReadHbase,
		UpdateContext: resourceMonitorUpdateHbase,
		DeleteContext: resourceMonitorDeleteHbase,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateHbase creates the sematext_monitor_hbase resource.
func resourceMonitorCreateHbase(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "HBase"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadHbase reads the sematext_monitor_hbase resource from Sematext Cloud.
func resourceMonitorReadHbase(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "HBase"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateHbase updates Sematext Cloud from the sematext_monitor_hbase resource.
func resourceMonitorUpdateHbase(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "HBase"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteHbase marks a sematext_monitor_hbase resource as retired.
func resourceMonitorDeleteHbase(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "HBase"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportHbase checks a sematext_monitor_hbase resource exists in Sematext Cloud.
func resourceSematextMonitorImportHbase(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "HBase"
	return CommonMonitorImport(d, meta, apptype)
}

*/
