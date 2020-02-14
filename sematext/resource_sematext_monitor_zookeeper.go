package sematext

import (
	"fmt"
	"github.com/thoas/go-funk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// resourceSematextMonitorZookeeper TODO Doc Comment
func resourceSematextMonitorZookeeper() *schema.Resource {
	return &schema.Resource{
		Create: resourceSematextMonitorZookeeperCreate,
		Read:   resourceSematextMonitorZookeeperRead,
		Update: resourceSematextMonitorZookeeperUpdate,
		Delete: resourceSematextMonitorZookeeperDelete,
		Exists: resourceSematextMonitorZookeeperExists,
		Importer: &schema.ResourceImporter{
			State: resourceSematextMonitorZookeeperImport,
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

			/*
			"metadata": &schema.Schema{
				Optional: true,
				ForceNew: true,
				Type:     schema.Schema {
					"aws_cloudwatch" : &schema.Schema{ // TODO pull out of ENV
						Required: true,
						ForceNew: true,
						Type : &schema.Schema{
							"access_key" : {
								Type:     schema.TypeString,
								Required: true,
							},
							"secret_key" : {
								Type:     schema.TypeString,
								Required: true,
							},
							"fetch_frequency" : {
								Type:     schema.TypeString,
								Required: true,
							},
							"region" : {
								Type:     schema.TypeString,
								Required: true,
							},

						}
					}
				}
			}
			*/

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


// resourceSematextMonitorZookeeperCreate TODO Doc Comment
func resourceSematextMonitorZookeeperCreate(d *schema.ResourceData, meta interface{}) error {

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

// resourceSematextMonitorZookeeperRead TODO Doc Comment
func resourceSematextMonitorZookeeperRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Provider.Config.Api.Client) // TODO check reference
	id := d.Id()
	app, err := App.Retrieve(id, client)
	if err != nil{
		return nil, err
	}
	plans := map[int]string{0: "basic",1: "standard",2:"pro",3: "enterprise":3}
	d.set("label", app.Name)
	d.set("description",app.Description)
	d.set("billing_plan", plans[App.Pan.ID])
	//d.set("discount_code",) // TODO - this seems to not be stored in the App.
	//d.set("lifecycle.delete_on_destroy",)  // TODO Computed from resource, not stored on backend?

}


// resourceSematextMonitorZookeeperUpdate TODO Doc Comment
func resourceSematextMonitorZookeeperUpdate(d *schema.ResourceData, meta interface{}) error {


	// TODO - how to do a plan update?

	client := meta.(*Provider.Config.Api.Client) // TODO check reference
	id := d.Id()
	dto := &Dto{}

	description := d.get("description")
	if description != nil {
		dto.Description = description
	}
	ignorePercentage := d.get("ignore_percentage")
	if ignorePercentage != nil {
		dto.ignorePercentage = ignorePercentage
	}
	maxEvents := d.get("max_events")
	if maxEvents != nil {
		dto.MaxEvents = maxEvents
	}
	maxLimitMB := d.get("max_limit_mb")
	if maxLimitMB != nil {
		dto.maxLimitMB = maxLimitMB
	}
	name := d.get("name") // TODO name vs resource name/label in terraform
	if name != nil {
		dto.Name = name
	}
	sampling := d.get("sampling")
	if description != nil {
		dto.Sampling = sampling
	}
	samplingPercentage := d.get("sampling_percentage")
	if samplingPercentage != nil {
		dto.SamplingPercentage = samplingPercentage
	}
	staggering := d.get("staggering")
	if staggering != nil {
		dto.Staggering = staggering
	}
	status := d.get("status")
	if status != nil {
		dto.Status = status // TODO Status should be reset as part of delete?
	}

	if dto.isValid() == false {
		return fmt.Errorf("Coding error - dto non-valid") // TODO Check correctness.
	}

	app, err := dto.Update(client, dto)
	if err != null {
		return err
	}

	return nil

}


// resourceSematextMonitorZookeeperDelete TODO Doc Comment
func resourceSematextMonitorZookeeperDelete(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected

	client := meta.(*Provider.Config.Api.Client) // TODO check reference
	id := d.Id()
	var app App
	err = app.Delete(id, client)
	if err != nil {
		return err
	}

	return nil
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorZookeeperExists TODO Doc Comment
func resourceSematextMonitorZookeeperExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {

	client := meta.(*Provider.Config.Api.Client) // TODO check reference
	id := d.Id()
	var app App
	b, e = app.Exists(id, client) // TODO Pointer-value mismatch?
}

// resourceSematextMonitorZookeeperImport TODO Doc Comment
func resourceSematextMonitorZookeeperImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	// TODO Decide if necessary

}
