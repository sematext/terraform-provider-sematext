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

// resourceSematextMonitorApache is the resource class that handles sematext_monitor_apache
func resourceSematextMonitorApache() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Apache")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateApache,
		ReadContext:   resourceMonitorReadApache,
		UpdateContext: resourceMonitorUpdateApache,
		DeleteContext: resourceMonitorDeleteApache,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateApache creates the sematext_monitor_apache resource.
func resourceMonitorCreateApache(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Apache"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadApache reads the sematext_monitor_apache resource from Sematext Cloud.
func resourceMonitorReadApache(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Apache"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateApache updates Sematext Cloud from the sematext_monitor_apache resource.
func resourceMonitorUpdateApache(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Apache"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteApache marks a sematext_monitor_apache resource as retired.
func resourceMonitorDeleteApache(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Apache"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportApache checks a sematext_monitor_apache resource exists in Sematext Cloud.
func resourceSematextMonitorImportApache(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Apache"
	return CommonMonitorImport(d, meta, apptype)
}

*/
