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
var _ resource.Resource = &AppLogseneResource{}
var _ resource.ResourceWithImportState = &AppLogseneResource{}


func NewAppLogseneResource() resource.Resource {
	return &AppLogseneResource{}
}

// AppLogseneResourceModel describes the resource data model.
type AppLogseneResourceModel = AppResourceModel


// AppLogseneResource defines the resource implementation.
type AppLogseneResource = AppResource


func (r *AppLogseneResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppLogseneResource"	
}


func (r *AppLogseneResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("Logsene")
}


func (r *AppLogseneResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Logsene")
}


func (r *AppLogseneResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	
	ResourceOpCreateApp(r, ctx, req, resp, "Logsene")

}


func (r *AppLogseneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Logsene")

}

func (r *AppLogseneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Logsene")

}

func (r *AppLogseneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Logsene")

}

func (r *AppLogseneResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Logsene")

}

