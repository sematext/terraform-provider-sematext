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

// resourceAppMongodb is the resource class that handles sematext_app_mongodb
func resourceAppMongodb() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("MongoDB"),
		CreateContext: resourceOperationCreateAppMongodb,
		ReadContext:   resourceOperationReadAppMongodb,
		UpdateContext: resourceOperationUpdateAppMongodb,
		DeleteContext: resourceOperationDeleteAppMongodb,
		Importer:      resourceOperationImportAppMongodb(),
	}
}

// resourceOperationCreateAppMongodb creates the sematext_app_mongodb resource.
func resourceOperationCreateAppMongodb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "MongoDB"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppMongodb reads the sematext_app_mongodb resource from Sematext Cloud.
func resourceOperationReadAppMongodb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "MongoDB"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppMongodb updates Sematext Cloud from the sematext_app_mongodb resource.
func resourceOperationUpdateAppMongodb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "MongoDB"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppMongodb marks a sematext_app_mongodb resource as retired.
func resourceOperationDeleteAppMongodb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "MongoDB"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}

// resourceOperationImportAppMongodb imports a sematext_app_mongodb resource into the state file.
func resourceOperationImportAppMongodb() *schema.ResourceImporter {
	apptype := "MongoDB"
	return sematext.ResourceOperationImportApp(apptype)
}
