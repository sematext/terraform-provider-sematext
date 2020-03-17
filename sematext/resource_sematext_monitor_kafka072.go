package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorKafka072 TODO Doc Comment
func resourceSematextMonitorKafka072() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateKafka072,
		Read:   resourceMonitorReadKafka072,
		Update: resourceMonitorUpdateKafka072,
		Delete: resourceMonitorDeleteKafka072,
		Exists: resourceMonitorExistsKafka072,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateKafka072 TODO Doc Comment
func resourceMonitorCreateKafka072(d *schema.ResourceData, meta interface{}) error {
	apptype := "Kafka-0.7.2"
	return CommonMonitorCreate(d, meta, apptype)
}

// resourceMonitorReadKafka072 TODO Doc Comment
func resourceMonitorReadKafka072(d *schema.ResourceData, meta interface{}) error {
	apptype := "Kafka-0.7.2"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateKafka072 TODO Doc Comment
func resourceMonitorUpdateKafka072(d *schema.ResourceData, meta interface{}) error {
	apptype := "Kafka-0.7.2"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteKafka072 TODO Doc Comment
func resourceMonitorDeleteKafka072(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Kafka-0.7.2"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsKafka072 TODO Doc Comment
func resourceMonitorExistsKafka072(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Kafka-0.7.2"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportKafka072 TODO Doc Comment
func resourceSematextMonitorImportKafka072(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Kafka-0.7.2"
	return CommonMonitorImport(d, meta, apptype)
}
