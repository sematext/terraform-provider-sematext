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

// resourceSematextMonitorMongodb is the resource class that handles sematext_monitor_mongodb
func resourceSematextMonitorMongodb() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("MongoDB")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateMongodb,
		ReadContext:   resourceMonitorReadMongodb,
		UpdateContext: resourceMonitorUpdateMongodb,
		DeleteContext: resourceMonitorDeleteMongodb,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateMongodb creates the sematext_monitor_mongodb resource.
func resourceMonitorCreateMongodb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "MongoDB"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadMongodb reads the sematext_monitor_mongodb resource from Sematext Cloud.
func resourceMonitorReadMongodb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "MongoDB"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateMongodb updates Sematext Cloud from the sematext_monitor_mongodb resource.
func resourceMonitorUpdateMongodb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "MongoDB"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteMongodb marks a sematext_monitor_mongodb resource as retired.
func resourceMonitorDeleteMongodb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "MongoDB"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportMongodb checks a sematext_monitor_mongodb resource exists in Sematext Cloud.
func resourceSematextMonitorImportMongodb(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "MongoDB"
	return CommonMonitorImport(d, meta, apptype)
}

*/
