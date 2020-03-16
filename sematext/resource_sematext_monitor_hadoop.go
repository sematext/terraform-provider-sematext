package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorHadoop TODO Doc Comment
func resourceSematextMonitorHadoop() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateHadoop,
		Read:   resourceMonitorReadHadoop,
		Update: resourceMonitorUpdateHadoop,
		Delete: resourceMonitorDeleteHadoop,
		Exists: resourceMonitorExistsHadoop,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateHadoop TODO Doc Comment
func resourceMonitorCreateHadoop(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Hadoop"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadHadoop TODO Doc Comment
func resourceMonitorReadHadoop(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Hadoop"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateHadoop TODO Doc Comment
func resourceMonitorUpdateHadoop(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Hadoop"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteHadoop TODO Doc Comment
func resourceMonitorDeleteHadoop(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Hadoop"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsHadoop TODO Doc Comment
func resourceMonitorExistsHadoop(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Hadoop"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportHadoop TODO Doc Comment
func resourceSematextMonitorImportHadoop(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Hadoop"))
	return CommonMonitorImport(d, meta)
}
