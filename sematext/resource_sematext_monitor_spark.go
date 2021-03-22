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

// resourceSematextMonitorSpark is the resource class that handles sematext_monitor_spark
func resourceSematextMonitorSpark() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Spark")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateSpark,
		ReadContext:   resourceMonitorReadSpark,
		UpdateContext: resourceMonitorUpdateSpark,
		DeleteContext: resourceMonitorDeleteSpark,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateSpark creates the sematext_monitor_spark resource.
func resourceMonitorCreateSpark(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Spark"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadSpark reads the sematext_monitor_spark resource from Sematext Cloud.
func resourceMonitorReadSpark(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Spark"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateSpark updates Sematext Cloud from the sematext_monitor_spark resource.
func resourceMonitorUpdateSpark(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Spark"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteSpark marks a sematext_monitor_spark resource as retired.
func resourceMonitorDeleteSpark(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Spark"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportSpark checks a sematext_monitor_spark resource exists in Sematext Cloud.
func resourceSematextMonitorImportSpark(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Spark"
	return CommonMonitorImport(d, meta, apptype)
}

*/
