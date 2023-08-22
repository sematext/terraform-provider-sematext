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
var _ resource.Resource = &AppAwsebsResource{}
var _ resource.ResourceWithImportState = &AppAwsebsResource{}


func NewAppAwsebsResource() resource.Resource {
	return &AppAwsebsResource{}
}

// AppAwsebsResourceModel describes the resource data model.
type AppAwsebsResourceModel = AppResourceModel


// AppAwsebsResource defines the resource implementation.
type AppAwsebsResource = AppResource


func (r *AppAwsebsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppAwsebsResource"	
}


func (r *AppAwsebsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("AWS EBS")
}


func (r *AppAwsebsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Awsebs")
}


func (r *AppAwsebsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	
	ResourceOpCreateApp(r, ctx, req, resp, "Awsebs")

}


func (r *AppAwsebsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Awsebs")

}

func (r *AppAwsebsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Awsebs")

}

func (r *AppAwsebsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Awsebs")

}

func (r *AppAwsebsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Awsebs")

}

