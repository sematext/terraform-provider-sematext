package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorSolrcloud is the resource class that handles sematext_monitor_solrcloud
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

// resourceMonitorCreateSolrcloud creates the sematext_monitor_solrcloud resource.
func resourceMonitorCreateSolrcloud(d *schema.ResourceData, meta interface{}) error {
	apptype := "SolrCloud"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadSolrcloud reads the sematext_monitor_solrcloud resource from Sematext Cloud.
func resourceMonitorReadSolrcloud(d *schema.ResourceData, meta interface{}) error {
	apptype := "SolrCloud"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateSolrcloud updates Sematext Cloud from the sematext_monitor_solrcloud resource.
func resourceMonitorUpdateSolrcloud(d *schema.ResourceData, meta interface{}) error {
	apptype := "SolrCloud"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteSolrcloud marks a sematext_monitor_solrcloud resource as retired.
func resourceMonitorDeleteSolrcloud(d *schema.ResourceData, meta interface{}) error {
	apptype := "SolrCloud"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsSolrcloud checks a sematext_monitor_solrcloud resource exists in Sematext Cloud.
func resourceMonitorExistsSolrcloud(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "SolrCloud"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportSolrcloud checks a sematext_monitor_solrcloud resource exists in Sematext Cloud.
func resourceSematextMonitorImportSolrcloud(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "SolrCloud"
	return CommonMonitorImport(d, meta, apptype)
}

*/
