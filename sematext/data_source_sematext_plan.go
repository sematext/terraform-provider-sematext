package sematext

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

/*
Plan
{
appType	string
custom	boolean
dataRetentionHours	number
defaultTrialPlan	boolean
free	boolean
freeTrialDays	integer($int64)
id	integer($int64)
maxAlerts	integer($int64)
maxDailyEvents	integer($int64)
name	string
planScheme	string
Enum:
[ SPM_1_0, SPM_2_0, SA_1_0, SEARCHENE_1_0, LOGSENE_1_0, LOGSENE_2_0, RUM_1_0, RUM_EA, SYNTHETICS_1_0 ]
sematextService	string
Enum:
[ LOGSENE, SPM, RUM, SYNTHETICS ]
trialPlan	boolean
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
