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
var _ resource.Resource = &AppMobilelogsResource{}
var _ resource.ResourceWithImportState = &AppMobilelogsResource{}


func NewAppMobilelogsResource() resource.Resource {
	return &AppMobilelogsResource{}
}

// AppMobilelogsResourceModel describes the resource data model.
type AppMobilelogsResourceModel = AppResourceModel


// AppMobilelogsResource defines the resource implementation.
type AppMobilelogsResource = AppResource


func (r *AppMobilelogsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppMobilelogsResource"	
}


func (r *AppMobilelogsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("mobile-logs")
}


func (r *AppMobilelogsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Mobilelogs")
}


func (r *AppMobilelogsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
		
	ResourceOpCreateApp(r, ctx, req, resp, "Mobilelogs")

}


func (r *AppMobilelogsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Mobilelogs")

}

func (r *AppMobilelogsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Mobilelogs")

}

func (r *AppMobilelogsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Mobilelogs")

}

func (r *AppMobilelogsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Mobilelogs")

}

