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

// resourceSematextMonitorInfra is the resource class that handles sematext_monitor_infra
func resourceSematextMonitorInfra() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Infra")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateInfra,
		ReadContext:   resourceMonitorReadInfra,
		UpdateContext: resourceMonitorUpdateInfra,
		DeleteContext: resourceMonitorDeleteInfra,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateInfra creates the sematext_monitor_infra resource.
func resourceMonitorCreateInfra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Infra"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadInfra reads the sematext_monitor_infra resource from Sematext Cloud.
func resourceMonitorReadInfra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Infra"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateInfra updates Sematext Cloud from the sematext_monitor_infra resource.
func resourceMonitorUpdateInfra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Infra"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteInfra marks a sematext_monitor_infra resource as retired.
func resourceMonitorDeleteInfra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Infra"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportInfra checks a sematext_monitor_infra resource exists in Sematext Cloud.
func resourceSematextMonitorImportInfra(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Infra"
	return CommonMonitorImport(d, meta, apptype)
}

*/
