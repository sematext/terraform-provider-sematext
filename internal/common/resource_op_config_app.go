package sematext

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

func Configure(resourceApp ResourceApp, ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse, appType string) {

	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*stcloud.APIClient) // @TODO push client into a struct

	if !ok {
		resp.Diagnostics.AddError(
			"Error: Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *stcloud.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	config = 


	diags := req.ProviderData
	
	.Get(ctx, &config)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.client = client

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" {

		// get AWS out of the environment
		var awsregion string
		var aws_access_key string
		var aws_secret_key string
		var ok bool

		if config.AwsAccessKey.IsUnknown() {
			aws_access_key, ok = os.LookupEnv("AWS_ACCESS_KEY_ID")
			if ok {
				config.AwsAccessKey
			} else {
				resp.Diagnostics.AddAttributeError(
					path.Root("aws_access_key"),
					"aws_access_key missing",
					"The provider cannot create or modify a Sematext Cloud app that requires AWS access (aws_access_key is missing or blank). "+
						"Either set the value statically in the resource configuration, or use the AWS_ACCESS_KEY_ID environment variable.",
				)
				return
			}
		}

		if config.AwsSecretKey.IsUnknown() {
			aws_secret_key, ok = os.LookupEnv("AWS_SECRET_ACCESS_KEY")
			if !ok {
				resp.Diagnostics.AddAttributeError(
					path.Root("aws_secret_key"),
					"aws_secret_key missing",
					"The provider cannot create or modify a Sematext Cloud app that requires AWS access (aws_access_key is missing or blank). "+
						"Either set the value statically in the resource configuration, or use the AWS_SECRET_ACCESS_KEY environment variable.",
				)
				return
			}

		}

		if config.AwsRegion.IsUnknown() {
			awsregion, ok = os.LookupEnv("AWS_REGION")
			if !ok {
				awsregion, ok = os.LookupEnv("AWS_DEFAULT_REGION")+++++++
			}
			if !ok {
				awsregion = "us-east-1"
			}
		}

	}

}
