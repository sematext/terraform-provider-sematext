package sematext

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sematext/terraform-provider-sematext/sematext.com/api"
	funk "github.com/thoas/go-funk"
)

// resourceSematextMonitorInfra TODO Doc Comment
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

			"id": &schema.Schema{ // TODO validate func
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"label": &schema.Schema{ // TODO validate func
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"description": &schema.Schema{ // TODO Need this?
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"billing_plan": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
				ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
					allowed := []string{"basic", "standard", "pro", "enterprise"}
					if !funk.Contains(allowed, value.(string)) {
						errs = append(errs, fmt.Errorf("billing_plan must be one of %v : got %s", allowed, value))
					}
					return warns, errs
				},
			},

			"discount_code": &schema.Schema{ // TODO validate func
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
					if len(value.(string)) == 0 {
						errs = append(errs, fmt.Errorf("discount_code is set with zero length"))
					}
					return warns, errs
				},
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
				Type: schema.Schema{
					"delete_on_destroy": {
						Type:     schema.TypeBool,
						Optional: true,
						DefaultFunc: func() (interface{}, error) {
							return false, nil
						},
						ValidateFunc: func(value bool, key string) (warns []string, errs []error) {
							if value == true {
								warns = append(warns, fmt.Errorf("Note: delete_on_destroy is ON - Sematext data collection for this resource will be turned off and will be deleted on resource-destroy"))
							} else {
								warns = append(warns, fmt.Errorf("Note: delete_on_destroy is OFF - Sematext data collection for this resource will be set to archived but wil be retainedon resource-destroy"))
							}
						},
					},
				},
			},
		},
	}
}

// resourceSematextMonitorInfraCreate TODO Doc Comment
func resourceSematextMonitorInfraCreate(d *schema.ResourceData, meta interface{}) error {
	return api.SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorInfraRead TODO Doc Comment
func resourceSematextMonitorInfraRead(d *schema.ResourceData, meta interface{}) error {
	return api.SematextMonitorRead(d, meta)
}

// resourceSematextMonitorInfraUpdate TODO Doc Comment
func resourceSematextMonitorInfraUpdate(d *schema.ResourceData, meta interface{}) error {
	return api.SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorInfraDelete TODO Doc Comment
func resourceSematextMonitorInfraDelete(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	return api.SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorInfraExists TODO Doc Comment
func resourceSematextMonitorInfraExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	return api.SematextMonitorExists(d, meta)
}

// resourceSematextMonitorInfraImport TODO Doc Comment
func resourceSematextMonitorInfraImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return api.SematextMonitorImport(d, meta)
}
