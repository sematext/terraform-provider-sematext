package sematext

import (
    "strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// resourceSematextMonitorSolrcloud TODO Doc Comment
func resourceSematextMonitorSolrcloud() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateSolrcloud,
		Read:   resourceSematextMonitorReadSolrcloud,
		Update: resourceSematextMonitorUpdateSolrcloud,
		Delete: resourceSematextMonitorDeleteSolrcloud,
		Exists: resourceSematextMonitorExistsSolrcloud,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateSolrcloud TODO Doc Comment
func resourceSematextMonitorCreateSolrcloud(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Solrcloud"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadSolrcloud TODO Doc Comment
func resourceSematextMonitorReadSolrcloud(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Solrcloud"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateSolrcloud TODO Doc Comment
func resourceSematextMonitorUpdateSolrcloud(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Solrcloud"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteSolrcloud TODO Doc Comment
func resourceSematextMonitorDeleteSolrcloud(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Solrcloud"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsSolrcloud TODO Doc Comment
func resourceSematextMonitorExistsSolrcloud(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Solrcloud"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportSolrcloud TODO Doc Comment
func resourceSematextMonitorImportSolrcloud(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Solrcloud"))
	return SematextMonitorImport(d, meta)
}
