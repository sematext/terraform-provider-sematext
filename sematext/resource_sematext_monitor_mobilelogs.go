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

// resourceSematextMonitorMobilelogs is the resource class that handles sematext_monitor_mobilelogs
func resourceSematextMonitorMobilelogs() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("mobile-logs")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateMobilelogs,
		ReadContext:   resourceMonitorReadMobilelogs,
		UpdateContext: resourceMonitorUpdateMobilelogs,
		DeleteContext: resourceMonitorDeleteMobilelogs,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateMobilelogs creates the sematext_monitor_mobilelogs resource.
func resourceMonitorCreateMobilelogs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "mobile-logs"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadMobilelogs reads the sematext_monitor_mobilelogs resource from Sematext Cloud.
func resourceMonitorReadMobilelogs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "mobile-logs"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateMobilelogs updates Sematext Cloud from the sematext_monitor_mobilelogs resource.
func resourceMonitorUpdateMobilelogs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "mobile-logs"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteMobilelogs marks a sematext_monitor_mobilelogs resource as retired.
func resourceMonitorDeleteMobilelogs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "mobile-logs"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportMobilelogs checks a sematext_monitor_mobilelogs resource exists in Sematext Cloud.
func resourceSematextMonitorImportMobilelogs(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "mobile-logs"
	return CommonMonitorImport(d, meta, apptype)
}

*/
