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
var _ resource.Resource = &AppRedisResource{}
var _ resource.ResourceWithImportState = &AppRedisResource{}


func NewAppRedisResource() resource.Resource {
	return &AppRedisResource{}
}

// AppRedisResourceModel describes the resource data model.
type AppRedisResourceModel = AppResourceModel


// AppRedisResource defines the resource implementation.
type AppRedisResource = AppResource


func (r *AppRedisResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppRedisResource"	
}


func (r *AppRedisResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("Redis")
}


func (r *AppRedisResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Redis")
}


func (r *AppRedisResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
		
	ResourceOpCreateApp(r, ctx, req, resp, "Redis")

}


func (r *AppRedisResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Redis")

}

func (r *AppRedisResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Redis")

}

func (r *AppRedisResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Redis")

}

func (r *AppRedisResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Redis")

}

