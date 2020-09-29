package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorKafka is the resource class that handles sematext_monitor_kafka
func resourceSematextMonitorKafka() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Kafka")

	return &schema.Resource{
		Create: resourceMonitorCreateKafka,
		Read:   resourceMonitorReadKafka,
		Update: resourceMonitorUpdateKafka,
		Delete: resourceMonitorDeleteKafka,
		Exists: resourceMonitorExistsKafka,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateKafka creates the sematext_monitor_kafka resource.
func resourceMonitorCreateKafka(d *schema.ResourceData, meta interface{}) error {
	apptype := "Kafka"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadKafka reads the sematext_monitor_kafka resource from Sematext Cloud.
func resourceMonitorReadKafka(d *schema.ResourceData, meta interface{}) error {
	apptype := "Kafka"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateKafka updates Sematext Cloud from the sematext_monitor_kafka resource.
func resourceMonitorUpdateKafka(d *schema.ResourceData, meta interface{}) error {
	apptype := "Kafka"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteKafka marks a sematext_monitor_kafka resource as retired.
func resourceMonitorDeleteKafka(d *schema.ResourceData, meta interface{}) error {
	apptype := "Kafka"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsKafka checks a sematext_monitor_kafka resource exists in Sematext Cloud.
func resourceMonitorExistsKafka(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Kafka"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportKafka checks a sematext_monitor_kafka resource exists in Sematext Cloud.
func resourceSematextMonitorImportKafka(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Kafka"
	return CommonMonitorImport(d, meta, apptype)
}

*/
