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

// resourceOperationAppStorm is the resource class that handles sematext_app_storm
func resourceOperationAppStorm() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceOperationCreateAppStorm,
		ReadContext:   resourceOperationReadAppStorm,
		UpdateContext: resourceOperationUpdateAppStorm,
		DeleteContext: resourceOperationDeleteAppStorm,
		Schema:        sematext.ResourceSchemaApp("Storm"),
		Importer:      sematext.ResourceImporterApp("Storm"),
	}
}

// resourceOperationCreateAppStorm creates the sematext_app_storm resource.
func resourceOperationCreateAppStorm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Storm"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppStorm reads the sematext_app_storm resource from Sematext Cloud.
func resourceOperationReadAppStorm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Storm"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppStorm updates Sematext Cloud from the sematext_app_storm resource.
func resourceOperationUpdateAppStorm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Storm"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppStorm marks a sematext_app_storm resource as retired.
func resourceOperationDeleteAppStorm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Storm"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}
