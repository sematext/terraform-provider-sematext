package sematext

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

// ResourceSchemaApp contains common resource fields
func ResourceSchemaMonitoringApp(appType string) map[string]*schema.Schema {

	return schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: appType, // TODO complete string

		Attributes: map[string]schema.Attribute{

			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Example identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},

			"name": {
				Description: "(Required) Label for the monitor app in Sematext Cloud.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
					if len(value.(string)) == 0 {
						errs = append(errs, errors.New("name field is missing or empty"))
					}
					return warns, errs
				},
			},
	
			"billing_plan_id": {
				Description: "(Required) Plan ID attached to the monitor app in Sematext Cloud.",
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    false,
				ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
	
					planID := value.(int)
	
					if appTypeForPlanID, found := stcloud.LookupPlanID2Apptypes[planID]; found {
						if appType != appTypeForPlanID {
							errs = append(errs, fmt.Errorf("billing_plan_id %d is not a valid billing code for resource type %s", planID, appType))
						}
					} else {
						errs = append(errs, fmt.Errorf("billing_plan_id %d is not rea valid billing code for resource type %s", planID, appType))
					}
	
					return warns, errs
	
				},
			},
	
			"discount_code": {
				Description: "(Optional) Iniitial discount code attached to the monitor app in Sematext Cloud.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    false,
				ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
					return warns, errs
				},
			},
	
			// At this point the association is with all app-tokens in the SC app. Association of individual app-token with container instance is set in the provisioner(s).
			"apptoken": {
				Description: "(required) apptoken(s) attached to the monitor app in Sematext Cloud.",
				Type:        schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"names": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Required: true,
							ForceNew: false,
						},
					},
				},
				Required: true,
				ForceNew: false,
			},
	
			/*
				// Above
				// Once the app is created this is pulled back from SC and lives only in state file. This is checked each time to see ids line up with names.
				"apptoken_entries": {
					Description: "One or more apptoken entries. Calculated, supplied by SC Cloud.",
					Type:        schema.TypeSet,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							
							"CreatedAt": {
								Type:     schema.TypeString,
								Computed: true,
								ForceNew: false,
							},
							"Enabled": {
								Type:     schema.TypeBool,
								Computed: true,
								ForceNew: false,
							},
							"ID": {
								Type:     schema.TypeInt,
								Computed: true,
								ForceNew: false,
							},
							"Name": {
								Type:     schema.TypeString,
								Computed: true,
								ForceNew: false,
							},
							"Readable": {
								Type:     schema.TypeBool,
								Computed: true,
								ForceNew: false,
							},
							"Token": {
								Type:     schema.TypeString,
								Computed: true,
								ForceNew: false,
							},
							"Writeable": {
								Type:     schema.TypeBool,
								Computed: true,
								ForceNew: false,
							},
						},
					},
					Computed: true,
					ForceNew: false,
				},
			*/
	
			// Once the app is created this is pulled back from SC and lives only in state file. This is checked each time to see ids line up with names.
			// Note : Simplified from above to make lookup more readible in scripts that use this.
			"sc_apptoken_entries": {
				Description: "Map of apptoken name -> id. Calculated, supplied by SC Cloud.",
				Type:        schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
				ForceNew: false,
			},
		}
	
		if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" {
	
			// TODO Pull aws credentials out of .config in OS dependant manner, override with env.
	
			resourceSchema["aws_access_key"] = &schema.Schema{
				Description: "The access key for retrieval of stats from AWS Cloudwatch. You can retrieve this\nfrom the 'Security & Credentials' section of the AWS console.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
					if matched, err := regexp.MatchString(`([0-9A-Za-z]{20})`, value.(string)); !matched {
						errs = append(errs, errors.New("invalid aws_access_key"))
						if err != nil {
							errs = append(errs, err)
						}
					}
					return warns, errs
				},
			}
	
			resourceSchema["aws_secret_key"] = &schema.Schema{
				Description: "The secret key for retrieval of stats from AWS Cloudwatch. You can retrieve this\nfrom the 'Security & Credentials' section of the AWS console.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
					if matched, err := regexp.MatchString(`([0-9A-Za-z+/=]{40})`, value.(string)); !matched {
						errs = append(errs, errors.New("invalid aws_secret_key"))
						if err != nil {
							errs = append(errs, err)
						}
	
					}
					return warns, errs
				},
			}
	
			resourceSchema["aws_fetch_frequency"] = &schema.Schema{
				Description: "How frequently to fetch metrics. One of MINUTE|FIVE_MINUTES|FIFTEEN_MINUTES",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
					if valid, _ := regexp.MatchString(`(MINUTE|FIVE_MINUTES|FIFTEEN_MINUTES)`, value.(string)); !valid {
						errs = append(errs, errors.New("invalid aws_fetch_frequency"))
					}
					return warns, errs
				},
			}
	
			resourceSchema["aws_region"] = &schema.Schema{
				Description: "The region where AWS operations will take place. Examples\nare us-east-1, us-west-2, etc.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
					if _, found := stcloud.AWSRegion2STRegion[value.(string)]; !found {
						errs = append(errs, errors.New("invalid aws_region"))
					}
					return warns, errs
				},
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"AWS_REGION",
					"AWS_DEFAULT_REGION",
				}, nil),
				InputDefault: "us-east-1",
			}
		},

	}

}



