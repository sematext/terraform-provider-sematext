package sematext

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)
/*
ReportInfo
{
addresses	string
Comma separated list of email addresses

appId	integer($int64)
endDate	string($date-time)
filters	string
reportName	string
startDate	string($date-time)
subject	string
text	string
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
