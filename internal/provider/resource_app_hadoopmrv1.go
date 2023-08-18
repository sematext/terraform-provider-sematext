package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_App.go.template
	Then run generate/generate.sh
*/

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/sematext/sematext-api-client-go/stcloud"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppHadoopmrv1Resource{}
var _ resource.ResourceWithImportState = &AppHadoopmrv1Resource{}


func NewAppHadoopmrv1Resource() resource.Resource {
	return &AppHadoopmrv1Resource{}
}

// AppHadoopmrv1ResourceModel describes the resource data model.
type AppHadoopmrv1ResourceModel = ResourceModel


// AppHadoopmrv1Resource defines the resource implementation.
type AppHadoopmrv1Resource struct {
	client *stcloud.Configuration 
}


func (r *AppHadoopmrv1Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppHadoopmrv1Resource"	
}


func (r *AppHadoopmrv1Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("Hadoop-MRv1")
}


func (r *AppHadoopmrv1Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*http.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	var config AppHadoopmrv1ResourceModel
    diags := req.Config.Get(ctx, &config)
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
		
		if config.AwsAccessKey.IsUnknown() {
			aws_access_key, ok = os.LookupEnv("AWS_ACCESS_KEY_ID")
			if ok {
				config.AwsAccessKey
			} else {
				resp.Diagnostics.AddAttributeError(
            		path.Root("aws_access_key"),
            			"aws_access_key missing from resource",
            			"The provider cannot create or modify a Sematext Cloud app that requires AWS access (aws_access_key is missing or blank). "+
						"Either set the value statically in the resource configuration, or use the AWS_ACCESS_KEY_ID environment variable.",
        		)
			} 
		}

		if config.AwsSecretKey.IsUnknown() {	
			aws_secret_key, ok = os.LookupEnv("AWS_SECRET_ACCESS_KEY")
			if !ok {
				resp.Diagnostics.AddAttributeError(
            		path.Root("aws_secret_key"),
            			"aws_access_key missing from resource",
            			"The provider cannot create or modify a Sematext Cloud app that requires AWS access (aws_access_key is missing or blank). "+
						"Either set the value statically in the resource configuration, or use the AWS_SECRET_ACCESS_KEY environment variable.",
        		)
			}
		}

		if resp.Diagnostics.HasError() {
        	return
    	}





		if config.AwsRegion.IsUnknown() {	
			awsregion, ok := os.LookupEnv("AWS_REGION")
			if !ok {
				awsregion, ok = os.LookupEnv("AWS_DEFAULT_REGION")
			}
			if !ok {
				awsregion = "us-east-1"
			}
		}

}
}

func (r *AppHadoopmrv1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AppHadoopmrv1ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create example, got error: %s", err))
	//     return
	// }

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	data.Id = types.StringValue("example-id")

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppHadoopmrv1Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AppHadoopmrv1ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}


	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppHadoopmrv1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AppHadoopmrv1ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update example, got error: %s", err))
	//     return
	// }

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppHadoopmrv1Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AppHadoopmrv1ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete example, got error: %s", err))
	//     return
	// }
}

func (r *AppHadoopmrv1Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

