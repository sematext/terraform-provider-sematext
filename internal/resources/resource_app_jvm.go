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


type AppJvmResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppJvmResource{}
var _ resource.ResourceWithImportState = &AppJvmResource{}


func NewAppJvmResource() resource.Resource {

	return &AppJvmResource{}

}


func (r *AppJvmResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_jvm"	

}


func (r *AppJvmResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "JVM"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("JVM")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("JVM")
	}

}


func (r *AppJvmResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "JVM"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Jvm")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Jvm")
	}

}


func (r *AppJvmResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "JVM"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Jvm")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Jvm")
	}

}


func (r *AppJvmResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "JVM"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Jvm")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Jvm")
	}

}


func (r *AppJvmResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "JVM"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Jvm")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Jvm")
	}

}


func (r *AppJvmResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "JVM"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Jvm")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Jvm")
	}
	

}


func (r *AppJvmResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "JVM"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Jvm")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Jvm")
	}

}

