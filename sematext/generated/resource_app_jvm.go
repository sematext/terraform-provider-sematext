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

// resourceOperationAppJvm is the resource class that handles sematext_app_jvm
func resourceOperationAppJvm() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceOperationCreateAppJvm,
		ReadContext:   resourceOperationReadAppJvm,
		UpdateContext: resourceOperationUpdateAppJvm,
		DeleteContext: resourceOperationDeleteAppJvm,
		Schema:        sematext.ResourceSchemaApp("JVM"),
		Importer:      sematext.ResourceImporterApp("JVM"),
	}
}

// resourceOperationCreateAppJvm creates the sematext_app_jvm resource.
func resourceOperationCreateAppJvm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "JVM"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppJvm reads the sematext_app_jvm resource from Sematext Cloud.
func resourceOperationReadAppJvm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "JVM"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppJvm updates Sematext Cloud from the sematext_app_jvm resource.
func resourceOperationUpdateAppJvm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "JVM"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppJvm marks a sematext_app_jvm resource as retired.
func resourceOperationDeleteAppJvm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "JVM"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}
