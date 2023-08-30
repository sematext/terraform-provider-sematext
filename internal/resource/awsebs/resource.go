package awsebs

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_App.go.template
	Then run generate/generate.sh
*/

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/sematext/sematext-api-client-go/stcloud"
	"github.com/sematext/terraform-provider-sematext/internal/common"
)


type Resource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Resource{}
var _ resource.ResourceWithImportState = &Resource{}


func NewResource() resource.Resource {

	return &Resource{}

}


func (r *Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_awsebs"	

}


func (r *Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "AWS EBS"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAWS("AWS EBS")
	} else{
		resp.Schema = common.ResourceSchemaDefault("AWS EBS")
	}

}


func (r *Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "AWS EBS"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAWS(r.client, ctx, req, resp, "Awsebs")
	} else {
		common.ResourceOpConfigureDefault(r.client, ctx, req, resp, "Awsebs")
	}

}


func (r *Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "AWS EBS"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAWS(r.client, ctx, req, resp, "Awsebs")
	} else{
		common.ResourceOpCreateDefault(r.client, ctx, req, resp, "Awsebs")
	}

}


func (r *Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "AWS EBS"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAWS(r.client, ctx, req, resp, "Awsebs")
	} else {
		common.ResourceOpReadDefault(r.client, ctx, req, resp, "Awsebs")
	}

}


func (r *Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "AWS EBS"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAWS(r.client, ctx, req, resp, "Awsebs")
	} else {
		common.ResourceOpUpdateDefault(r.client, ctx, req, resp, "Awsebs")
	}

}


func (r *Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "AWS EBS"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAWS(r.client, ctx, req, resp, "Awsebs")
	} else { 
		common.ResourceOpDeleteDefault(r.client, ctx, req, resp, "Awsebs")
	}
	

}


func (r *Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "AWS EBS"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAWS(r.client, ctx, req, resp, "Awsebs")
	} else {
		common.ResourceOpImportDefault(r.client, ctx, req, resp, "Awsebs")
	}

}

