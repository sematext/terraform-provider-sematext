package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorSolr TODO Doc Comment
func resourceSematextMonitorSolr() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateSolr,
		Read:   resourceMonitorReadSolr,
		Update: resourceMonitorUpdateSolr,
		Delete: resourceMonitorDeleteSolr,
		Exists: resourceMonitorExistsSolr,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateSolr TODO Doc Comment
func resourceMonitorCreateSolr(d *schema.ResourceData, meta interface{}) error {
	apptype := "Solr"
	return CommonMonitorCreate(d, meta, apptype)
}

// resourceMonitorReadSolr TODO Doc Comment
func resourceMonitorReadSolr(d *schema.ResourceData, meta interface{}) error {
	apptype := "Solr"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateSolr TODO Doc Comment
func resourceMonitorUpdateSolr(d *schema.ResourceData, meta interface{}) error {
	apptype := "Solr"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteSolr TODO Doc Comment
func resourceMonitorDeleteSolr(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Solr"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsSolr TODO Doc Comment
func resourceMonitorExistsSolr(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Solr"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportSolr TODO Doc Comment
func resourceSematextMonitorImportSolr(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Solr"
	return CommonMonitorImport(d, meta, apptype)
}
