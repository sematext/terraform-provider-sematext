package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorKafka TODO Doc Comment
func resourceSematextMonitorKafka() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateKafka,
		Read:   resourceSematextMonitorReadKafka,
		Update: resourceSematextMonitorUpdateKafka,
		Delete: resourceSematextMonitorDeleteKafka,
		Exists: resourceSematextMonitorExistsKafka,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateKafka TODO Doc Comment
func resourceSematextMonitorCreateKafka(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Kafka"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadKafka TODO Doc Comment
func resourceSematextMonitorReadKafka(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Kafka"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateKafka TODO Doc Comment
func resourceSematextMonitorUpdateKafka(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Kafka"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteKafka TODO Doc Comment
func resourceSematextMonitorDeleteKafka(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Kafka"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsKafka TODO Doc Comment
func resourceSematextMonitorExistsKafka(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Kafka"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportKafka TODO Doc Comment
func resourceSematextMonitorImportKafka(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Kafka"))
	return SematextMonitorImport(d, meta)
}
