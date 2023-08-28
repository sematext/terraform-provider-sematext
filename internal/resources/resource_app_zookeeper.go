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


type AppZookeeperResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppZookeeperResource{}
var _ resource.ResourceWithImportState = &AppZookeeperResource{}


func NewAppZookeeperResource() resource.Resource {

	return &AppZookeeperResource{}

}


func (r *AppZookeeperResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_zookeeper"	

}


func (r *AppZookeeperResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "ZooKeeper"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("ZooKeeper")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("ZooKeeper")
	}

}


func (r *AppZookeeperResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "ZooKeeper"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Zookeeper")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Zookeeper")
	}

}


func (r *AppZookeeperResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "ZooKeeper"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Zookeeper")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Zookeeper")
	}

}


func (r *AppZookeeperResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "ZooKeeper"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Zookeeper")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Zookeeper")
	}

}


func (r *AppZookeeperResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "ZooKeeper"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Zookeeper")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Zookeeper")
	}

}


func (r *AppZookeeperResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "ZooKeeper"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Zookeeper")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Zookeeper")
	}
	

}


func (r *AppZookeeperResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "ZooKeeper"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Zookeeper")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Zookeeper")
	}

}

