package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorSolrcloud TODO Doc Comment
func resourceSematextMonitorSolrcloud() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

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
	d.Set("app_type", strings.ToLower("Solrcloud"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadSolrcloud TODO Doc Comment
func resourceMonitorReadSolrcloud(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Solrcloud"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateSolrcloud TODO Doc Comment
func resourceMonitorUpdateSolrcloud(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Solrcloud"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteSolrcloud TODO Doc Comment
func resourceMonitorDeleteSolrcloud(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Solrcloud"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsSolrcloud TODO Doc Comment
func resourceMonitorExistsSolrcloud(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Solrcloud"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportSolrcloud TODO Doc Comment
func resourceSematextMonitorImportSolrcloud(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Solrcloud"))
	return CommonMonitorImport(d, meta)
}
