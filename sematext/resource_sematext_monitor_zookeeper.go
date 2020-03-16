package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorZookeeper TODO Doc Comment
func resourceSematextMonitorZookeeper() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

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
	d.Set("app_type", strings.ToLower("Zookeeper"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadZookeeper TODO Doc Comment
func resourceMonitorReadZookeeper(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Zookeeper"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateZookeeper TODO Doc Comment
func resourceMonitorUpdateZookeeper(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Zookeeper"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteZookeeper TODO Doc Comment
func resourceMonitorDeleteZookeeper(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Zookeeper"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsZookeeper TODO Doc Comment
func resourceMonitorExistsZookeeper(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Zookeeper"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportZookeeper TODO Doc Comment
func resourceSematextMonitorImportZookeeper(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Zookeeper"))
	return CommonMonitorImport(d, meta)
}
