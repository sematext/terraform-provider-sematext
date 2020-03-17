package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorSearchanalytics TODO Doc Comment
func resourceSematextMonitorSearchanalytics() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateSearchanalytics,
		Read:   resourceMonitorReadSearchanalytics,
		Update: resourceMonitorUpdateSearchanalytics,
		Delete: resourceMonitorDeleteSearchanalytics,
		Exists: resourceMonitorExistsSearchanalytics,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateSearchanalytics TODO Doc Comment
func resourceMonitorCreateSearchanalytics(d *schema.ResourceData, meta interface{}) error {
	apptype := "SearchAnalytics"
	return CommonMonitorCreate(d, meta, apptype)
}

// resourceMonitorReadSearchanalytics TODO Doc Comment
func resourceMonitorReadSearchanalytics(d *schema.ResourceData, meta interface{}) error {
	apptype := "SearchAnalytics"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateSearchanalytics TODO Doc Comment
func resourceMonitorUpdateSearchanalytics(d *schema.ResourceData, meta interface{}) error {
	apptype := "SearchAnalytics"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteSearchanalytics TODO Doc Comment
func resourceMonitorDeleteSearchanalytics(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "SearchAnalytics"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsSearchanalytics TODO Doc Comment
func resourceMonitorExistsSearchanalytics(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "SearchAnalytics"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportSearchanalytics TODO Doc Comment
func resourceSematextMonitorImportSearchanalytics(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "SearchAnalytics"
	return CommonMonitorImport(d, meta, apptype)
}
