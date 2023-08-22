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
var _ resource.Resource = &AppClickhouseResource{}
var _ resource.ResourceWithImportState = &AppClickhouseResource{}


func NewAppClickhouseResource() resource.Resource {
	return &AppClickhouseResource{}
}

// AppClickhouseResourceModel describes the resource data model.
type AppClickhouseResourceModel = AppResourceModel


// AppClickhouseResource defines the resource implementation.
type AppClickhouseResource = AppResource


func (r *AppClickhouseResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppClickhouseResource"	
}


func (r *AppClickhouseResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("ClickHouse")
}


func (r *AppClickhouseResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Clickhouse")
}


func (r *AppClickhouseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	
	ResourceOpCreateApp(r, ctx, req, resp, "Clickhouse")

}


func (r *AppClickhouseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Clickhouse")

}

func (r *AppClickhouseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Clickhouse")

}

func (r *AppClickhouseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Clickhouse")

}

func (r *AppClickhouseResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Clickhouse")

}

