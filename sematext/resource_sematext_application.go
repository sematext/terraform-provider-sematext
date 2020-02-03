package sematext

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSematextApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSematextApplicationCreate,
		Read:   resourceSematextApplicationRead,
		Update: resourceSematextApplicationUpdate,
		Delete: resourceSematextApplicationDelete,
		Exists: resourceSematextApplicationExists,
		Importer: &schema.ResourceImporter{
			State: resourceSematextApplicationImport,
		},

		Schema: map[string]*schema.Schema{},
	}
}

func resourceSematextApplicationCreate(d *schema.ResourceData, meta interface{}) error {
	application, err := buildSematextApplication(d)
	if err != nil {
		return fmt.Errorf("Failed to parse resource configuration: %s", err.Error())
	}
	application, err = meta.(*sematext.Client).CreateApplication(application)
	if err != nil {
		return fmt.Errorf("Failed to create application using API: %s", err.Error())
	}
	return resourceSematextApplicationRead(d, meta)
}


func resourceSematextApplicationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sematext.Config.Client)
	id := d.get(Id)
	application, err = client.getApplication(id)
	if err != "" {
		return err
	}
	d.set("id", application.id))
	d.set("token", application.token))
	d.set("appType", application.appType)
	d.set("description", application.description)
	d.set("integration", application.integration) // TODO - ServiceIntegration sub-schema in seperate file?
	d.set("name", application.name) // TODO - ServiceIntegration sub-schema in seperate file?
	d.set("owningOrganization", application.owningOrganization) // TODO - ServiceIntegration sub-schema in seperate file?
	d.set("prepaidAccount", application.prepaidAccount)
	d.set("status", application.status)
	d.set("userRoles", application.userRoles) // TODO - ServiceIntegration sub-schema in seperate file?

	if err != nil {
		return err
	}

}

func resourceSematextApplicationUpdate(d *schema.ResourceData, meta interface{}) error {
	application, err := buildSematextApplication(d)
	if err != nil {
		return fmt.Errorf("Failed to parse resource configuration: %s", err.Error())
	}
	application, err = meta.(*sematext.Client).UpdateApplication(application)
	if err != nil {
		return fmt.Errorf("Failed to create application using API: %s", err.Error())
	}
	return resourceSematextApplicationRead(d, meta)
}

func resourceSematextApplicationDelete(d *schema.ResourceData, meta interface{}) error { // TODO check this is protected
	id := d.Id()
	if err := meta.(*sematext.Client).DeleteApplication(id); err != nil {
		return err
	}
	return nil
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
func resourceSematextApplicationExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	id := d.Id()
	if _, err := meta.(*sematext.Config.Client).GetApplication(id); err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func resourceSematextApplicationImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := resourceSematextAppRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

// Helper Functions

func buildSematextApplication(d *schema.ResourceData) (*sematext.Application, error) {
	var application sematext.Application

	// TODO Field validations
	application.setId(d.Id())
	if v, ok := d.GetOk("field1"); ok {
		application.SetField1(v.(string))
	}
	if v, ok := d.GetOk("field2"); ok {
		application.SetField2(v.(string))
	}
	if v, ok := d.GetOk("field3"); ok {
		application.SetField3(v.(string))
	}

	return &application, nil

}
