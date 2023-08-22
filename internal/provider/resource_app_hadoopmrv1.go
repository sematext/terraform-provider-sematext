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
var _ resource.Resource = &AppHadoopmrv1Resource{}
var _ resource.ResourceWithImportState = &AppHadoopmrv1Resource{}


func NewAppHadoopmrv1Resource() resource.Resource {
	return &AppHadoopmrv1Resource{}
}

// AppHadoopmrv1ResourceModel describes the resource data model.
type AppHadoopmrv1ResourceModel = AppResourceModel


// AppHadoopmrv1Resource defines the resource implementation.
type AppHadoopmrv1Resource = AppResource


func (r *AppHadoopmrv1Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppHadoopmrv1Resource"	
}


func (r *AppHadoopmrv1Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("Hadoop-MRv1")
}


func (r *AppHadoopmrv1Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Hadoopmrv1")
}


func (r *AppHadoopmrv1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	
	ResourceOpCreateApp(r, ctx, req, resp, "Hadoopmrv1")

}


func (r *AppHadoopmrv1Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Hadoopmrv1")

}

func (r *AppHadoopmrv1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Hadoopmrv1")

}

func (r *AppHadoopmrv1Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Hadoopmrv1")

}

func (r *AppHadoopmrv1Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Hadoopmrv1")

}

