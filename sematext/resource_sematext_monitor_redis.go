package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorRedis is the resource class that handles sematext_monitor_redis
func resourceSematextMonitorRedis() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Redis")

	return &schema.Resource{
		Create: resourceMonitorCreateRedis,
		Read:   resourceMonitorReadRedis,
		Update: resourceMonitorUpdateRedis,
		Delete: resourceMonitorDeleteRedis,
		Exists: resourceMonitorExistsRedis,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateRedis creates the sematext_monitor_redis resource.
func resourceMonitorCreateRedis(d *schema.ResourceData, meta interface{}) error {
	apptype := "Redis"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadRedis reads the sematext_monitor_redis resource from Sematext Cloud.
func resourceMonitorReadRedis(d *schema.ResourceData, meta interface{}) error {
	apptype := "Redis"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateRedis updates Sematext Cloud from the sematext_monitor_redis resource.
func resourceMonitorUpdateRedis(d *schema.ResourceData, meta interface{}) error {
	apptype := "Redis"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteRedis marks a sematext_monitor_redis resource as retired.
func resourceMonitorDeleteRedis(d *schema.ResourceData, meta interface{}) error {
	apptype := "Redis"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsRedis checks a sematext_monitor_redis resource exists in Sematext Cloud.
func resourceMonitorExistsRedis(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Redis"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportRedis checks a sematext_monitor_redis resource exists in Sematext Cloud.
func resourceSematextMonitorImportRedis(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Redis"
	return CommonMonitorImport(d, meta, apptype)
}

*/
