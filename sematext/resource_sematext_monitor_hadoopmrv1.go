package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorHadoopmrv1 is the resource class that handles sematext_monitor_hadoopmrv1
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

// resourceMonitorCreateHadoopmrv1 creates the sematext_monitor_hadoopmrv1 resource.
func resourceMonitorCreateHadoopmrv1(d *schema.ResourceData, meta interface{}) error {
	apptype := "Hadoop-MRv1"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadHadoopmrv1 reads the sematext_monitor_hadoopmrv1 resource from Sematext Cloud.
func resourceMonitorReadHadoopmrv1(d *schema.ResourceData, meta interface{}) error {
	apptype := "Hadoop-MRv1"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateHadoopmrv1 updates Sematext Cloud from the sematext_monitor_hadoopmrv1 resource.
func resourceMonitorUpdateHadoopmrv1(d *schema.ResourceData, meta interface{}) error {
	apptype := "Hadoop-MRv1"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteHadoopmrv1 marks a sematext_monitor_hadoopmrv1 resource as retired.
func resourceMonitorDeleteHadoopmrv1(d *schema.ResourceData, meta interface{}) error {
	apptype := "Hadoop-MRv1"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsHadoopmrv1 checks a sematext_monitor_hadoopmrv1 resource exists in Sematext Cloud.
func resourceMonitorExistsHadoopmrv1(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Hadoop-MRv1"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportHadoopmrv1 checks a sematext_monitor_hadoopmrv1 resource exists in Sematext Cloud.
func resourceSematextMonitorImportHadoopmrv1(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Hadoop-MRv1"
	return CommonMonitorImport(d, meta, apptype)
}

*/
