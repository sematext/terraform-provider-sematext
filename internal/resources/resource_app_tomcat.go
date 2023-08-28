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


type AppTomcatResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppTomcatResource{}
var _ resource.ResourceWithImportState = &AppTomcatResource{}


func NewAppTomcatResource() resource.Resource {

	return &AppTomcatResource{}

}


func (r *AppTomcatResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_tomcat"	

}


func (r *AppTomcatResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Tomcat"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Tomcat")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Tomcat")
	}

}


func (r *AppTomcatResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Tomcat"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Tomcat")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Tomcat")
	}

}


func (r *AppTomcatResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Tomcat"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Tomcat")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Tomcat")
	}

}


func (r *AppTomcatResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Tomcat"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Tomcat")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Tomcat")
	}

}


func (r *AppTomcatResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Tomcat"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Tomcat")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Tomcat")
	}

}


func (r *AppTomcatResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Tomcat"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Tomcat")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Tomcat")
	}
	

}


func (r *AppTomcatResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Tomcat"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Tomcat")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Tomcat")
	}

}

