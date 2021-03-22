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

// resourceSematextMonitorClickhouse is the resource class that handles sematext_monitor_clickhouse
func resourceSematextMonitorClickhouse() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("ClickHouse")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateClickhouse,
		ReadContext:   resourceMonitorReadClickhouse,
		UpdateContext: resourceMonitorUpdateClickhouse,
		DeleteContext: resourceMonitorDeleteClickhouse,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateClickhouse creates the sematext_monitor_clickhouse resource.
func resourceMonitorCreateClickhouse(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "ClickHouse"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadClickhouse reads the sematext_monitor_clickhouse resource from Sematext Cloud.
func resourceMonitorReadClickhouse(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "ClickHouse"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateClickhouse updates Sematext Cloud from the sematext_monitor_clickhouse resource.
func resourceMonitorUpdateClickhouse(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "ClickHouse"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteClickhouse marks a sematext_monitor_clickhouse resource as retired.
func resourceMonitorDeleteClickhouse(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "ClickHouse"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportClickhouse checks a sematext_monitor_clickhouse resource exists in Sematext Cloud.
func resourceSematextMonitorImportClickhouse(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "ClickHouse"
	return CommonMonitorImport(d, meta, apptype)
}

*/
