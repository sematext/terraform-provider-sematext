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

// resourceOperationAppNodejs is the resource class that handles sematext_app_nodejs
func resourceOperationAppNodejs() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceOperationCreateAppNodejs,
		ReadContext:   resourceOperationReadAppNodejs,
		UpdateContext: resourceOperationUpdateAppNodejs,
		DeleteContext: resourceOperationDeleteAppNodejs,
		Schema:        sematext.ResourceSchemaApp("Node.js"),
		Importer:      sematext.ResourceImporterApp("Node.js"),
	}
}

// resourceOperationCreateAppNodejs creates the sematext_app_nodejs resource.
func resourceOperationCreateAppNodejs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Node.js"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppNodejs reads the sematext_app_nodejs resource from Sematext Cloud.
func resourceOperationReadAppNodejs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Node.js"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppNodejs updates Sematext Cloud from the sematext_app_nodejs resource.
func resourceOperationUpdateAppNodejs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Node.js"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppNodejs marks a sematext_app_nodejs resource as retired.
func resourceOperationDeleteAppNodejs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Node.js"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}
