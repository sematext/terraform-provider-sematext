package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorSolr is the resource class that handles sematext_monitor_solr
func resourceSematextMonitorSolr() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Solr")

	return &schema.Resource{
		Create: resourceMonitorCreateSolr,
		Read:   resourceMonitorReadSolr,
		Update: resourceMonitorUpdateSolr,
		Delete: resourceMonitorDeleteSolr,
		Exists: resourceMonitorExistsSolr,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateSolr creates the sematext_monitor_solr resource.
func resourceMonitorCreateSolr(d *schema.ResourceData, meta interface{}) error {
	apptype := "Solr"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadSolr reads the sematext_monitor_solr resource from Sematext Cloud.
func resourceMonitorReadSolr(d *schema.ResourceData, meta interface{}) error {
	apptype := "Solr"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateSolr updates Sematext Cloud from the sematext_monitor_solr resource.
func resourceMonitorUpdateSolr(d *schema.ResourceData, meta interface{}) error {
	apptype := "Solr"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteSolr marks a sematext_monitor_solr resource as retired.
func resourceMonitorDeleteSolr(d *schema.ResourceData, meta interface{}) error {
	apptype := "Solr"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsSolr checks a sematext_monitor_solr resource exists in Sematext Cloud.
func resourceMonitorExistsSolr(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Solr"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportSolr checks a sematext_monitor_solr resource exists in Sematext Cloud.
func resourceSematextMonitorImportSolr(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Solr"
	return CommonMonitorImport(d, meta, apptype)
}

*/
