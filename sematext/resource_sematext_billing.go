package sematext

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	return nil // TODO Implement
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil // TODO Implement
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil // TODO Implement
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil // TODO Implement
}
