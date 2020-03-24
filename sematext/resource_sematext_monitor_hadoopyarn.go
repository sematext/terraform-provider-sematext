package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorHadoopyarn TODO Doc Comment
func resourceSematextMonitorHadoopyarn() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Hadoopyarn")

	return &schema.Resource{
		Create: resourceMonitorCreateHadoopyarn,
		Read:   resourceMonitorReadHadoopyarn,
		Update: resourceMonitorUpdateHadoopyarn,
		Delete: resourceMonitorDeleteHadoopyarn,
		Exists: resourceMonitorExistsHadoopyarn,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateHadoopyarn TODO Doc Comment
func resourceMonitorCreateHadoopyarn(d *schema.ResourceData, meta interface{}) error {
	apptype := "Hadoop-YARN"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadHadoopyarn TODO Doc Comment
func resourceMonitorReadHadoopyarn(d *schema.ResourceData, meta interface{}) error {
	apptype := "Hadoop-YARN"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateHadoopyarn TODO Doc Comment
func resourceMonitorUpdateHadoopyarn(d *schema.ResourceData, meta interface{}) error {
	apptype := "Hadoop-YARN"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteHadoopyarn TODO Doc Comment
func resourceMonitorDeleteHadoopyarn(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	apptype := "Hadoop-YARN"
	return CommonMonitorDelete(d, meta, apptype)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceMonitorExistsHadoopyarn TODO Doc Comment
func resourceMonitorExistsHadoopyarn(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Hadoop-YARN"
	return CommonMonitorExists(d, meta, apptype)
}

// resourceSematextMonitorImportHadoopyarn TODO Doc Comment
func resourceSematextMonitorImportHadoopyarn(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Hadoop-YARN"
	return CommonMonitorImport(d, meta, apptype)
}
