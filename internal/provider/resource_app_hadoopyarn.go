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
var _ resource.Resource = &AppHadoopyarnResource{}
var _ resource.ResourceWithImportState = &AppHadoopyarnResource{}


func NewAppHadoopyarnResource() resource.Resource {
	return &AppHadoopyarnResource{}
}

// AppHadoopyarnResourceModel describes the resource data model.
type AppHadoopyarnResourceModel = AppResourceModel


// AppHadoopyarnResource defines the resource implementation.
type AppHadoopyarnResource = AppResource


func (r *AppHadoopyarnResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppHadoopyarnResource"	
}


func (r *AppHadoopyarnResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("Hadoop-YARN")
}


func (r *AppHadoopyarnResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Hadoopyarn")
}


func (r *AppHadoopyarnResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	
	ResourceOpCreateApp(r, ctx, req, resp, "Hadoopyarn")

}


func (r *AppHadoopyarnResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Hadoopyarn")

}

func (r *AppHadoopyarnResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Hadoopyarn")

}

func (r *AppHadoopyarnResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Hadoopyarn")

}

func (r *AppHadoopyarnResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Hadoopyarn")

}

