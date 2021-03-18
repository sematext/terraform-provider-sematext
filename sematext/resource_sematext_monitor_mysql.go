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

// resourceSematextMonitorMysql is the resource class that handles sematext_monitor_mysql
func resourceSematextMonitorMysql() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("MySQL")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateMysql,
		ReadContext:   resourceMonitorReadMysql,
		UpdateContext: resourceMonitorUpdateMysql,
		DeleteContext: resourceMonitorDeleteMysql,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateMysql creates the sematext_monitor_mysql resource.
func resourceMonitorCreateMysql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "MySQL"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadMysql reads the sematext_monitor_mysql resource from Sematext Cloud.
func resourceMonitorReadMysql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "MySQL"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateMysql updates Sematext Cloud from the sematext_monitor_mysql resource.
func resourceMonitorUpdateMysql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "MySQL"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteMysql marks a sematext_monitor_mysql resource as retired.
func resourceMonitorDeleteMysql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "MySQL"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportMysql checks a sematext_monitor_mysql resource exists in Sematext Cloud.
func resourceSematextMonitorImportMysql(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "MySQL"
	return CommonMonitorImport(d, meta, apptype)
}

*/
