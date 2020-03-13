package sematext

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorTomcat TODO Doc Comment
func resourceSematextMonitorTomcat() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// AWS replacement target

	return &schema.Resource{
		Create: resourceSematextMonitorCreateTomcat,
		Read:   resourceSematextMonitorReadTomcat,
		Update: resourceSematextMonitorUpdateTomcat,
		Delete: resourceSematextMonitorDeleteTomcat,
		Exists: resourceSematextMonitorExistsTomcat,
		Schema: fieldSchema,
	}
}

// resourceSematextMonitorCreateTomcat TODO Doc Comment
func resourceSematextMonitorCreateTomcat(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Tomcat"))
	return SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorReadTomcat TODO Doc Comment
func resourceSematextMonitorReadTomcat(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Tomcat"))
	return SematextMonitorRead(d, meta)
}

// resourceSematextMonitorUpdateTomcat TODO Doc Comment
func resourceSematextMonitorUpdateTomcat(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", strings.ToLower("Tomcat"))
	return SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorDeleteTomcat TODO Doc Comment
func resourceSematextMonitorDeleteTomcat(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", strings.ToLower("Tomcat"))
	return SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorExistsTomcat TODO Doc Comment
func resourceSematextMonitorExistsTomcat(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", strings.ToLower("Tomcat"))
	return SematextMonitorExists(d, meta)
}

// resourceSematextMonitorImportTomcat TODO Doc Comment
func resourceSematextMonitorImportTomcat(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", strings.ToLower("Tomcat"))
	return SematextMonitorImport(d, meta)
}
