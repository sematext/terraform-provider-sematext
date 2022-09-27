package sematext

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zclconf/go-cty/cty"
)

// ResourceSchemaAlertRule contains common resource fields
func ResourceSchemaAlertRule(appType string) map[string]*schema.Schema {

	allowedTriggers := map[string]string{
		"LESS":          "LESS",
		"MORE":          "MORE",
		"EQUAL":         "EQUAL",
		"UN_EQUAL":      "UN_EQUAL",
		"LESS_OR_EQUAL": "LESS_OR_EQUAL",
		"MORE_OR_EQUAL": "MORE_OR_EQUAL",
	}

	allowedNotifyApplicability := map[string]string{
		"NONE":       "NONE",
		"USE_ALWAYS": "USE_ALWAYS",
	}

	allowedNotifyType := map[string]string{
		"PAGER_DUTY":           "PAGER_DUTY",
		"NAGIOS":               "NAGIOS",
		"WEB_HOOKS":            "WEB_HOOKS",
		"WEB_HOOKS_TEMPLATE":   "WEB_HOOKS_TEMPLATE",
		"HIP_CHAT":             "HIP_CHAT",
		"EMAIL_LIST":           "EMAIL_LIST",
		"TEMPORARY_EMAIL_LIST": "TEMPORARY_EMAIL_LIST",
	}

	allowedStateType := map[string]string{
		"ACTIVE":   "ACTIVE",
		"DISABLED": "DISABLED",
		"DELETED":  "DELETED",
	}

	resourceSchema := map[string]*schema.Schema{

		"id": {
			Description: "(Computed) Internal id assigned by Sematext Cloud for this Alert Rule.",
			Type:        schema.TypeInt,
			Computed:    true,
			ForceNew:    true,
		},

		"name": {
			Description: "(Required) Name for this AlertRule.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				if len(value.(string)) == 0 {
					diag := diag.Diagnostic{}
					diag.Severity = diag.Error
					diag.Summary = "wrong value in name field"
					diag.Detail = fmt.Sprintf("string expected but got %q", value)
					diags = append(diags, diag)
				}
				return diags
			},
		},

		"description": {
			Description: "(Optional) Description of AlertRule.",
			Type:        schema.TypeString,
			Required:    false,
			ForceNew:    false,
			ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				if len(value.(string)) == 0 {
					diag := diag.Diagnostic{}
					diag.Severity = diag.Error
					diag.Summary = "wrong value in name field"
					diag.Detail = "description field is missing or empty"
					diags = append(diags, diag)
				}
				return diags
			},
		},

		"app": {
			Description: "(Required) Id for the Sematext Cloud App this AlertRule is associated with",
			Type:        schema.TypeString,
			Computed:    true,
			ForceNew:    true,
			ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				// TODO: enforce dependancy to an existing app?
				if len(value.(string)) == 0 {
					diag := diag.Diagnostic{}
					diag.Severity = diag.Error
					diag.Summary = "wrong value in name field"
					diag.Detail = "description field is missing or empty"
					diags = append(diags, diag)
				}
				return diags
			},
		},

		"backToNormalNeeded": {
			Description: "Set if a notfication is required for a back to normal status.",
			Type:        schema.TypeBool,
			Optional:    true,
			ForceNew:    true,
			Default:     true,
			ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				return diags
			},
		},

		"chartKey": {
			Description: "(Optional) Charting key legend for this AlertRule report.",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				return diags
			},
		},

		"color": {
			Description: "(Optional) Charting color for this AlertRule measurement.",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				// todo test legal colours
				return diags
			},
		},

		"enabled": {
			Description: "(Required) Set this AlertRule active.",
			Type:        schema.TypeBool,
			Required:    true,
			ForceNew:    true,
			Default:     true,
			ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				return diags
			},
		},

		"measure": {
			Description: "(Required) Set this AlertRule active.",
			Type:        schema.TypeList,
			Required:    true,
			ForceNew:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"trigger": {
						Description: "(Required) This AlertRule will trigger once measurement is LESS, MORE, EQUAL, UN_EQUAL, LESS_OR_EQUAL, MORE_OR_EQUAL.",
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							if trigger, ok := allowedTriggers[value.(string)]; !ok {
								diag := diag.Diagnostic{}
								diag.Severity = diag.Error
								diag.Summary = "bad value in measure.trigger field"
								diag.Detail = "measure.trigger is " + trigger + " should be one of LESS, MORE, EQUAL, UN_EQUAL, LESS_OR_EQUAL, MORE_OR_EQUAL"
								diags = append(diags, diag)
							}
							return diags

						},
					},

					"threshold": {
						Description: "(Required) AlertRule trigger threshold.",
						Type:        schema.TypeFloat,
						Required:    true,
						ForceNew:    true,
						ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics

							//todo validate measurement ranges

							return diags
						},
					},
				},
			},
			ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				// TODO validate set length = 1
				return diags
			},
		},

		"filter": {
			Description: "List of filters.",
			Type:        schema.TypeList,
			Optional:    true,
			ForceNew:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					"aggType": {
						Description: "TODO",
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							// TODO - looks like it should be an enum
							return diags
						},
					},

					"filterName": {
						Description: "TODO",
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
							var diags diag.Diagnostics
							// TODO
							return diags
						},
					},

					"key": {
						Description: "TODO",
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
							var diags diag.Diagnostics
							// TODO
							return diags
						},
					},

					"label": {
						Description: "TODO",
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
							var diags diag.Diagnostics
							// TODO
							return diags
						},
					},

					"name": {
						Description: "TODO",
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
							var diags diag.Diagnostics
							// TODO
							return diags
						},
					},

					"values": {
						Description: "TODO",
						Type:        schema.TypeList,
						Required:    true,
						ForceNew:    true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
							var diags diag.Diagnostics
							// TODO
							return diags
						},
					},
				},
			},
		},

		"ignoreRegularEvents": { // TODO - find out more about this, e.g. default
			Description: "Ignore regular events",
			Type:        schema.TypeBool,
			Optional:    true,
			ForceNew:    true,
			ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
				var diags diag.Diagnostics
				// TODO
				return diags
			},
		},

		"metadata": { // TODO - find out more about this
			Description: "Metadata - TODO",
			Type:        schema.TypeMap,
			Optional:    true,
			ForceNew:    true,
			ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
				var diags diag.Diagnostics
				// TODO - validate against apptype
				return diags
			},
		},

		"metricLabel": { // TODO
			Description: "Label for this metric..",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
				var diags diag.Diagnostics
				// TODO
				return diags
			},
		},

		"minDelayBetweenNotificationsInMinutes": { // TODO - default value?
			Description: "Minimum delay between notifications (minutes).",
			Type:        schema.TypeString, // TODO - should be TypeInt? Bug in Schema?
			Optional:    true,
			ForceNew:    true,
			ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
				var diags diag.Diagnostics
				// TODO
				return diags
			},
		},

		"notificationEmails": { // TODO
			Description: "TODO",
			Type:        schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
			ForceNew: false,
			ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
				var diags diag.Diagnostics
				// TODO
				return diags
			},
		},

		"notify": {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					"applicability": {
						Description: "Individual notify.applicability - should be one of NONE, USE_ALWAYS ",
						Type:        schema.TypeString,
						ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
							var diags diag.Diagnostics
							if applicability, ok := allowedNotifyApplicability[key]; !ok {
								diag := diag.Diagnostic{}
								diag.Severity = diag.Error
								diag.Summary = "bad value in notify.type field"
								diag.Detail = "notify.type is " + applicability + " should be one of NONE, USE_ALWAYS"
								diags = append(diags, diag)
							}
							return diags
						},
					},

					"type": {
						Description: "Individual notify.type - should be one one of PAGER_DUTY, NAGIOS, WEB_HOOKS, WEB_HOOKS_TEMPLATE, HIP_CHAT, EMAIL_LIST, TEMPORARY_EMAIL_LIST",
						Type:        schema.TypeString,
						ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
							var diags diag.Diagnostics
							if notifyType, ok := allowedNotifyType[key]; !ok {
								diag := diag.Diagnostic{}
								diag.Severity = diag.Error
								diag.Summary = "bad value in notify.type field"
								diag.Detail = "notify.type is set to " + notifyType + " should be one of one of PAGER_DUTY, NAGIOS, WEB_HOOKS, WEB_HOOKS_TEMPLATE, HIP_CHAT, EMAIL_LIST, TEMPORARY_EMAIL_LIST"
								diags = append(diags, diag)

							}
							return diags
						},
					},

					"name": { // TODO find out more about this
						Description: "TODO",
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
							// TODO - validate length
							return warns, errs
						},
					},

					"params": {
						Description: "TODO - highly variable - requires knowledge of integration type ... assume keyvalue pairs",
						Type:        schema.TypeMap,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Optional: true,
						ForceNew: false,
						ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
							var diags diag.Diagnostics
							// TODO - validate dependant on notify.type
							return diags
						},
					},

					"state": {
						Description: "TODO",
						Type:        schema.TypeString,
						Computed:    true,
						ForceNew:    false,
						ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
							var diags diag.Diagnostics
							if notifyState, ok := allowedStateType[key]; !ok {
								diag := diag.Diagnostic{}
								diag.Severity = diag.Error
								diag.Summary = "bad value in notify.state field"
								diag.Detail = "notify.state is set to " + notifyState + " should be one of one of ACTIVE, DISABLED or DELETED"
								diags = append(diags, diag)

							}
							return diags
						},
					},
				},
			},
		},

		"notificationsEnabled": { // TODO can this be set by the user or is this a system thing?
			Description: "TODO",
			Type:        schema.TypeBool,
			Required:    true,
			ForceNew:    true,
			ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
				var diags diag.Diagnostics
				return diags
			},
		},

		"query": { // TODO
			Description: "TODO",
			Type:        schema.TypeString,
			Required:    false,
			ForceNew:    true,
			ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
				var diags diag.Diagnostics
				return diags
			},
		},

		"reportName": {
			Description: "TODO",
			Type:        schema.TypeString,
			Required:    false,
			ForceNew:    true,
			ValidateDiagFunc: func(value interface{}, key string) (warns []string, errs []error) {
				var diags diag.Diagnostics
				return diags
			},
		},

		"ruleType": { // TODO
			Description: "TODO",
			Type:        schema.TypeString,
			Required:    false,
			ForceNew:    true,
			ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
				// TODO - validate enum [ ??? ]
				return warns, errs
			},
		},

		"runbook": { // TODO
			Description: "TODO",
			Type:        schema.TypeString,
			Required:    false,
			ForceNew:    true,
			ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
				// TODO - validate enum [ ??? ]
				return warns, errs
			},
		},

		"schedule": { // TODO array of AlertRuleScheduleWeekdayDto
			Description: "TODO",
			Type:        schema.TypeList,
			Optional:    true,
			ForceNew:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"day": {
						Description: "TODO",
						Type:        schema.TypeString,
						Required:    false,
						ForceNew:    true,
						ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
							// TODO - validate enum [ ??? ]
							return warns, errs
						},
					},

					"index": { // TODO - find out more about this
						Description: "TODO",
						Type:        schema.TypeInt,
						Required:    false,
						ForceNew:    true,
						ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
							// TODO - validate ???
							return warns, errs
						},
					},

					"interval": {
						Type:     schema.TypeList,
						Optional: true,
						ForceNew: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"start": {
									Type:     schema.TypeString,
									Required: false,
									ForceNew: true,
									ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
										// TODO - validate as datetime?
										return warns, errs
									},
								},
								"end": {
									Type:     schema.TypeString,
									Required: false,
									ForceNew: true,
									ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
										// TODO - validate as datetime?
										return warns, errs
									},
								},
							},
						},
					},

					"label": {
						Type:     schema.TypeString,
						Required: false,
						ForceNew: true,
						ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
							// TODO - validate string length
							return warns, errs
						},
					},

					"type": {
						Type:        schema.TypeString,
						Description: "TODO",
						Required:    false,
						ForceNew:    true,
						ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
							// TODO - validate enum [ ??? ]
							return warns, errs
						},
					},
				},
			},
		},

		"sendToEmail": { // TODO - find out more about this, smells obsolete.
			Description: "TODO",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
				// TODO - validate email address
				return warns, errs
			},
		},

		"timezone": { // TODO
			Description: "TODO",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
				// TODO - validate timezone enum
				return warns, errs
			},
		},

		"useOnlyAlertRuleIntegrations": {
			Description: "TODO",
			Type:        schema.TypeBool,
			Computed:    true,
			Default:     true,
		},
	}

	return resourceSchema

}
