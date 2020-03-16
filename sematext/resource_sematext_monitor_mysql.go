package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorMysql TODO Doc Comment
func resourceSematextMonitorMysql() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

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
	d.Set("app_type", strings.ToLower("Mysql"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadMysql TODO Doc Comment
func resourceMonitorReadMysql(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Mysql"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateMysql TODO Doc Comment
func resourceMonitorUpdateMysql(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Mysql"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteMysql TODO Doc Comment
func resourceMonitorDeleteMysql(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Mysql"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsMysql TODO Doc Comment
func resourceMonitorExistsMysql(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Mysql"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportMysql TODO Doc Comment
func resourceSematextMonitorImportMysql(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Mysql"))
	return CommonMonitorImport(d, meta)
}
