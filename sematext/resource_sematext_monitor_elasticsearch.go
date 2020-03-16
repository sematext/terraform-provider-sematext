package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorElasticsearch TODO Doc Comment
func resourceSematextMonitorElasticsearch() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

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
	d.Set("app_type", strings.ToLower("Elasticsearch"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadElasticsearch TODO Doc Comment
func resourceMonitorReadElasticsearch(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Elasticsearch"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateElasticsearch TODO Doc Comment
func resourceMonitorUpdateElasticsearch(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Elasticsearch"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteElasticsearch TODO Doc Comment
func resourceMonitorDeleteElasticsearch(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Elasticsearch"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsElasticsearch TODO Doc Comment
func resourceMonitorExistsElasticsearch(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Elasticsearch"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportElasticsearch TODO Doc Comment
func resourceSematextMonitorImportElasticsearch(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Elasticsearch"))
	return CommonMonitorImport(d, meta)
}
