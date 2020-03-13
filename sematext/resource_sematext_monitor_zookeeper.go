package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorZookeeper TODO Doc Comment
func resourceSematextMonitorZookeeper() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateZookeeper,
		Read:   resourceSematextMonitorReadZookeeper,
		Update: resourceSematextMonitorUpdateZookeeper,
		Delete: resourceSematextMonitorDeleteZookeeper,
		Exists: resourceSematextMonitorExistsZookeeper,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateZookeeper TODO Doc Comment
func resourceSematextMonitorCreateZookeeper(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Zookeeper"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadZookeeper TODO Doc Comment
func resourceSematextMonitorReadZookeeper(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Zookeeper"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateZookeeper TODO Doc Comment
func resourceSematextMonitorUpdateZookeeper(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Zookeeper"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteZookeeper TODO Doc Comment
func resourceSematextMonitorDeleteZookeeper(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Zookeeper"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsZookeeper TODO Doc Comment
func resourceSematextMonitorExistsZookeeper(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Zookeeper"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportZookeeper TODO Doc Comment
func resourceSematextMonitorImportZookeeper(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Zookeeper"))
	return SematextMonitorImport(d, meta)
}
