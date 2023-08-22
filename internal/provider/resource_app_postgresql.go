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
var _ resource.Resource = &AppPostgresqlResource{}
var _ resource.ResourceWithImportState = &AppPostgresqlResource{}


func NewAppPostgresqlResource() resource.Resource {
	return &AppPostgresqlResource{}
}

// AppPostgresqlResourceModel describes the resource data model.
type AppPostgresqlResourceModel = AppResourceModel


// AppPostgresqlResource defines the resource implementation.
type AppPostgresqlResource = AppResource


func (r *AppPostgresqlResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppPostgresqlResource"	
}


func (r *AppPostgresqlResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("postgresql")
}


func (r *AppPostgresqlResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Postgresql")
}


func (r *AppPostgresqlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	
	ResourceOpCreateApp(r, ctx, req, resp, "Postgresql")

}


func (r *AppPostgresqlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Postgresql")

}

func (r *AppPostgresqlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Postgresql")

}

func (r *AppPostgresqlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Postgresql")

}

func (r *AppPostgresqlResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Postgresql")

}

