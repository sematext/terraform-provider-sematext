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

// resourceSematextMonitorNginx is the resource class that handles sematext_monitor_nginx
func resourceSematextMonitorNginx() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Nginx")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateNginx,
		ReadContext:   resourceMonitorReadNginx,
		UpdateContext: resourceMonitorUpdateNginx,
		DeleteContext: resourceMonitorDeleteNginx,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateNginx creates the sematext_monitor_nginx resource.
func resourceMonitorCreateNginx(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Nginx"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadNginx reads the sematext_monitor_nginx resource from Sematext Cloud.
func resourceMonitorReadNginx(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Nginx"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateNginx updates Sematext Cloud from the sematext_monitor_nginx resource.
func resourceMonitorUpdateNginx(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Nginx"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteNginx marks a sematext_monitor_nginx resource as retired.
func resourceMonitorDeleteNginx(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Nginx"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportNginx checks a sematext_monitor_nginx resource exists in Sematext Cloud.
func resourceSematextMonitorImportNginx(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Nginx"
	return CommonMonitorImport(d, meta, apptype)
}

*/
