package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorAwsec2 is the resource class that handles sematext_monitor_awsec2
func resourceSematextMonitorAwsec2() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("AWS EC2")

	return &schema.Resource{
		Create: resourceMonitorCreateAwsec2,
		Read:   resourceMonitorReadAwsec2,
		Update: resourceMonitorUpdateAwsec2,
		Delete: resourceMonitorDeleteAwsec2,
		Exists: resourceMonitorExistsAwsec2,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateAwsec2 creates the sematext_monitor_awsec2 resource.
func resourceMonitorCreateAwsec2(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS EC2"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadAwsec2 reads the sematext_monitor_awsec2 resource from Sematext Cloud.
func resourceMonitorReadAwsec2(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS EC2"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateAwsec2 updates Sematext Cloud from the sematext_monitor_awsec2 resource.
func resourceMonitorUpdateAwsec2(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS EC2"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteAwsec2 marks a sematext_monitor_awsec2 resource as retired.
func resourceMonitorDeleteAwsec2(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS EC2"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsAwsec2 checks a sematext_monitor_awsec2 resource exists in Sematext Cloud.
func resourceMonitorExistsAwsec2(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "AWS EC2"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportAwsec2 checks a sematext_monitor_awsec2 resource exists in Sematext Cloud.
func resourceSematextMonitorImportAwsec2(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "AWS EC2"
	return CommonMonitorImport(d, meta, apptype)
}

*/
