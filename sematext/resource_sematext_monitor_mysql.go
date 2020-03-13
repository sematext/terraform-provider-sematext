package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorMysql TODO Doc Comment
func resourceSematextMonitorMysql() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateMysql,
		Read:   resourceSematextMonitorReadMysql,
		Update: resourceSematextMonitorUpdateMysql,
		Delete: resourceSematextMonitorDeleteMysql,
		Exists: resourceSematextMonitorExistsMysql,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateMysql TODO Doc Comment
func resourceSematextMonitorCreateMysql(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Mysql"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadMysql TODO Doc Comment
func resourceSematextMonitorReadMysql(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Mysql"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateMysql TODO Doc Comment
func resourceSematextMonitorUpdateMysql(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Mysql"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteMysql TODO Doc Comment
func resourceSematextMonitorDeleteMysql(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Mysql"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsMysql TODO Doc Comment
func resourceSematextMonitorExistsMysql(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Mysql"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportMysql TODO Doc Comment
func resourceSematextMonitorImportMysql(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Mysql"))
	return SematextMonitorImport(d, meta)
}
