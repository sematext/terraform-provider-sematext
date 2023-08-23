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
var _ resource.Resource = &AppMongodbResource{}
var _ resource.ResourceWithImportState = &AppMongodbResource{}


func NewAppMongodbResource() resource.Resource {
	return &AppMongodbResource{}
}

// AppMongodbResourceModel describes the resource data model.
type AppMongodbResourceModel = AppResourceModel


// AppMongodbResource defines the resource implementation.
type AppMongodbResource = AppResource


func (r *AppMongodbResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppMongodbResource"	
}


func (r *AppMongodbResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("MongoDB")
}


func (r *AppMongodbResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Mongodb")
}


func (r *AppMongodbResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
		
	ResourceOpCreateApp(r, ctx, req, resp, "Mongodb")

}


func (r *AppMongodbResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Mongodb")

}

func (r *AppMongodbResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Mongodb")

}

func (r *AppMongodbResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Mongodb")

}

func (r *AppMongodbResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Mongodb")

}

