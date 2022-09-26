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

// resourceAppInfra is the resource class that handles sematext_app_infra
func resourceAppInfra() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("Infra"),
		CreateContext: resourceOperationCreateAppInfra,
		ReadContext:   resourceOperationReadAppInfra,
		UpdateContext: resourceOperationUpdateAppInfra,
		DeleteContext: resourceOperationDeleteAppInfra,
		Importer:      resourceOperationImportAppInfra(),
	}
}

// resourceOperationCreateAppInfra creates the sematext_app_infra resource.
func resourceOperationCreateAppInfra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Infra"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppInfra reads the sematext_app_infra resource from Sematext Cloud.
func resourceOperationReadAppInfra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Infra"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppInfra updates Sematext Cloud from the sematext_app_infra resource.
func resourceOperationUpdateAppInfra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Infra"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppInfra marks a sematext_app_infra resource as retired.
func resourceOperationDeleteAppInfra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Infra"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}

// resourceOperationImportAppInfra imports a sematext_app_infra resource into the state file.
func resourceOperationImportAppInfra() *schema.ResourceImporter {
	apptype := "Infra"
	return sematext.ResourceOperationImportApp(apptype)
}
