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
var _ resource.Resource = &AppAwsec2Resource{}
var _ resource.ResourceWithImportState = &AppAwsec2Resource{}


func NewAppAwsec2Resource() resource.Resource {
	return &AppAwsec2Resource{}
}

// AppAwsec2Resource defines the resource implementation.
type AppAwsec2Resource struct {
	client *stcloud.Configuration 
}


// AppAwsec2ResourceModel describes the resource data model.
type AppAwsec2ResourceModel struct {
	ResourceModel //TODO does this organisation work?
}


func (r *AppAwsec2Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = req.ProviderTypeName + "_AppAwsec2Resource"	

}


func (r *AppAwsec2Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	resp.Schema = ResourceSchemaApp("AWS EC2")
	
}


func (r *AppAwsec2Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	
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

	r.client = client
}

func (r *AppAwsec2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *AppAwsec2ResourceModel

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

func (r *AppAwsec2Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *AppAwsec2ResourceModel

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

func (r *AppAwsec2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *AppAwsec2ResourceModel

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

func (r *AppAwsec2Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *AppAwsec2ResourceModel

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

func (r *AppAwsec2Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

