package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorZookeeper is the resource class that handles sematext_monitor_zookeeper
func resourceSematextMonitorZookeeper() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("ZooKeeper")

	return &schema.Resource{
		Create: resourceMonitorCreateZookeeper,
		Read:   resourceMonitorReadZookeeper,
		Update: resourceMonitorUpdateZookeeper,
		Delete: resourceMonitorDeleteZookeeper,
		Exists: resourceMonitorExistsZookeeper,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateZookeeper creates the sematext_monitor_zookeeper resource.
func resourceMonitorCreateZookeeper(d *schema.ResourceData, meta interface{}) error {
	apptype := "ZooKeeper"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadZookeeper reads the sematext_monitor_zookeeper resource from Sematext Cloud.
func resourceMonitorReadZookeeper(d *schema.ResourceData, meta interface{}) error {
	apptype := "ZooKeeper"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateZookeeper updates Sematext Cloud from the sematext_monitor_zookeeper resource.
func resourceMonitorUpdateZookeeper(d *schema.ResourceData, meta interface{}) error {
	apptype := "ZooKeeper"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteZookeeper marks a sematext_monitor_zookeeper resource as retired.
func resourceMonitorDeleteZookeeper(d *schema.ResourceData, meta interface{}) error {
	apptype := "ZooKeeper"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsZookeeper checks a sematext_monitor_zookeeper resource exists in Sematext Cloud.
func resourceMonitorExistsZookeeper(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "ZooKeeper"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportZookeeper checks a sematext_monitor_zookeeper resource exists in Sematext Cloud.
func resourceSematextMonitorImportZookeeper(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "ZooKeeper"
	return CommonMonitorImport(d, meta, apptype)
}

*/
