package sematext

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sematext/sematext-api-client/stcloud"
)

// MonitorSchemaCommon contains common resource fields
func MonitorSchemaCommon(appType string) map[string]*schema.Schema {

	resourceSchema := map[string]*schema.Schema{

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
				fmt.Println(value.(string))
				if matched, err := regexp.MatchString(`([0-9A-Za-z]{20})`, value.(string)); !matched {
					errs = append(errs, errors.New("Invalid aws_access_key"))
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
				fmt.Println(value.(string))
				if matched, err := regexp.MatchString(`([0-9A-Za-z+/=]{40})`, value.(string)); !matched {
					errs = append(errs, errors.New("Invalid aws_secret_key"))
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
				fmt.Println(value.(string))
				if valid, _ := regexp.MatchString(`(MINUTE|FIVE_MINUTES|FIFTEEN_MINUTES)`, value.(string)); !valid {
					errs = append(errs, errors.New("Invalid aws_fetch_frequency"))
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
				fmt.Println(value.(string))
				if _, found := stcloud.AWSRegion2STRegion[value.(string)]; !found {
					errs = append(errs, errors.New("Invalid aws_region"))
				}
				return warns, errs
			},
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{
				"AWS_REGION",
				"AWS_DEFAULT_REGION",
			}, nil),
			InputDefault: "us-east-1",
		}
	}

	return resourceSchema

}
