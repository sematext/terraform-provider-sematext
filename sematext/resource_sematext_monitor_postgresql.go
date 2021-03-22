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

// resourceSematextMonitorPostgresql is the resource class that handles sematext_monitor_postgresql
func resourceSematextMonitorPostgresql() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("postgresql")

	return &schema.Resource{
		CreateContext: resourceMonitorCreatePostgresql,
		ReadContext:   resourceMonitorReadPostgresql,
		UpdateContext: resourceMonitorUpdatePostgresql,
		DeleteContext: resourceMonitorDeletePostgresql,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreatePostgresql creates the sematext_monitor_postgresql resource.
func resourceMonitorCreatePostgresql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "postgresql"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadPostgresql reads the sematext_monitor_postgresql resource from Sematext Cloud.
func resourceMonitorReadPostgresql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "postgresql"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdatePostgresql updates Sematext Cloud from the sematext_monitor_postgresql resource.
func resourceMonitorUpdatePostgresql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "postgresql"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeletePostgresql marks a sematext_monitor_postgresql resource as retired.
func resourceMonitorDeletePostgresql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "postgresql"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportPostgresql checks a sematext_monitor_postgresql resource exists in Sematext Cloud.
func resourceSematextMonitorImportPostgresql(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "postgresql"
	return CommonMonitorImport(d, meta, apptype)
}

*/
