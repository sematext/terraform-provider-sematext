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

// resourceSematextMonitorAwselb is the resource class that handles sematext_monitor_awselb
func resourceSematextMonitorAwselb() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("AWS ELB")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateAwselb,
		ReadContext:   resourceMonitorReadAwselb,
		UpdateContext: resourceMonitorUpdateAwselb,
		DeleteContext: resourceMonitorDeleteAwselb,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateAwselb creates the sematext_monitor_awselb resource.
func resourceMonitorCreateAwselb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS ELB"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadAwselb reads the sematext_monitor_awselb resource from Sematext Cloud.
func resourceMonitorReadAwselb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS ELB"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateAwselb updates Sematext Cloud from the sematext_monitor_awselb resource.
func resourceMonitorUpdateAwselb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS ELB"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteAwselb marks a sematext_monitor_awselb resource as retired.
func resourceMonitorDeleteAwselb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS ELB"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportAwselb checks a sematext_monitor_awselb resource exists in Sematext Cloud.
func resourceSematextMonitorImportAwselb(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "AWS ELB"
	return CommonMonitorImport(d, meta, apptype)
}

*/
