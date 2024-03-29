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

// resourceAppNginx is the resource class that handles sematext_app_nginx
func resourceAppNginx() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("Nginx"),
		CreateContext: resourceOperationCreateAppNginx,
		ReadContext:   resourceOperationReadAppNginx,
		UpdateContext: resourceOperationUpdateAppNginx,
		DeleteContext: resourceOperationDeleteAppNginx,
		Importer:      resourceOperationImportAppNginx(),
	}
}

// resourceOperationCreateAppNginx creates the sematext_app_nginx resource.
func resourceOperationCreateAppNginx(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Nginx"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppNginx reads the sematext_app_nginx resource from Sematext Cloud.
func resourceOperationReadAppNginx(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Nginx"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppNginx updates Sematext Cloud from the sematext_app_nginx resource.
func resourceOperationUpdateAppNginx(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Nginx"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppNginx marks a sematext_app_nginx resource as retired.
func resourceOperationDeleteAppNginx(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Nginx"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}

// resourceOperationImportAppNginx imports a sematext_app_nginx resource into the state file.
func resourceOperationImportAppNginx() *schema.ResourceImporter {
	apptype := "Nginx"
	return sematext.ResourceOperationImportApp(apptype)
}
