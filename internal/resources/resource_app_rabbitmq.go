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


type AppRabbitmqResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppRabbitmqResource{}
var _ resource.ResourceWithImportState = &AppRabbitmqResource{}


func NewAppRabbitmqResource() resource.Resource {

	return &AppRabbitmqResource{}

}


func (r *AppRabbitmqResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_rabbitmq"	

}


func (r *AppRabbitmqResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "rabbitmq"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("rabbitmq")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("rabbitmq")
	}

}


func (r *AppRabbitmqResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "rabbitmq"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Rabbitmq")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Rabbitmq")
	}

}


func (r *AppRabbitmqResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "rabbitmq"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Rabbitmq")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Rabbitmq")
	}

}


func (r *AppRabbitmqResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "rabbitmq"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Rabbitmq")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Rabbitmq")
	}

}


func (r *AppRabbitmqResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "rabbitmq"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Rabbitmq")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Rabbitmq")
	}

}


func (r *AppRabbitmqResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "rabbitmq"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Rabbitmq")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Rabbitmq")
	}
	

}


func (r *AppRabbitmqResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "rabbitmq"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Rabbitmq")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Rabbitmq")
	}

}

