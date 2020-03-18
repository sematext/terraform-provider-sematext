package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorKafka TODO Doc Comment
func resourceSematextMonitorKafka() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

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
	apptype := "Kafka"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadKafka TODO Doc Comment
func resourceMonitorReadKafka(d *schema.ResourceData, meta interface{}) error {
	apptype := "Kafka"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateKafka TODO Doc Comment
func resourceMonitorUpdateKafka(d *schema.ResourceData, meta interface{}) error {
	apptype := "Kafka"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteKafka TODO Doc Comment
func resourceMonitorDeleteKafka(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Kafka"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsKafka TODO Doc Comment
func resourceMonitorExistsKafka(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Kafka"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportKafka TODO Doc Comment
func resourceSematextMonitorImportKafka(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Kafka"
	return CommonMonitorImport(d, meta, apptype)
}
