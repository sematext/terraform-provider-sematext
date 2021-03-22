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

// resourceSematextMonitorSolrcloud is the resource class that handles sematext_monitor_solrcloud
func resourceSematextMonitorSolrcloud() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("SolrCloud")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateSolrcloud,
		ReadContext:   resourceMonitorReadSolrcloud,
		UpdateContext: resourceMonitorUpdateSolrcloud,
		DeleteContext: resourceMonitorDeleteSolrcloud,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateSolrcloud creates the sematext_monitor_solrcloud resource.
func resourceMonitorCreateSolrcloud(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "SolrCloud"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadSolrcloud reads the sematext_monitor_solrcloud resource from Sematext Cloud.
func resourceMonitorReadSolrcloud(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "SolrCloud"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateSolrcloud updates Sematext Cloud from the sematext_monitor_solrcloud resource.
func resourceMonitorUpdateSolrcloud(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "SolrCloud"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteSolrcloud marks a sematext_monitor_solrcloud resource as retired.
func resourceMonitorDeleteSolrcloud(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "SolrCloud"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportSolrcloud checks a sematext_monitor_solrcloud resource exists in Sematext Cloud.
func resourceSematextMonitorImportSolrcloud(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "SolrCloud"
	return CommonMonitorImport(d, meta, apptype)
}

*/
