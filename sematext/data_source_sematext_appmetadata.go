package sematext

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

/*

AppMetaData

{
awsCloudWatchAccessKey	string
example: zzzzz
awsCloudWatchSecretKey	string
example: xxxxx
awsFetchFrequency	string
example: FIVE_MINUTES
Enum:
[ MINUTE, FIVE_MINUTES, FIFTEEN_MINUTES ]
awsRegion	string
example: US_EAST_1
Enum:
[ US_EAST_1, US_WEST_1, EU_WEST_1, US_WEST_2, AP_SOUTHEAST_1, AP_SOUTHEAST_2, AP_NORTHEAST_1, SA_EAST_1, GovCloud, CN_NORTH_1, US_EAST_2, AP_SOUTH_1, AP_NORTHEAST_2, CA_CENTRAL_1, EU_CENTRAL_1, EU_WEST_2 ]
subTypes	[
example: aws_ec2,aws_elb
Comma separated list of AWS types monitored by created app

string
Enum:
[ aws_ec2, aws_elb, aws_ebs ]
]
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
