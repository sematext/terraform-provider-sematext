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

			"id": { // TODO validate func
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"name": { // TODO validate func
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"description": { // TODO Need this?
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"billing_plan": {
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

			"discount_code": { // TODO validate func
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
		},
	}
}

// resourceSematextMonitorInfraCreate TODO Doc Comment
func resourceSematextMonitorInfraCreate(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", "infra")
	return api.SematextMonitorCreate(d, meta)
}

// resourceSematextMonitorInfraRead TODO Doc Comment
func resourceSematextMonitorInfraRead(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", "infra")
	return api.SematextMonitorRead(d, meta)
}

// resourceSematextMonitorInfraUpdate TODO Doc Comment
func resourceSematextMonitorInfraUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Set("app_type", "infra")
	return api.SematextMonitorUpdate(d, meta)
}

// resourceSematextMonitorInfraDelete TODO Doc Comment
func resourceSematextMonitorInfraDelete(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected
	d.Set("app_type", "infra")
	return api.SematextMonitorDelete(d, meta)
}

// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
// resourceSematextMonitorInfraExists TODO Doc Comment
func resourceSematextMonitorInfraExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	d.Set("app_type", "infra")
	return api.SematextMonitorExists(d, meta)
}

// resourceSematextMonitorInfraImport TODO Doc Comment
func resourceSematextMonitorInfraImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("app_type", "infra")
	return api.SematextMonitorImport(d, meta)
}
