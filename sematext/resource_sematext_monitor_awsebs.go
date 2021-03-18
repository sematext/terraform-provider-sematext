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

// resourceSematextMonitorAwsebs is the resource class that handles sematext_monitor_awsebs
func resourceSematextMonitorAwsebs() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("AWS EBS")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateAwsebs,
		ReadContext:   resourceMonitorReadAwsebs,
		UpdateContext: resourceMonitorUpdateAwsebs,
		DeleteContext: resourceMonitorDeleteAwsebs,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateAwsebs creates the sematext_monitor_awsebs resource.
func resourceMonitorCreateAwsebs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS EBS"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadAwsebs reads the sematext_monitor_awsebs resource from Sematext Cloud.
func resourceMonitorReadAwsebs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS EBS"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateAwsebs updates Sematext Cloud from the sematext_monitor_awsebs resource.
func resourceMonitorUpdateAwsebs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS EBS"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteAwsebs marks a sematext_monitor_awsebs resource as retired.
func resourceMonitorDeleteAwsebs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS EBS"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportAwsebs checks a sematext_monitor_awsebs resource exists in Sematext Cloud.
func resourceSematextMonitorImportAwsebs(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "AWS EBS"
	return CommonMonitorImport(d, meta, apptype)
}

*/
