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

// resourceSematextMonitorHaproxy is the resource class that handles sematext_monitor_haproxy
func resourceSematextMonitorHaproxy() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("HAProxy")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateHaproxy,
		ReadContext:   resourceMonitorReadHaproxy,
		UpdateContext: resourceMonitorUpdateHaproxy,
		DeleteContext: resourceMonitorDeleteHaproxy,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateHaproxy creates the sematext_monitor_haproxy resource.
func resourceMonitorCreateHaproxy(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "HAProxy"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadHaproxy reads the sematext_monitor_haproxy resource from Sematext Cloud.
func resourceMonitorReadHaproxy(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "HAProxy"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateHaproxy updates Sematext Cloud from the sematext_monitor_haproxy resource.
func resourceMonitorUpdateHaproxy(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "HAProxy"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteHaproxy marks a sematext_monitor_haproxy resource as retired.
func resourceMonitorDeleteHaproxy(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "HAProxy"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportHaproxy checks a sematext_monitor_haproxy resource exists in Sematext Cloud.
func resourceSematextMonitorImportHaproxy(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "HAProxy"
	return CommonMonitorImport(d, meta, apptype)
}

*/
