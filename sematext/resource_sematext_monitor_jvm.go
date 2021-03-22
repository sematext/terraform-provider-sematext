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

// resourceSematextMonitorJvm is the resource class that handles sematext_monitor_jvm
func resourceSematextMonitorJvm() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("JVM")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateJvm,
		ReadContext:   resourceMonitorReadJvm,
		UpdateContext: resourceMonitorUpdateJvm,
		DeleteContext: resourceMonitorDeleteJvm,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateJvm creates the sematext_monitor_jvm resource.
func resourceMonitorCreateJvm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "JVM"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadJvm reads the sematext_monitor_jvm resource from Sematext Cloud.
func resourceMonitorReadJvm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "JVM"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateJvm updates Sematext Cloud from the sematext_monitor_jvm resource.
func resourceMonitorUpdateJvm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "JVM"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteJvm marks a sematext_monitor_jvm resource as retired.
func resourceMonitorDeleteJvm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "JVM"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportJvm checks a sematext_monitor_jvm resource exists in Sematext Cloud.
func resourceSematextMonitorImportJvm(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "JVM"
	return CommonMonitorImport(d, meta, apptype)
}

*/
