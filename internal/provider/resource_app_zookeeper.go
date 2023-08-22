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
var _ resource.Resource = &AppZookeeperResource{}
var _ resource.ResourceWithImportState = &AppZookeeperResource{}


func NewAppZookeeperResource() resource.Resource {
	return &AppZookeeperResource{}
}

// AppZookeeperResourceModel describes the resource data model.
type AppZookeeperResourceModel = AppResourceModel


// AppZookeeperResource defines the resource implementation.
type AppZookeeperResource = AppResource


func (r *AppZookeeperResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppZookeeperResource"	
}


func (r *AppZookeeperResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("ZooKeeper")
}


func (r *AppZookeeperResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Zookeeper")
}


func (r *AppZookeeperResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	
	ResourceOpCreateApp(r, ctx, req, resp, "Zookeeper")

}


func (r *AppZookeeperResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Zookeeper")

}

func (r *AppZookeeperResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Zookeeper")

}

func (r *AppZookeeperResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Zookeeper")

}

func (r *AppZookeeperResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Zookeeper")

}

