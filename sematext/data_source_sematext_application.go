package sematext

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSematextApplication() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceSematextApplicationRead,

		Schema: map[string]*schema.Schema{

			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Todo:     true, // TODO Plan sub-struct in seperate file?
			},

			"appType": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Todo:     true, // TODO enum
			},

			"appType": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Todo:     true, // TODO enum
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Todo:     true, // TODO
			},

			"integration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Todo:     true, // TODO - ServiceIntegration sub-schema in seperate file?
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Todo:     true, // TODO
			},

			"owningOrganization": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Todo:     true, // TODO BasicOrganizationDto sub-struct in seperate file?
			},

			"plan": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Todo:     true, // TODO Plan sub-struct in seperate file?
			},

			"prepaidAccount": &schema.Schema{
				Type:     schema.TypeBoolean,
				Optional: true,
				Todo:     true, // TODO Required?
			},

			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Todo:     true, // TODO Plan sub-struct in seperate file?
			},

			"userRoles": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Todo:     true, // TODO array of UserRole sub-struct in seperate file?
			},
		},
	}
}

func dataSourceSematextApplicationRead(d *schema.ResourceData, meta interface{}) error {

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

	return nil
}
