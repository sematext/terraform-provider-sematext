package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorTomcat TODO Doc Comment
func resourceSematextMonitorTomcat() *schema.Resource {

	fieldSchema := MonitorSchemaCommon

	// TODO AWS* metadata replacement target

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
	apptype := "Tomcat"
	return CommonMonitorCreate(d, meta, apptype)
}

// resourceMonitorReadTomcat TODO Doc Comment
func resourceMonitorReadTomcat(d *schema.ResourceData, meta interface{}) error {
	apptype := "Tomcat"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateTomcat TODO Doc Comment
func resourceMonitorUpdateTomcat(d *schema.ResourceData, meta interface{}) error {
	apptype := "Tomcat"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteTomcat TODO Doc Comment
func resourceMonitorDeleteTomcat(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Tomcat"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsTomcat TODO Doc Comment
func resourceMonitorExistsTomcat(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Tomcat"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportTomcat TODO Doc Comment
func resourceSematextMonitorImportTomcat(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Tomcat"
	return CommonMonitorImport(d, meta, apptype)
}
