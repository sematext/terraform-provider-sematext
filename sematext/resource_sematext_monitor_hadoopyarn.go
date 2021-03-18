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

// resourceSematextMonitorHadoopyarn is the resource class that handles sematext_monitor_hadoopyarn
func resourceSematextMonitorHadoopyarn() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Hadoop-YARN")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateHadoopyarn,
		ReadContext:   resourceMonitorReadHadoopyarn,
		UpdateContext: resourceMonitorUpdateHadoopyarn,
		DeleteContext: resourceMonitorDeleteHadoopyarn,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateHadoopyarn creates the sematext_monitor_hadoopyarn resource.
func resourceMonitorCreateHadoopyarn(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Hadoop-YARN"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadHadoopyarn reads the sematext_monitor_hadoopyarn resource from Sematext Cloud.
func resourceMonitorReadHadoopyarn(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Hadoop-YARN"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateHadoopyarn updates Sematext Cloud from the sematext_monitor_hadoopyarn resource.
func resourceMonitorUpdateHadoopyarn(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Hadoop-YARN"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteHadoopyarn marks a sematext_monitor_hadoopyarn resource as retired.
func resourceMonitorDeleteHadoopyarn(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Hadoop-YARN"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportHadoopyarn checks a sematext_monitor_hadoopyarn resource exists in Sematext Cloud.
func resourceSematextMonitorImportHadoopyarn(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Hadoop-YARN"
	return CommonMonitorImport(d, meta, apptype)
}

*/
