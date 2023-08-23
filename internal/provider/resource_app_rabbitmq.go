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
var _ resource.Resource = &AppRabbitmqResource{}
var _ resource.ResourceWithImportState = &AppRabbitmqResource{}


func NewAppRabbitmqResource() resource.Resource {
	return &AppRabbitmqResource{}
}

// AppRabbitmqResourceModel describes the resource data model.
type AppRabbitmqResourceModel = AppResourceModel


// AppRabbitmqResource defines the resource implementation.
type AppRabbitmqResource = AppResource


func (r *AppRabbitmqResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppRabbitmqResource"	
}


func (r *AppRabbitmqResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("rabbitmq")
}


func (r *AppRabbitmqResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Rabbitmq")
}


func (r *AppRabbitmqResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
		
	ResourceOpCreateApp(r, ctx, req, resp, "Rabbitmq")

}


func (r *AppRabbitmqResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Rabbitmq")

}

func (r *AppRabbitmqResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Rabbitmq")

}

func (r *AppRabbitmqResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Rabbitmq")

}

func (r *AppRabbitmqResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Rabbitmq")

}

