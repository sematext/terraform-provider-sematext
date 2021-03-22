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

// resourceSematextMonitorSolr is the resource class that handles sematext_monitor_solr
func resourceSematextMonitorSolr() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Solr")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateSolr,
		ReadContext:   resourceMonitorReadSolr,
		UpdateContext: resourceMonitorUpdateSolr,
		DeleteContext: resourceMonitorDeleteSolr,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateSolr creates the sematext_monitor_solr resource.
func resourceMonitorCreateSolr(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Solr"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadSolr reads the sematext_monitor_solr resource from Sematext Cloud.
func resourceMonitorReadSolr(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Solr"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateSolr updates Sematext Cloud from the sematext_monitor_solr resource.
func resourceMonitorUpdateSolr(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Solr"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteSolr marks a sematext_monitor_solr resource as retired.
func resourceMonitorDeleteSolr(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Solr"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportSolr checks a sematext_monitor_solr resource exists in Sematext Cloud.
func resourceSematextMonitorImportSolr(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Solr"
	return CommonMonitorImport(d, meta, apptype)
}

*/
