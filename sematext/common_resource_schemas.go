package sematext

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	funk "github.com/thoas/go-funk"
)

// MonitorSchemaCommon contains common resource fields
var MonitorSchemaCommon = map[string]*schema.Schema{

	"name": { // TODO validate func
		Type:     schema.TypeString,
		Required: true,
		ForceNew: false,
		ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
			if len(value.(string)) == 0 {
				errs = append(errs, errors.New("name is missing or empty"))
			}
			return warns, errs
		},
	},
	"description": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: false,
		ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
			if len(value.(string)) == 0 {
				errs = append(errs, errors.New("description is missing or empty"))
			}
			return warns, errs
		},
	},
	"billing_plan": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: false,
		ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
			allowed := []string{"basic", "standard", "pro"}
			if !funk.Contains(allowed, value.(string)) {
				errs = append(errs, fmt.Errorf("billing_plan must be one of %v : got %s", allowed, value))
			}
			return warns, errs
		},
	},
	"discount_code": {
		Type:     schema.TypeString,
		Optional: true,
		ForceNew: false,
		ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
			if len(value.(string)) == 0 {
				errs = append(errs, errors.New("discount_code is present but is empty"))
			}
			return warns, errs
		},
	},
	"ignore_percentage": {
		Type:     schema.TypeInt,
		Optional: true,
		ForceNew: false,
		ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
			// TODO Validate func on ignore_percentage
			return warns, errs
		},
	},
	"max_events": {
		Type:     schema.TypeInt,
		Optional: true,
		ForceNew: false,
		ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
			// TODO Validate func on max_events
			return warns, errs
		},
	},
	"max_limit_mb": {
		Type:     schema.TypeInt,
		Optional: true,
		ForceNew: false,
		ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
			// TODO Validate func on max_limit_mb
			return warns, errs
		},
	},
	"sampling": {
		Type:     schema.TypeBool,
		Optional: true,
		ForceNew: false,
		ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
			// TODO Validate func on sampling
			return warns, errs
		},
	},
	"sampling_percentage": {
		Type:     schema.TypeInt,
		Optional: true,
		ForceNew: false,
		ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
			// TODO Validate func on sampling_percentage
			return warns, errs
		},
	},
	"staggering": {
		Type:     schema.TypeBool,
		Optional: true,
		ForceNew: false,
		ValidateFunc: func(value interface{}, key string) (warns []string, errs []error) {
			// TODO Validate func on staggering
			return warns, errs
		},
	},
}

/*
var MonitorMetadataAWS = &schema.Schema{
	Optional: true,
	ForceNew: true,
	Type: schema.Schema{
		"aws_cloudwatch": &schema.Schema{ // TODO Pull out of ENV?
			Required: true,
			ForceNew: true,
			Type: &schema.Schema{
				"access_key": {
					Type:     schema.TypeString,
					Required: true,
				},
				"secret_key": {
					Type:     schema.TypeString,
					Required: true,
				},
				"fetch_frequency": {
					Type:     schema.TypeString,
					Required: true,
				},
				"region": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	},
}

*/
