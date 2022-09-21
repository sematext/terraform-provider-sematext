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

// resourceOperationAppHadoopmrv1 is the resource class that handles sematext_app_hadoopmrv1
func resourceOperationAppHadoopmrv1() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceOperationCreateAppHadoopmrv1,
		ReadContext:   resourceOperationReadAppHadoopmrv1,
		UpdateContext: resourceOperationUpdateAppHadoopmrv1,
		DeleteContext: resourceOperationDeleteAppHadoopmrv1,
		Schema:        sematext.ResourceSchemaApp("Hadoop-MRv1"),
		Importer:      sematext.ResourceImporterApp("Hadoop-MRv1"),
	}
}

// resourceOperationCreateAppHadoopmrv1 creates the sematext_app_hadoopmrv1 resource.
func resourceOperationCreateAppHadoopmrv1(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Hadoop-MRv1"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppHadoopmrv1 reads the sematext_app_hadoopmrv1 resource from Sematext Cloud.
func resourceOperationReadAppHadoopmrv1(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Hadoop-MRv1"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppHadoopmrv1 updates Sematext Cloud from the sematext_app_hadoopmrv1 resource.
func resourceOperationUpdateAppHadoopmrv1(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Hadoop-MRv1"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppHadoopmrv1 marks a sematext_app_hadoopmrv1 resource as retired.
func resourceOperationDeleteAppHadoopmrv1(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Hadoop-MRv1"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}
