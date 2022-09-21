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

// resourceOperationAppTomcat is the resource class that handles sematext_app_tomcat
func resourceOperationAppTomcat() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceOperationCreateAppTomcat,
		ReadContext:   resourceOperationReadAppTomcat,
		UpdateContext: resourceOperationUpdateAppTomcat,
		DeleteContext: resourceOperationDeleteAppTomcat,
		Schema:        sematext.ResourceSchemaApp("Tomcat"),
		Importer:      sematext.ResourceImporterApp("Tomcat"),
	}
}

// resourceOperationCreateAppTomcat creates the sematext_app_tomcat resource.
func resourceOperationCreateAppTomcat(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Tomcat"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppTomcat reads the sematext_app_tomcat resource from Sematext Cloud.
func resourceOperationReadAppTomcat(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Tomcat"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppTomcat updates Sematext Cloud from the sematext_app_tomcat resource.
func resourceOperationUpdateAppTomcat(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Tomcat"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppTomcat marks a sematext_app_tomcat resource as retired.
func resourceOperationDeleteAppTomcat(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Tomcat"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}
