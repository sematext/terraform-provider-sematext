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

// resourceSematextMonitorStorm is the resource class that handles sematext_monitor_storm
func resourceSematextMonitorStorm() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Storm")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateStorm,
		ReadContext:   resourceMonitorReadStorm,
		UpdateContext: resourceMonitorUpdateStorm,
		DeleteContext: resourceMonitorDeleteStorm,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateStorm creates the sematext_monitor_storm resource.
func resourceMonitorCreateStorm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Storm"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadStorm reads the sematext_monitor_storm resource from Sematext Cloud.
func resourceMonitorReadStorm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Storm"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateStorm updates Sematext Cloud from the sematext_monitor_storm resource.
func resourceMonitorUpdateStorm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Storm"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteStorm marks a sematext_monitor_storm resource as retired.
func resourceMonitorDeleteStorm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Storm"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportStorm checks a sematext_monitor_storm resource exists in Sematext Cloud.
func resourceSematextMonitorImportStorm(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Storm"
	return CommonMonitorImport(d, meta, apptype)
}

*/
