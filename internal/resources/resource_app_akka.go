package resources

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


type AppAkkaResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppAkkaResource{}
var _ resource.ResourceWithImportState = &AppAkkaResource{}


func NewAppAkkaResource() resource.Resource {

	return &AppAkkaResource{}

}


func (r *AppAkkaResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_akka"	

}


func (r *AppAkkaResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Akka"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Akka")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Akka")
	}

}


func (r *AppAkkaResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Akka"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Akka")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Akka")
	}

}


func (r *AppAkkaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Akka"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Akka")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Akka")
	}

}


func (r *AppAkkaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Akka"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Akka")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Akka")
	}

}


func (r *AppAkkaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Akka"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Akka")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Akka")
	}

}


func (r *AppAkkaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Akka"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Akka")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Akka")
	}
	

}


func (r *AppAkkaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Akka"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Akka")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Akka")
	}

}

