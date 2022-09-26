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

// resourceAppAwselb is the resource class that handles sematext_app_awselb
func resourceAppAwselb() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("AWS ELB"),
		CreateContext: resourceOperationCreateAppAwselb,
		ReadContext:   resourceOperationReadAppAwselb,
		UpdateContext: resourceOperationUpdateAppAwselb,
		DeleteContext: resourceOperationDeleteAppAwselb,
		Importer:      resourceOperationImportAppAwselb(),
	}
}

// resourceOperationCreateAppAwselb creates the sematext_app_awselb resource.
func resourceOperationCreateAppAwselb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS ELB"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppAwselb reads the sematext_app_awselb resource from Sematext Cloud.
func resourceOperationReadAppAwselb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS ELB"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppAwselb updates Sematext Cloud from the sematext_app_awselb resource.
func resourceOperationUpdateAppAwselb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS ELB"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppAwselb marks a sematext_app_awselb resource as retired.
func resourceOperationDeleteAppAwselb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS ELB"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}

// resourceOperationImportAppAwselb imports a sematext_app_awselb resource into the state file.
func resourceOperationImportAppAwselb() *schema.ResourceImporter {
	apptype := "AWS ELB"
	return sematext.ResourceOperationImportApp(apptype)
}
