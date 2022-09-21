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

// resourceOperationAppRedis is the resource class that handles sematext_app_redis
func resourceOperationAppRedis() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceOperationCreateAppRedis,
		ReadContext:   resourceOperationReadAppRedis,
		UpdateContext: resourceOperationUpdateAppRedis,
		DeleteContext: resourceOperationDeleteAppRedis,
		Schema:        sematext.ResourceSchemaApp("Redis"),
		Importer:      sematext.ResourceImporterApp("Redis"),
	}
}

// resourceOperationCreateAppRedis creates the sematext_app_redis resource.
func resourceOperationCreateAppRedis(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Redis"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppRedis reads the sematext_app_redis resource from Sematext Cloud.
func resourceOperationReadAppRedis(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Redis"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppRedis updates Sematext Cloud from the sematext_app_redis resource.
func resourceOperationUpdateAppRedis(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Redis"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppRedis marks a sematext_app_redis resource as retired.
func resourceOperationDeleteAppRedis(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Redis"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}
