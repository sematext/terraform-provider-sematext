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
var _ resource.Resource = &AppAkkaResource{}
var _ resource.ResourceWithImportState = &AppAkkaResource{}


func NewAppAkkaResource() resource.Resource {
	return &AppAkkaResource{}
}

// AppAkkaResourceModel describes the resource data model.
type AppAkkaResourceModel = AppResourceModel


// AppAkkaResource defines the resource implementation.
type AppAkkaResource = AppResource


func (r *AppAkkaResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppAkkaResource"	
}


func (r *AppAkkaResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("Akka")
}


func (r *AppAkkaResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Akka")
}


func (r *AppAkkaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
		
	ResourceOpCreateApp(r, ctx, req, resp, "Akka")

}


func (r *AppAkkaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Akka")

}

func (r *AppAkkaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Akka")

}

func (r *AppAkkaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Akka")

}

func (r *AppAkkaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Akka")

}

