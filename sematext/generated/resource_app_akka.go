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

// resourceAppAkka is the resource class that handles sematext_app_akka
func resourceAppAkka() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("Akka"),
		CreateContext: resourceOperationCreateAppAkka,
		ReadContext:   resourceOperationReadAppAkka,
		UpdateContext: resourceOperationUpdateAppAkka,
		DeleteContext: resourceOperationDeleteAppAkka,
		Importer:      resourceOperationImportAppAkka(),
	}
}

// resourceOperationCreateAppAkka creates the sematext_app_akka resource.
func resourceOperationCreateAppAkka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Akka"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppAkka reads the sematext_app_akka resource from Sematext Cloud.
func resourceOperationReadAppAkka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Akka"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppAkka updates Sematext Cloud from the sematext_app_akka resource.
func resourceOperationUpdateAppAkka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Akka"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppAkka marks a sematext_app_akka resource as retired.
func resourceOperationDeleteAppAkka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Akka"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}

// resourceOperationImportAppAkka imports a sematext_app_akka resource into the state file.
func resourceOperationImportAppAkka() *schema.ResourceImporter {
	apptype := "Akka"
	return sematext.ResourceOperationImportApp(apptype)
}
