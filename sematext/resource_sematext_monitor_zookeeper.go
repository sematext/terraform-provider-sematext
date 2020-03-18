package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorZookeeper TODO Doc Comment
func resourceSematextMonitorZookeeper() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateZookeeper,
		Read:   resourceMonitorReadZookeeper,
		Update: resourceMonitorUpdateZookeeper,
		Delete: resourceMonitorDeleteZookeeper,
		Exists: resourceMonitorExistsZookeeper,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateZookeeper TODO Doc Comment
func resourceMonitorCreateZookeeper(d *schema.ResourceData, meta interface{}) error {
	apptype := "ZooKeeper"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadZookeeper TODO Doc Comment
func resourceMonitorReadZookeeper(d *schema.ResourceData, meta interface{}) error {
	apptype := "ZooKeeper"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateZookeeper TODO Doc Comment
func resourceMonitorUpdateZookeeper(d *schema.ResourceData, meta interface{}) error {
	apptype := "ZooKeeper"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteZookeeper TODO Doc Comment
func resourceMonitorDeleteZookeeper(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "ZooKeeper"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsZookeeper TODO Doc Comment
func resourceMonitorExistsZookeeper(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "ZooKeeper"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportZookeeper TODO Doc Comment
func resourceSematextMonitorImportZookeeper(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "ZooKeeper"
	return CommonMonitorImport(d, meta, apptype)
}
