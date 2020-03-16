package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorTomcat TODO Doc Comment
func resourceSematextMonitorTomcat() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS metadata replacement target

	return &schema.Resource{
		Create: resourceMonitorCreateTomcat,
		Read:   resourceMonitorReadTomcat,
		Update: resourceMonitorUpdateTomcat,
		Delete: resourceMonitorDeleteTomcat,
		Exists: resourceMonitorExistsTomcat,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateTomcat TODO Doc Comment
func resourceMonitorCreateTomcat(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Tomcat"))
	return CommonMonitorCreate(d, meta)
}

// resourceMonitorReadTomcat TODO Doc Comment
func resourceMonitorReadTomcat(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Tomcat"))
	return CommonMonitorRead(d, meta)
}

// resourceMonitorUpdateTomcat TODO Doc Comment
func resourceMonitorUpdateTomcat(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Tomcat"))
	return CommonMonitorUpdate(d, meta)
}

// resourceMonitorDeleteTomcat TODO Doc Comment
func resourceMonitorDeleteTomcat(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Tomcat"))
	return CommonMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsTomcat TODO Doc Comment
func resourceMonitorExistsTomcat(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Tomcat"))
	return CommonMonitorExists(d, meta)
}

// resourceSematextMonitorImportTomcat TODO Doc Comment
func resourceSematextMonitorImportTomcat(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Tomcat"))
	return CommonMonitorImport(d, meta)
}
