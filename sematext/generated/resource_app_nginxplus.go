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

// resourceOperationAppNginxplus is the resource class that handles sematext_app_nginxplus
func resourceOperationAppNginxplus() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceOperationCreateAppNginxplus,
		ReadContext:   resourceOperationReadAppNginxplus,
		UpdateContext: resourceOperationUpdateAppNginxplus,
		DeleteContext: resourceOperationDeleteAppNginxplus,
		Schema:        sematext.ResourceSchemaApp("Nginx-Plus"),
		Importer:      sematext.ResourceImporterApp("Nginx-Plus"),
	}
}

// resourceOperationCreateAppNginxplus creates the sematext_app_nginxplus resource.
func resourceOperationCreateAppNginxplus(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Nginx-Plus"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppNginxplus reads the sematext_app_nginxplus resource from Sematext Cloud.
func resourceOperationReadAppNginxplus(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Nginx-Plus"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppNginxplus updates Sematext Cloud from the sematext_app_nginxplus resource.
func resourceOperationUpdateAppNginxplus(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Nginx-Plus"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppNginxplus marks a sematext_app_nginxplus resource as retired.
func resourceOperationDeleteAppNginxplus(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Nginx-Plus"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}
