package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorRabbitmq is the resource class that handles sematext_monitor_rabbitmq
func resourceSematextMonitorRabbitmq() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("RabbitMQ")

	return &schema.Resource{
		Create: resourceMonitorCreateRabbitmq,
		Read:   resourceMonitorReadRabbitmq,
		Update: resourceMonitorUpdateRabbitmq,
		Delete: resourceMonitorDeleteRabbitmq,
		Exists: resourceMonitorExistsRabbitmq,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateRabbitmq creates the sematext_monitor_rabbitmq resource.
func resourceMonitorCreateRabbitmq(d *schema.ResourceData, meta interface{}) error {
	apptype := "RabbitMQ"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadRabbitmq reads the sematext_monitor_rabbitmq resource from Sematext Cloud.
func resourceMonitorReadRabbitmq(d *schema.ResourceData, meta interface{}) error {
	apptype := "RabbitMQ"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateRabbitmq updates Sematext Cloud from the sematext_monitor_rabbitmq resource.
func resourceMonitorUpdateRabbitmq(d *schema.ResourceData, meta interface{}) error {
	apptype := "RabbitMQ"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteRabbitmq marks a sematext_monitor_rabbitmq resource as retired.
func resourceMonitorDeleteRabbitmq(d *schema.ResourceData, meta interface{}) error {
	apptype := "RabbitMQ"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsRabbitmq checks a sematext_monitor_rabbitmq resource exists in Sematext Cloud.
func resourceMonitorExistsRabbitmq(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "RabbitMQ"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportRabbitmq checks a sematext_monitor_rabbitmq resource exists in Sematext Cloud.
func resourceSematextMonitorImportRabbitmq(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "RabbitMQ"
	return CommonMonitorImport(d, meta, apptype)
}

*/
