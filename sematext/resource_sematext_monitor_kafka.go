package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorKafka TODO Doc Comment
func resourceSematextMonitorKafka() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateKafka,
		Read:   resourceMonitorReadKafka,
		Update: resourceMonitorUpdateKafka,
		Delete: resourceMonitorDeleteKafka,
		Exists: resourceMonitorExistsKafka,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateKafka TODO Doc Comment
func resourceMonitorCreateKafka(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Kafka"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadKafka TODO Doc Comment
func resourceMonitorReadKafka(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Kafka"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateKafka TODO Doc Comment
func resourceMonitorUpdateKafka(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Kafka"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteKafka TODO Doc Comment
func resourceMonitorDeleteKafka(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Kafka"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsKafka TODO Doc Comment
func resourceMonitorExistsKafka(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Kafka"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportKafka TODO Doc Comment
func resourceSematextMonitorImportKafka(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Kafka"))
	return CommonMonitorImport(d, meta)
}
