package sematext

import (
	"fmt"
	"github.com/thoas/go-funk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSematextMonitorInfra() *schema.Resource {
	return &schema.Resource{
		Create: resourceSematextMonitorInfraCreate,
		Read:   resourceSematextMonitorInfraRead,
		Update: resourceSematextMonitorInfraUpdate,
		Delete: resourceSematextMonitorInfraDelete,
		Exists: resourceSematextMonitorInfraExists,
		Importer: &schema.ResourceImporter{
			State: resourceSematextMonitorInfraImport,
		},

		Schema: map[string]*schema.Schema{

			"id": &schema.Schema{  // TODO validate func
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"label": &schema.Schema{  // TODO validate func
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"description": &schema.Schema{  // TODO Need this?
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"billing_plan": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
				ValidateFunc : func(value string, key string) (warns []string, errs []error) {
					allowed := []string{"basic","standard","pro","enterprise"}
					if !funk.Contains(allowed, value) {
						errs = append(errs, fmt.Errorf("billing_plan must be one of %v : got %s", allowed, value))
					}
					return
				}
			},

			"discount_code": &schema.Schema{  // TODO validate func
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				ValidateFunc : func(value string, key string) (warns []string, errs []error) {
					if value.len == 0 {
						errs = append(errs, fmt.Errorf("discount_code is set with zero length"))
					}
					return
				}
			},

			"lifecycle": &schema.Schema{
				Optional: true,
				ForceNew: true, // TODO Confirm use of ForceNew.
				Type:     schema.Schema {
					"delete_on_destroy" : {
						Type:     schema.TypeBool,
						Optional: true,
						DefaultFunc func() {interface{}, error} {
							return false, nil
						},
						ValidateFunc : func(value bool, key string) (warns []string, errs []error) {
							if value == true {
								warns = append(warns, fmt.Errorf("Note: delete_on_destroy is ON - Sematext data collection for this resource will be turned off and will be deleted on resource-destroy!"))
							} else {
								warns = append(warns, fmt.Errorf("Note: delete_on_destroy is OFF - Sematext data collection for this resource will be set to archived but wil be retainedon resource-destroy!"))
							}
							return
						}

					}
				}
			}
		}
	}
}


func resourceSematextMonitorInfraCreate(d *schema.ResourceData, meta interface{}) error {

	plans := map[string]int{"basic":0,"standard":1,"pro":2,"enterprise":3}
	createAppInfo :=  &CreateAppInfo{}
	createAppInfo.AppType = "infra"
	createAppInfo.DiscountCode = d.Get("discount_code")
	createAppInfo.InitialPlanID = plans[d.Get("billing_plan")]
	//createAppInfo.MetaData = &AppMetaData{} // not used for this resource
	createAppInfo.Name = d.Get("label")
	client := meta.(*ProviderConfig).Config.Api.Client
	app, err := createAppInfo.Create(client)
	if (err != nil){
		return error
	}

	d.setId(app.Id)
}


func resourceSematextMonitorInfraRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sematext.Config.Client)
	id := d.get(Id)
	application, err = client.getApplication(id)
	if err != nil {
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

func resourceSematextMonitorInfraUpdate(d *schema.ResourceData, meta interface{}) error {
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

func resourceSematextMonitorInfraDelete(d *schema.ResourceData, meta interface{}) error { // TODO check this is protected
	id := d.Id()
	if err := meta.(*sematext.Client).DeleteApplication(id); err != nil {
		return err
	}
	return nil
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
func resourceSematextMonitorInfraExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	id := d.Id()
	if _, err := meta.(*sematext.Config.Client).GetApplication(id); err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func resourceSematextMonitorInfra(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
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
