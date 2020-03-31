package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorHadoopmrv1 TODO Doc Comment
func resourceSematextMonitorHadoopmrv1() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Hadoop-MRv1")

	return &schema.Resource{
		Create: resourceMonitorCreateHadoopmrv1,
		Read:   resourceMonitorReadHadoopmrv1,
		Update: resourceMonitorUpdateHadoopmrv1,
		Delete: resourceMonitorDeleteHadoopmrv1,
		Exists: resourceMonitorExistsHadoopmrv1,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateHadoopmrv1 TODO Doc Comment
func resourceMonitorCreateHadoopmrv1(d *schema.ResourceData, meta interface{}) error {
	apptype := "Hadoop-MRv1"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadHadoopmrv1 TODO Doc Comment
func resourceMonitorReadHadoopmrv1(d *schema.ResourceData, meta interface{}) error {
	apptype := "Hadoop-MRv1"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateHadoopmrv1 TODO Doc Comment
func resourceMonitorUpdateHadoopmrv1(d *schema.ResourceData, meta interface{}) error {
	apptype := "Hadoop-MRv1"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteHadoopmrv1 TODO Doc Comment
func resourceMonitorDeleteHadoopmrv1(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Hadoop-MRv1"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsHadoopmrv1 TODO Doc Comment
func resourceMonitorExistsHadoopmrv1(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Hadoop-MRv1"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportHadoopmrv1 TODO Doc Comment
func resourceSematextMonitorImportHadoopmrv1(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Hadoop-MRv1"
	return CommonMonitorImport(d, meta, apptype)
}
