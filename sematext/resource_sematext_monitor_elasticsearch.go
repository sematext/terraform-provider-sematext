package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorElasticsearch is the resource class that handles sematext_monitor_elasticsearch
func resourceSematextMonitorElasticsearch() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Elastic Search")

	return &schema.Resource{
		Create: resourceMonitorCreateElasticsearch,
		Read:   resourceMonitorReadElasticsearch,
		Update: resourceMonitorUpdateElasticsearch,
		Delete: resourceMonitorDeleteElasticsearch,
		Exists: resourceMonitorExistsElasticsearch,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateElasticsearch creates the sematext_monitor_elasticsearch resource.
func resourceMonitorCreateElasticsearch(d *schema.ResourceData, meta interface{}) error {
	apptype := "Elastic Search"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadElasticsearch reads the sematext_monitor_elasticsearch resource from Sematext Cloud.
func resourceMonitorReadElasticsearch(d *schema.ResourceData, meta interface{}) error {
	apptype := "Elastic Search"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateElasticsearch updates Sematext Cloud from the sematext_monitor_elasticsearch resource.
func resourceMonitorUpdateElasticsearch(d *schema.ResourceData, meta interface{}) error {
	apptype := "Elastic Search"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteElasticsearch marks a sematext_monitor_elasticsearch resource as retired.
func resourceMonitorDeleteElasticsearch(d *schema.ResourceData, meta interface{}) error {
	apptype := "Elastic Search"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsElasticsearch checks a sematext_monitor_elasticsearch resource exists in Sematext Cloud.
func resourceMonitorExistsElasticsearch(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Elastic Search"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportElasticsearch checks a sematext_monitor_elasticsearch resource exists in Sematext Cloud.
func resourceSematextMonitorImportElasticsearch(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Elastic Search"
	return CommonMonitorImport(d, meta, apptype)
}

*/
