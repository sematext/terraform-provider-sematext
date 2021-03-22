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

// resourceSematextMonitorNodejs is the resource class that handles sematext_monitor_nodejs
func resourceSematextMonitorNodejs() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Node.js")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateNodejs,
		ReadContext:   resourceMonitorReadNodejs,
		UpdateContext: resourceMonitorUpdateNodejs,
		DeleteContext: resourceMonitorDeleteNodejs,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateNodejs creates the sematext_monitor_nodejs resource.
func resourceMonitorCreateNodejs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Node.js"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadNodejs reads the sematext_monitor_nodejs resource from Sematext Cloud.
func resourceMonitorReadNodejs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Node.js"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateNodejs updates Sematext Cloud from the sematext_monitor_nodejs resource.
func resourceMonitorUpdateNodejs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Node.js"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteNodejs marks a sematext_monitor_nodejs resource as retired.
func resourceMonitorDeleteNodejs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Node.js"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportNodejs checks a sematext_monitor_nodejs resource exists in Sematext Cloud.
func resourceSematextMonitorImportNodejs(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Node.js"
	return CommonMonitorImport(d, meta, apptype)
}

*/
