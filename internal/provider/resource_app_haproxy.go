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
var _ resource.Resource = &AppHaproxyResource{}
var _ resource.ResourceWithImportState = &AppHaproxyResource{}


func NewAppHaproxyResource() resource.Resource {
	return &AppHaproxyResource{}
}

// AppHaproxyResourceModel describes the resource data model.
type AppHaproxyResourceModel = AppResourceModel


// AppHaproxyResource defines the resource implementation.
type AppHaproxyResource = AppResource


func (r *AppHaproxyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppHaproxyResource"	
}


func (r *AppHaproxyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("HAProxy")
}


func (r *AppHaproxyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Haproxy")
}


func (r *AppHaproxyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	
	ResourceOpCreateApp(r, ctx, req, resp, "Haproxy")

}


func (r *AppHaproxyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Haproxy")

}

func (r *AppHaproxyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Haproxy")

}

func (r *AppHaproxyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Haproxy")

}

func (r *AppHaproxyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Haproxy")

}

