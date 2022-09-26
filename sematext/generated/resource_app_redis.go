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

// resourceAppRedis is the resource class that handles sematext_app_redis
func resourceAppRedis() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("Redis"),
		CreateContext: resourceOperationCreateAppRedis,
		ReadContext:   resourceOperationReadAppRedis,
		UpdateContext: resourceOperationUpdateAppRedis,
		DeleteContext: resourceOperationDeleteAppRedis,
		Importer:      resourceOperationImportAppRedis(),
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

// resourceOperationImportAppRedis imports a sematext_app_redis resource into the state file.
func resourceOperationImportAppRedis() *schema.ResourceImporter {
	apptype := "Redis"
	return sematext.ResourceOperationImportApp(apptype)
}
