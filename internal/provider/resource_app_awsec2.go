package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_App.go.template
	Then run generate/generate.sh
*/

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppAwsec2Resource{}
var _ resource.ResourceWithImportState = &AppAwsec2Resource{}


func NewAppAwsec2Resource() resource.Resource {
	return &AppAwsec2Resource{}
}

// AppAwsec2ResourceModel describes the resource data model.
type AppAwsec2ResourceModel = AppResourceModel


// AppAwsec2Resource defines the resource implementation.
type AppAwsec2Resource = AppResource


func (r *AppAwsec2Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppAwsec2Resource"	
}


func (r *AppAwsec2Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("AWS EC2")
}


func (r *AppAwsec2Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Awsec2")
}


func (r *AppAwsec2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
		
	ResourceOpCreateApp(r, ctx, req, resp, "Awsec2")

}


func (r *AppAwsec2Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Awsec2")

}

func (r *AppAwsec2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Awsec2")

}

func (r *AppAwsec2Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Awsec2")

}

func (r *AppAwsec2Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Awsec2")

}

