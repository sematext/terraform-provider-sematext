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

// resourceSematextMonitorAwsec2 is the resource class that handles sematext_monitor_awsec2
func resourceSematextMonitorAwsec2() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("AWS EC2")

	return &schema.Resource{
		CreateContext: resourceMonitorCreateAwsec2,
		ReadContext:   resourceMonitorReadAwsec2,
		UpdateContext: resourceMonitorUpdateAwsec2,
		DeleteContext: resourceMonitorDeleteAwsec2,
		Schema:        fieldSchema,
	}
}

// resourceMonitorCreateAwsec2 creates the sematext_monitor_awsec2 resource.
func resourceMonitorCreateAwsec2(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS EC2"
	err := CommonMonitorCreate(ctx, d, meta, apptype)

	return err
}

// resourceMonitorReadAwsec2 reads the sematext_monitor_awsec2 resource from Sematext Cloud.
func resourceMonitorReadAwsec2(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS EC2"
	return CommonMonitorRead(ctx, d, meta, apptype)
}

// resourceMonitorUpdateAwsec2 updates Sematext Cloud from the sematext_monitor_awsec2 resource.
func resourceMonitorUpdateAwsec2(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS EC2"
	return CommonMonitorUpdate(ctx, d, meta, apptype)
}

// resourceMonitorDeleteAwsec2 marks a sematext_monitor_awsec2 resource as retired.
func resourceMonitorDeleteAwsec2(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS EC2"
	return CommonMonitorDelete(ctx, d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportAwsec2 checks a sematext_monitor_awsec2 resource exists in Sematext Cloud.
func resourceSematextMonitorImportAwsec2(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "AWS EC2"
	return CommonMonitorImport(d, meta, apptype)
}

*/
