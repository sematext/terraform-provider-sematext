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

// resourceSematextMonitorKafka is the resource class that handles sematext_monitor_kafka
func resourceSematextMonitorKafka() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Kafka")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateKafka,
		ReadContext:   resourceMonitorReadKafka,
		UpdateContext: resourceMonitorUpdateKafka,
		DeleteContext: resourceMonitorDeleteKafka,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateKafka creates the sematext_monitor_kafka resource.
func resourceMonitorCreateKafka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Kafka"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadKafka reads the sematext_monitor_kafka resource from Sematext Cloud.
func resourceMonitorReadKafka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Kafka"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateKafka updates Sematext Cloud from the sematext_monitor_kafka resource.
func resourceMonitorUpdateKafka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Kafka"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteKafka marks a sematext_monitor_kafka resource as retired.
func resourceMonitorDeleteKafka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Kafka"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportKafka checks a sematext_monitor_kafka resource exists in Sematext Cloud.
func resourceSematextMonitorImportKafka(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Kafka"
	return CommonMonitorImport(d, meta, apptype)
}

*/
