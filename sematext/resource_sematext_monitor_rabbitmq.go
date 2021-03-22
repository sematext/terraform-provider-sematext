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

// resourceSematextMonitorRabbitmq is the resource class that handles sematext_monitor_rabbitmq
func resourceSematextMonitorRabbitmq() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("rabbitmq")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateRabbitmq,
		ReadContext:   resourceMonitorReadRabbitmq,
		UpdateContext: resourceMonitorUpdateRabbitmq,
		DeleteContext: resourceMonitorDeleteRabbitmq,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateRabbitmq creates the sematext_monitor_rabbitmq resource.
func resourceMonitorCreateRabbitmq(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "rabbitmq"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadRabbitmq reads the sematext_monitor_rabbitmq resource from Sematext Cloud.
func resourceMonitorReadRabbitmq(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "rabbitmq"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateRabbitmq updates Sematext Cloud from the sematext_monitor_rabbitmq resource.
func resourceMonitorUpdateRabbitmq(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "rabbitmq"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteRabbitmq marks a sematext_monitor_rabbitmq resource as retired.
func resourceMonitorDeleteRabbitmq(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "rabbitmq"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportRabbitmq checks a sematext_monitor_rabbitmq resource exists in Sematext Cloud.
func resourceSematextMonitorImportRabbitmq(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "rabbitmq"
	return CommonMonitorImport(d, meta, apptype)
}

*/
