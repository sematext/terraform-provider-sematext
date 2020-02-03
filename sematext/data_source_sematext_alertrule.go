package sematext

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

/*
	AlertRule

{
accountEmail	string
readOnly: true
allowedAppTypes	[...]
analyzingTime	string
appDisplayState	string
readOnly: true
appId	integer($int64)
appName	string
readOnly: true
appState	string
readOnly: true
appToken	string
readOnly: true
appType	string
readOnly: true
backToNormalNeeded	boolean
chartKey	string
color	string
creatorEmail	string
readOnly: true
defaultAggType	string
readOnly: true
description	string
disallowedAppTypes	[integer($int64)]
enabled	boolean
estimateOperation	string
Enum:
[ LESS, MORE, EQUAL, UN_EQUAL, LESS_OR_EQUAL, MORE_OR_EQUAL ]
estimateValue	number($double)
filterValues	string
filterValuesObj	[FilterValue{...}]
ignoreRegularEventsEnabled	boolean
integrations	string
readOnly: true
lastDataReceivedDate	integer($int64)
readOnly: true
lastSent	integer($int64)
readOnly: true
lastTriggered	integer($int64)
readOnly: true
metadata	{
}
metricKey	string
readOnly: true
metricLabel	string
minDelayBetweenNotificationsInMinutes	string
name	string
notificationEmails	[string]
notificationIntegrations	[NotificationIntegration{...}]
notificationsEnabled	boolean
query	string
reportName	string
ruleKey	integer($int64)
readOnly: true
ruleType	string
runbook	string
savedQueryId	integer($int64)
readOnly: true
schedule	[AlertRuleScheduleWeekdayDto{...}]
sematextService	string
readOnly: true
sendToEmail	string
timezone	string
useOnlyAlertRuleIntegrations	boolean
userPermissions	UserPermissions{
canDelete	boolean
canEdit	boolean
canView	boolean
}
valueColumnName	string
readOnly: true
valueName	string
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
