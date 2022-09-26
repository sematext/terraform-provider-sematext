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

// resourceAppSpark is the resource class that handles sematext_app_spark
func resourceAppSpark() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("Spark"),
		CreateContext: resourceOperationCreateAppSpark,
		ReadContext:   resourceOperationReadAppSpark,
		UpdateContext: resourceOperationUpdateAppSpark,
		DeleteContext: resourceOperationDeleteAppSpark,
		Importer:      resourceOperationImportAppSpark(),
	}
}

// resourceOperationCreateAppSpark creates the sematext_app_spark resource.
func resourceOperationCreateAppSpark(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Spark"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppSpark reads the sematext_app_spark resource from Sematext Cloud.
func resourceOperationReadAppSpark(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Spark"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppSpark updates Sematext Cloud from the sematext_app_spark resource.
func resourceOperationUpdateAppSpark(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Spark"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppSpark marks a sematext_app_spark resource as retired.
func resourceOperationDeleteAppSpark(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Spark"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}

// resourceOperationImportAppSpark imports a sematext_app_spark resource into the state file.
func resourceOperationImportAppSpark() *schema.ResourceImporter {
	apptype := "Spark"
	return sematext.ResourceOperationImportApp(apptype)
}
