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

// resourceSematextMonitorRedis is the resource class that handles sematext_monitor_redis
func resourceSematextMonitorRedis() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Redis")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateRedis,
		ReadContext:   resourceMonitorReadRedis,
		UpdateContext: resourceMonitorUpdateRedis,
		DeleteContext: resourceMonitorDeleteRedis,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateRedis creates the sematext_monitor_redis resource.
func resourceMonitorCreateRedis(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Redis"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadRedis reads the sematext_monitor_redis resource from Sematext Cloud.
func resourceMonitorReadRedis(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Redis"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateRedis updates Sematext Cloud from the sematext_monitor_redis resource.
func resourceMonitorUpdateRedis(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Redis"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteRedis marks a sematext_monitor_redis resource as retired.
func resourceMonitorDeleteRedis(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Redis"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportRedis checks a sematext_monitor_redis resource exists in Sematext Cloud.
func resourceSematextMonitorImportRedis(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Redis"
	return CommonMonitorImport(d, meta, apptype)
}

*/
