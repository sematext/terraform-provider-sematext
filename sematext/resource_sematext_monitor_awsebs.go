package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorAwsebs TODO Doc Comment
func resourceSematextMonitorAwsebs() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("AWS EBS")

	return &schema.Resource{
		Create: resourceMonitorCreateAwsebs,
		Read:   resourceMonitorReadAwsebs,
		Update: resourceMonitorUpdateAwsebs,
		Delete: resourceMonitorDeleteAwsebs,
		Exists: resourceMonitorExistsAwsebs,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateAwsebs TODO Doc Comment
func resourceMonitorCreateAwsebs(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS EBS"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadAwsebs TODO Doc Comment
func resourceMonitorReadAwsebs(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS EBS"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateAwsebs TODO Doc Comment
func resourceMonitorUpdateAwsebs(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS EBS"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteAwsebs TODO Doc Comment
func resourceMonitorDeleteAwsebs(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "AWS EBS"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsAwsebs TODO Doc Comment
func resourceMonitorExistsAwsebs(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "AWS EBS"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportAwsebs TODO Doc Comment
func resourceSematextMonitorImportAwsebs(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "AWS EBS"
	return CommonMonitorImport(d, meta, apptype)
}
