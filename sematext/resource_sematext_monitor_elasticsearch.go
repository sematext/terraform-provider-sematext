package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorElasticsearch TODO Doc Comment
func resourceSematextMonitorElasticsearch() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateElasticsearch,
		Read:   resourceSematextMonitorReadElasticsearch,
		Update: resourceSematextMonitorUpdateElasticsearch,
		Delete: resourceSematextMonitorDeleteElasticsearch,
		Exists: resourceSematextMonitorExistsElasticsearch,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateElasticsearch TODO Doc Comment
func resourceSematextMonitorCreateElasticsearch(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Elasticsearch"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadElasticsearch TODO Doc Comment
func resourceSematextMonitorReadElasticsearch(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Elasticsearch"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateElasticsearch TODO Doc Comment
func resourceSematextMonitorUpdateElasticsearch(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Elasticsearch"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteElasticsearch TODO Doc Comment
func resourceSematextMonitorDeleteElasticsearch(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Elasticsearch"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsElasticsearch TODO Doc Comment
func resourceSematextMonitorExistsElasticsearch(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Elasticsearch"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportElasticsearch TODO Doc Comment
func resourceSematextMonitorImportElasticsearch(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Elasticsearch"))
	return SematextMonitorImport(d, meta)
}
