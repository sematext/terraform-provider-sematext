package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorSolr TODO Doc Comment
func resourceSematextMonitorSolr() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateSolr,
		Read:   resourceSematextMonitorReadSolr,
		Update: resourceSematextMonitorUpdateSolr,
		Delete: resourceSematextMonitorDeleteSolr,
		Exists: resourceSematextMonitorExistsSolr,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateSolr TODO Doc Comment
func resourceSematextMonitorCreateSolr(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Solr"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadSolr TODO Doc Comment
func resourceSematextMonitorReadSolr(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Solr"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateSolr TODO Doc Comment
func resourceSematextMonitorUpdateSolr(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Solr"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteSolr TODO Doc Comment
func resourceSematextMonitorDeleteSolr(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Solr"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsSolr TODO Doc Comment
func resourceSematextMonitorExistsSolr(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Solr"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportSolr TODO Doc Comment
func resourceSematextMonitorImportSolr(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Solr"))
	return SematextMonitorImport(d, meta)
}
