package generated

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_App.go.template
	Then run generate/generate.sh
*/

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sematext/terraform-provider-sematext/sematext"
)

// resourceOperationAppZookeeper is the resource class that handles sematext_app_zookeeper
func resourceOperationAppZookeeper() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceOperationCreateAppZookeeper,
		ReadContext:   resourceOperationReadAppZookeeper,
		UpdateContext: resourceOperationUpdateAppZookeeper,
		DeleteContext: resourceOperationDeleteAppZookeeper,
		Schema:        sematext.ResourceSchemaApp("ZooKeeper"),
		Importer:      sematext.ResourceImporterApp("ZooKeeper"),
	}
}

// resourceOperationCreateAppZookeeper creates the sematext_app_zookeeper resource.
func resourceOperationCreateAppZookeeper(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "ZooKeeper"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppZookeeper reads the sematext_app_zookeeper resource from Sematext Cloud.
func resourceOperationReadAppZookeeper(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "ZooKeeper"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppZookeeper updates Sematext Cloud from the sematext_app_zookeeper resource.
func resourceOperationUpdateAppZookeeper(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "ZooKeeper"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppZookeeper marks a sematext_app_zookeeper resource as retired.
func resourceOperationDeleteAppZookeeper(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "ZooKeeper"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}
