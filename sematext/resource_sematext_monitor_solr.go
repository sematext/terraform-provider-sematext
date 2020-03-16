package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorSolr TODO Doc Comment
func resourceSematextMonitorSolr() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

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
	d.Set("app_type", strings.ToLower("Solr"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadSolr TODO Doc Comment
func resourceMonitorReadSolr(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Solr"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateSolr TODO Doc Comment
func resourceMonitorUpdateSolr(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Solr"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteSolr TODO Doc Comment
func resourceMonitorDeleteSolr(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Solr"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsSolr TODO Doc Comment
func resourceMonitorExistsSolr(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Solr"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportSolr TODO Doc Comment
func resourceSematextMonitorImportSolr(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Solr"))
	return CommonMonitorImport(d, meta)
}
