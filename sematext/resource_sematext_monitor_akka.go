package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_monitor.go.template
	Then run generate/generate.sh
*/

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceSematextMonitorAkka is the resource class that handles sematext_monitor_akka
func resourceSematextMonitorAkka() *schema.Resource {

	fieldSchema := MonitorSchemaCommon("Akka")

	return &schema.Resource{
		Create: resourceMonitorCreateAkka,
		Read:   resourceMonitorReadAkka,
		Update: resourceMonitorUpdateAkka,
		Delete: resourceMonitorDeleteAkka,
		Exists: resourceMonitorExistsAkka,
		Schema: fieldSchema,
	}
}

// resourceMonitorCreateAkka creates the sematext_monitor_akka resource.
func resourceMonitorCreateAkka(d *schema.ResourceData, meta interface{}) error {
	apptype := "Akka"
	err := CommonMonitorCreate(d, meta, apptype)

	return err
}

// resourceMonitorReadAkka reads the sematext_monitor_akka resource from Sematext Cloud.
func resourceMonitorReadAkka(d *schema.ResourceData, meta interface{}) error {
	apptype := "Akka"
	return CommonMonitorRead(d, meta, apptype)
}

// resourceMonitorUpdateAkka updates Sematext Cloud from the sematext_monitor_akka resource.
func resourceMonitorUpdateAkka(d *schema.ResourceData, meta interface{}) error {
	apptype := "Akka"
	return CommonMonitorUpdate(d, meta, apptype)
}

// resourceMonitorDeleteAkka marks a sematext_monitor_akka resource as retired.
func resourceMonitorDeleteAkka(d *schema.ResourceData, meta interface{}) error {
	apptype := "Akka"
	return CommonMonitorDelete(d, meta, apptype)
}

// resourceMonitorExistsAkka checks a sematext_monitor_akka resource exists in Sematext Cloud.
func resourceMonitorExistsAkka(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	apptype := "Akka"
	return CommonMonitorExists(d, meta, apptype)
}

/*

Placeholder - not implemented

// resourceSematextMonitorImportAkka checks a sematext_monitor_akka resource exists in Sematext Cloud.
func resourceSematextMonitorImportAkka(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	apptype := "Akka"
	return CommonMonitorImport(d, meta, apptype)
}

*/
