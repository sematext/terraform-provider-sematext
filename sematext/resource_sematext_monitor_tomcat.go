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

// resourceSematextMonitorTomcat is the resource class that handles sematext_monitor_tomcat
func resourceSematextMonitorTomcat() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Tomcat")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateTomcat,
		ReadContext:   resourceMonitorReadTomcat,
		UpdateContext: resourceMonitorUpdateTomcat,
		DeleteContext: resourceMonitorDeleteTomcat,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateTomcat creates the sematext_monitor_tomcat resource.
func resourceMonitorCreateTomcat(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Tomcat"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadTomcat reads the sematext_monitor_tomcat resource from Sematext Cloud.
func resourceMonitorReadTomcat(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Tomcat"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateTomcat updates Sematext Cloud from the sematext_monitor_tomcat resource.
func resourceMonitorUpdateTomcat(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Tomcat"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteTomcat marks a sematext_monitor_tomcat resource as retired.
func resourceMonitorDeleteTomcat(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Tomcat"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportTomcat checks a sematext_monitor_tomcat resource exists in Sematext Cloud.
func resourceSematextMonitorImportTomcat(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Tomcat"
	return CommonMonitorImport(d, meta, apptype)
}

*/
