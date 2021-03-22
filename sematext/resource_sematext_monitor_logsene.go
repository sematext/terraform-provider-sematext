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

// resourceSematextMonitorLogsene is the resource class that handles sematext_monitor_logsene
func resourceSematextMonitorLogsene() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Logsene")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateLogsene,
		ReadContext:   resourceMonitorReadLogsene,
		UpdateContext: resourceMonitorUpdateLogsene,
		DeleteContext: resourceMonitorDeleteLogsene,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateLogsene creates the sematext_monitor_logsene resource.
func resourceMonitorCreateLogsene(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Logsene"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadLogsene reads the sematext_monitor_logsene resource from Sematext Cloud.
func resourceMonitorReadLogsene(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Logsene"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateLogsene updates Sematext Cloud from the sematext_monitor_logsene resource.
func resourceMonitorUpdateLogsene(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Logsene"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteLogsene marks a sematext_monitor_logsene resource as retired.
func resourceMonitorDeleteLogsene(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Logsene"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportLogsene checks a sematext_monitor_logsene resource exists in Sematext Cloud.
func resourceSematextMonitorImportLogsene(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Logsene"
	return CommonMonitorImport(d, meta, apptype)
}

*/
