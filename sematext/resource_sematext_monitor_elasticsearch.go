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

// resourceSematextMonitorElasticsearch is the resource class that handles sematext_monitor_elasticsearch
func resourceSematextMonitorElasticsearch() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Elastic Search")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateElasticsearch,
		ReadContext:   resourceMonitorReadElasticsearch,
		UpdateContext: resourceMonitorUpdateElasticsearch,
		DeleteContext: resourceMonitorDeleteElasticsearch,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateElasticsearch creates the sematext_monitor_elasticsearch resource.
func resourceMonitorCreateElasticsearch(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Elastic Search"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadElasticsearch reads the sematext_monitor_elasticsearch resource from Sematext Cloud.
func resourceMonitorReadElasticsearch(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Elastic Search"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateElasticsearch updates Sematext Cloud from the sematext_monitor_elasticsearch resource.
func resourceMonitorUpdateElasticsearch(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Elastic Search"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteElasticsearch marks a sematext_monitor_elasticsearch resource as retired.
func resourceMonitorDeleteElasticsearch(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Elastic Search"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportElasticsearch checks a sematext_monitor_elasticsearch resource exists in Sematext Cloud.
func resourceSematextMonitorImportElasticsearch(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Elastic Search"
	return CommonMonitorImport(d, meta, apptype)
}

*/
