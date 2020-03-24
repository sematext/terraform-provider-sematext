package sematext

import (
	"errors"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sematext/sematext-api-client/api"
)

// MonitorSchemaCommon contains common resource fields
func MonitorSchemaCommon(appType string) map[string]*schema.Schema {

	resourceSchema := map[string]*schema.Schema{

		"name": { // TODO validate func
			Description: "TODO",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    false,
			ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
				if len(value.(string)) == 0 {
					errs = append(errs, errors.New("name is missing or empty"))
				}
				return warns, errs
			},
		},

		"billing_plan_id": {
			Description: "TODO",
			Type:        schema.TypeInt,
			Required:    true,
			ForceNew:    false,
			ValidateFunc: func(planID interface{}, key string) (warns []string, errs []error) {

				if valid, found := api.LookupPlanID2Apptypes[planID.(int)]; found {
					if appType != valid {
						errs = append(errs, errors.New("billing_plan_id is not a valid billing code for resource type"))
					}
				} else {
					errs = append(errs, errors.New("billing_plan_id not recognized"))
				}

				return warns, errs
			},
		},

		"discount_code": {
			Description: "TODO",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    false,
			ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
				if len(value.(string)) == 0 {
					errs = append(errs, errors.New("discount_code is present but is empty"))
				}
				return warns, errs
			},
		},
	}

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { // TODO Pull out of env - look at how aws do this.

		resourceSchema["aws_access_key"] = &schema.Schema{
			Description: "TODO",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    false,
			ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
				if valid, _ := regexp.MatchString(`(?<![A-Z0-9])[A-Z0-9]{20}(?![A-Z0-9])`, value.(string)); !valid { // TODO check regex
					errs = append(errs, errors.New("Invalid aws_access_key"))
				}
				return warns, errs
			},
		}

		resourceSchema["aws_secret_key"] = &schema.Schema{
			Description: "TODO",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    false,
			ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
				if valid, _ := regexp.MatchString(`(?<![A-Za-z0-9/+=])[A-Za-z0-9/+=]{40}(?![A-Za-z0-9/+=])`, value.(string)); !valid { // TODO check regex
					errs = append(errs, errors.New("Invalid aws_secret_key"))
				}
				return warns, errs
			},
		}

		resourceSchema["aws_fetch_frequency"] = &schema.Schema{
			Description: "TODO",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    false,
			ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
				if valid, _ := regexp.MatchString(`(MINUTE|FIVE_MINUTES|FIFTEEN_MINUTES)`, value.(string)); !valid { // TODO check regex
					errs = append(errs, errors.New("Invalid aws_fetch_frequency"))
				}
				return warns, errs
			},
		}

		resourceSchema["aws_region"] = &schema.Schema{
			Description: "TODO",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    false,
			ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
				if _, found := api.LookupAWSRegion[value.(string)]; !found {
					errs = append(errs, errors.New("Invalid aws_region"))
				}
				return warns, errs
			},
			DefaultFunc: schema.MultiEnvDefaultFunc([]string{ // TODO check what this does
				"AWS_REGION",
				"AWS_DEFAULT_REGION",
			}, nil),
		}
	}

	return resourceSchema

}
