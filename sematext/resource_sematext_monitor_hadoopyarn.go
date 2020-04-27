package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorHadoopyarn is the resource class that handles sematext_monitor_hadoopyarn
func resourceSematextMonitorHadoopyarn() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Hadoop-YARN")

	return &schema.Resource{
		Create: resourceMonitorCreateHadoopyarn,
		Read:   resourceMonitorReadHadoopyarn,
		Update: resourceMonitorUpdateHadoopyarn,
		Delete: resourceMonitorDeleteHadoopyarn,
		Exists: resourceMonitorExistsHadoopyarn,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateHadoopyarn creates the sematext_monitor_hadoopyarn resource.
func resourceMonitorCreateHadoopyarn(d *schema.ResourceData, meta interface{}) error {
	apptype := "Hadoop-YARN"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadHadoopyarn reads the sematext_monitor_hadoopyarn resource from Sematext Cloud.
func resourceMonitorReadHadoopyarn(d *schema.ResourceData, meta interface{}) error {
	apptype := "Hadoop-YARN"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateHadoopyarn updates Sematext Cloud from the sematext_monitor_hadoopyarn resource.
func resourceMonitorUpdateHadoopyarn(d *schema.ResourceData, meta interface{}) error {
	apptype := "Hadoop-YARN"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteHadoopyarn marks a sematext_monitor_hadoopyarn resource as retired.
func resourceMonitorDeleteHadoopyarn(d *schema.ResourceData, meta interface{}) error {
	apptype := "Hadoop-YARN"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsHadoopyarn checks a sematext_monitor_hadoopyarn resource exists in Sematext Cloud.
func resourceMonitorExistsHadoopyarn(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Hadoop-YARN"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportHadoopyarn checks a sematext_monitor_hadoopyarn resource exists in Sematext Cloud.
func resourceSematextMonitorImportHadoopyarn(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Hadoop-YARN"
	return CommonMonitorImport(d, meta, apptype)
}

*/
