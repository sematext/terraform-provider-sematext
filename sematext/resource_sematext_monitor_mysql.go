package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorMysql TODO Doc Comment
func resourceSematextMonitorMysql() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateMysql,
		Read:   resourceMonitorReadMysql,
		Update: resourceMonitorUpdateMysql,
		Delete: resourceMonitorDeleteMysql,
		Exists: resourceMonitorExistsMysql,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateMysql TODO Doc Comment
func resourceMonitorCreateMysql(d *schema.ResourceData, meta interface{}) error {
	apptype := "MySQL"
	return CommonMonitorCreate(d, meta, apptype)
}

// resourceMonitorReadMysql TODO Doc Comment
func resourceMonitorReadMysql(d *schema.ResourceData, meta interface{}) error {
	apptype := "MySQL"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateMysql TODO Doc Comment
func resourceMonitorUpdateMysql(d *schema.ResourceData, meta interface{}) error {
	apptype := "MySQL"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteMysql TODO Doc Comment
func resourceMonitorDeleteMysql(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "MySQL"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsMysql TODO Doc Comment
func resourceMonitorExistsMysql(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "MySQL"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportMysql TODO Doc Comment
func resourceSematextMonitorImportMysql(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "MySQL"
	return CommonMonitorImport(d, meta, apptype)
}
