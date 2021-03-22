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

// resourceSematextMonitorAkka is the resource class that handles sematext_monitor_akka
func resourceSematextMonitorAkka() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Akka")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateAkka,
		ReadContext:   resourceMonitorReadAkka,
		UpdateContext: resourceMonitorUpdateAkka,
		DeleteContext: resourceMonitorDeleteAkka,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateAkka creates the sematext_monitor_akka resource.
func resourceMonitorCreateAkka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Akka"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadAkka reads the sematext_monitor_akka resource from Sematext Cloud.
func resourceMonitorReadAkka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Akka"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateAkka updates Sematext Cloud from the sematext_monitor_akka resource.
func resourceMonitorUpdateAkka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Akka"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteAkka marks a sematext_monitor_akka resource as retired.
func resourceMonitorDeleteAkka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Akka"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportAkka checks a sematext_monitor_akka resource exists in Sematext Cloud.
func resourceSematextMonitorImportAkka(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Akka"
	return CommonMonitorImport(d, meta, apptype)
}

*/
