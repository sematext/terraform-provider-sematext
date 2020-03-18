package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorElasticsearch TODO Doc Comment
func resourceSematextMonitorElasticsearch() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateElasticsearch,
		Read:   resourceMonitorReadElasticsearch,
		Update: resourceMonitorUpdateElasticsearch,
		Delete: resourceMonitorDeleteElasticsearch,
		Exists: resourceMonitorExistsElasticsearch,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateElasticsearch TODO Doc Comment
func resourceMonitorCreateElasticsearch(d *schema.ResourceData, meta interface{}) error {
	apptype := "Elastic Search"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadElasticsearch TODO Doc Comment
func resourceMonitorReadElasticsearch(d *schema.ResourceData, meta interface{}) error {
	apptype := "Elastic Search"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateElasticsearch TODO Doc Comment
func resourceMonitorUpdateElasticsearch(d *schema.ResourceData, meta interface{}) error {
	apptype := "Elastic Search"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteElasticsearch TODO Doc Comment
func resourceMonitorDeleteElasticsearch(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Elastic Search"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsElasticsearch TODO Doc Comment
func resourceMonitorExistsElasticsearch(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Elastic Search"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportElasticsearch TODO Doc Comment
func resourceSematextMonitorImportElasticsearch(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Elastic Search"
	return CommonMonitorImport(d, meta, apptype)
}
