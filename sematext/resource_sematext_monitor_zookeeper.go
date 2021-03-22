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

// resourceSematextMonitorZookeeper is the resource class that handles sematext_monitor_zookeeper
func resourceSematextMonitorZookeeper() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("ZooKeeper")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateZookeeper,
		ReadContext:   resourceMonitorReadZookeeper,
		UpdateContext: resourceMonitorUpdateZookeeper,
		DeleteContext: resourceMonitorDeleteZookeeper,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateZookeeper creates the sematext_monitor_zookeeper resource.
func resourceMonitorCreateZookeeper(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "ZooKeeper"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadZookeeper reads the sematext_monitor_zookeeper resource from Sematext Cloud.
func resourceMonitorReadZookeeper(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "ZooKeeper"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateZookeeper updates Sematext Cloud from the sematext_monitor_zookeeper resource.
func resourceMonitorUpdateZookeeper(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "ZooKeeper"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteZookeeper marks a sematext_monitor_zookeeper resource as retired.
func resourceMonitorDeleteZookeeper(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "ZooKeeper"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportZookeeper checks a sematext_monitor_zookeeper resource exists in Sematext Cloud.
func resourceSematextMonitorImportZookeeper(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "ZooKeeper"
	return CommonMonitorImport(d, meta, apptype)
}

*/
