package sematext

import (
    "strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// resourceSematextMonitorHadoop TODO Doc Comment
func resourceSematextMonitorHadoop() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateHadoop,
		Read:   resourceSematextMonitorReadHadoop,
		Update: resourceSematextMonitorUpdateHadoop,
		Delete: resourceSematextMonitorDeleteHadoop,
		Exists: resourceSematextMonitorExistsHadoop,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateHadoop TODO Doc Comment
func resourceSematextMonitorCreateHadoop(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Hadoop"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadHadoop TODO Doc Comment
func resourceSematextMonitorReadHadoop(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Hadoop"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateHadoop TODO Doc Comment
func resourceSematextMonitorUpdateHadoop(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Hadoop"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteHadoop TODO Doc Comment
func resourceSematextMonitorDeleteHadoop(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Hadoop"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsHadoop TODO Doc Comment
func resourceSematextMonitorExistsHadoop(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Hadoop"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportHadoop TODO Doc Comment
func resourceSematextMonitorImportHadoop(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Hadoop"))
	return SematextMonitorImport(d, meta)
}
