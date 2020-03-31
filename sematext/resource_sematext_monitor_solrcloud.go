package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorSolrcloud TODO Doc Comment
func resourceSematextMonitorSolrcloud() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("SolrCloud")

	return &schema.Resource{
		Create: resourceMonitorCreateSolrcloud,
		Read:   resourceMonitorReadSolrcloud,
		Update: resourceMonitorUpdateSolrcloud,
		Delete: resourceMonitorDeleteSolrcloud,
		Exists: resourceMonitorExistsSolrcloud,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateSolrcloud TODO Doc Comment
func resourceMonitorCreateSolrcloud(d *schema.ResourceData, meta interface{}) error {
	apptype := "SolrCloud"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadSolrcloud TODO Doc Comment
func resourceMonitorReadSolrcloud(d *schema.ResourceData, meta interface{}) error {
	apptype := "SolrCloud"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateSolrcloud TODO Doc Comment
func resourceMonitorUpdateSolrcloud(d *schema.ResourceData, meta interface{}) error {
	apptype := "SolrCloud"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteSolrcloud TODO Doc Comment
func resourceMonitorDeleteSolrcloud(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "SolrCloud"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsSolrcloud TODO Doc Comment
func resourceMonitorExistsSolrcloud(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "SolrCloud"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportSolrcloud TODO Doc Comment
func resourceSematextMonitorImportSolrcloud(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "SolrCloud"
	return CommonMonitorImport(d, meta, apptype)
}
