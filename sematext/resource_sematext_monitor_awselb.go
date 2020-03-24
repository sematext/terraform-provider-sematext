package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorAwselb TODO Doc Comment
func resourceSematextMonitorAwselb() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Awselb")

	return &schema.Resource{
		Create: resourceMonitorCreateAwselb,
		Read:   resourceMonitorReadAwselb,
		Update: resourceMonitorUpdateAwselb,
		Delete: resourceMonitorDeleteAwselb,
		Exists: resourceMonitorExistsAwselb,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateAwselb TODO Doc Comment
func resourceMonitorCreateAwselb(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS ELB"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadAwselb TODO Doc Comment
func resourceMonitorReadAwselb(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS ELB"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateAwselb TODO Doc Comment
func resourceMonitorUpdateAwselb(d *schema.ResourceData, meta interface{}) error {
	apptype := "AWS ELB"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteAwselb TODO Doc Comment
func resourceMonitorDeleteAwselb(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "AWS ELB"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsAwselb TODO Doc Comment
func resourceMonitorExistsAwselb(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "AWS ELB"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportAwselb TODO Doc Comment
func resourceSematextMonitorImportAwselb(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "AWS ELB"
	return CommonMonitorImport(d, meta, apptype)
}
