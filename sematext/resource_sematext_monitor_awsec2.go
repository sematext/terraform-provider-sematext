package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorAwsec2 TODO Doc Comment
func resourceSematextMonitorAwsec2() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateAwsec2,
		Read:   resourceMonitorReadAwsec2,
		Update: resourceMonitorUpdateAwsec2,
		Delete: resourceMonitorDeleteAwsec2,
		Exists: resourceMonitorExistsAwsec2,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateAwsec2 TODO Doc Comment
func resourceMonitorCreateAwsec2(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS EC2"
	return CommonMonitorCreate(d, meta, apptype)
}

// resourceMonitorReadAwsec2 TODO Doc Comment
func resourceMonitorReadAwsec2(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS EC2"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateAwsec2 TODO Doc Comment
func resourceMonitorUpdateAwsec2(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS EC2"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteAwsec2 TODO Doc Comment
func resourceMonitorDeleteAwsec2(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "AWS EC2"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsAwsec2 TODO Doc Comment
func resourceMonitorExistsAwsec2(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "AWS EC2"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportAwsec2 TODO Doc Comment
func resourceSematextMonitorImportAwsec2(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "AWS EC2"
	return CommonMonitorImport(d, meta, apptype)
}
