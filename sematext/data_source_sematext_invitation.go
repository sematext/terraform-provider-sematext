package sematext

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)
/*
Invitation
{
app	App{
ajaxThreshold	integer($int64)
appType	string
appTypeId	integer($int64)
creatorEmail	string
creditCardExpiry	string
creditCardNumber	string
description	string
displayStatus	string
firstDataSavedDate	integer($int64)
id	integer($int64)
integration	ServiceIntegration{
appTypeId	integer($int64)
appTypeName	string
displayName	string
enabled	boolean
externalProductId	integer($int64)
externalProductName	string
id	integer($int64)
integrationType	string
ordinal	integer($int32)
parentIntegrationId	integer($int64)
sematextService	string
visible	boolean
}
lastDataReceivedDate	integer($int64)
lastDataSavedDate	integer($int64)
loggedInUserAppRole	string
monthlyInvoiceAccount	boolean
name	string
ownerEmail	string
owningOrganization	BasicOrganizationDto{
authMethods	[BasicAuthMethodDto{
authType	string
Enum:
[ SEMATEXT_ACCOUNT, LDAP ]
uuid	string
}]
name	string
status	string
Enum:
[ ACTIVE, IN_REGISTRATION, DISABLED, EXPIRED, INVITED, DEMO ]
uuid	string
}
pageLoadThreshold	integer($int64)
paymentMethodId	integer($int64)
plan	Plan{
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
prepaidAccount	boolean
status	string
token	string
trialEndDate	integer($int64)
urlGroupLimit	integer($int32)
userRoles	[UserRole{
role	string
Enum:
Array [ 6 ]
roleStatus	string
Enum:
Array [ 2 ]
userEmail	string
}]
}
apps	[...]
id	integer($int64)
readOnly: true
inviteDate	string($date-time)
readOnly: true
inviteStatus	string
readOnly: true
Enum:
Array [ 4 ]
inviteeEmail	string
example: guest@sematext.com
inviteeRole	string
example: DEMO
Enum:
Array [ 6 ]
inviteeStatus	string
example: ACTIVE
readOnly: true
Enum:
Array [ 2 ]
inviterEmail	string
readOnly: true
uuid	string
readOnly: true
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
