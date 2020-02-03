package sematext

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

/*
UpdateAppInfo
{
description	string
example: New Description of App
ignorePercentage	integer($int32)
example: 1
maxEvents	integer($int64)
example: 1000
maxLimitMB	integer($int64)
example: 10
name	string
example: New Name
sampling	boolean
example: false
samplingPercentage	integer($int32)
example: 10
staggering	boolean
example: false
status	string
example: ACTIVE
Enum:
[ ACTIVE, DISABLED ]
}
*/

func dataSourceTodoPlaceholder() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTodoPlaceholderRead,

		Schema: map[string]*schema.Schema{
			"field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"nested": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},

						"field2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},

						"field3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},

						"field4": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},

			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceTodoPlaceholderRead(d *schema.ResourceData, meta interface{}) error {

	// TODO Implement.

	return nil
}
