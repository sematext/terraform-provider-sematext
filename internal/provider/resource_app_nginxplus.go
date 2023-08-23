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
var _ resource.Resource = &AppNginxplusResource{}
var _ resource.ResourceWithImportState = &AppNginxplusResource{}


func NewAppNginxplusResource() resource.Resource {
	return &AppNginxplusResource{}
}

// AppNginxplusResourceModel describes the resource data model.
type AppNginxplusResourceModel = AppResourceModel


// AppNginxplusResource defines the resource implementation.
type AppNginxplusResource = AppResource


func (r *AppNginxplusResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppNginxplusResource"	
}


func (r *AppNginxplusResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("Nginx-Plus")
}


func (r *AppNginxplusResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Nginxplus")
}


func (r *AppNginxplusResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
		
	ResourceOpCreateApp(r, ctx, req, resp, "Nginxplus")

}


func (r *AppNginxplusResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Nginxplus")

}

func (r *AppNginxplusResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Nginxplus")

}

func (r *AppNginxplusResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Nginxplus")

}

func (r *AppNginxplusResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Nginxplus")

}

