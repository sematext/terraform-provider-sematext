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

// resourceSematextMonitorNginxplus is the resource class that handles sematext_monitor_nginxplus
func resourceSematextMonitorNginxplus() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Nginx-Plus")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateNginxplus,
		ReadContext:   resourceMonitorReadNginxplus,
		UpdateContext: resourceMonitorUpdateNginxplus,
		DeleteContext: resourceMonitorDeleteNginxplus,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateNginxplus creates the sematext_monitor_nginxplus resource.
func resourceMonitorCreateNginxplus(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Nginx-Plus"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadNginxplus reads the sematext_monitor_nginxplus resource from Sematext Cloud.
func resourceMonitorReadNginxplus(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Nginx-Plus"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateNginxplus updates Sematext Cloud from the sematext_monitor_nginxplus resource.
func resourceMonitorUpdateNginxplus(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Nginx-Plus"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteNginxplus marks a sematext_monitor_nginxplus resource as retired.
func resourceMonitorDeleteNginxplus(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Nginx-Plus"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportNginxplus checks a sematext_monitor_nginxplus resource exists in Sematext Cloud.
func resourceSematextMonitorImportNginxplus(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Nginx-Plus"
	return CommonMonitorImport(d, meta, apptype)
}

*/
