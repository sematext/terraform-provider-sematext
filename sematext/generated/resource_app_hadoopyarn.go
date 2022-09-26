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

// resourceAppHadoopyarn is the resource class that handles sematext_app_hadoopyarn
func resourceAppHadoopyarn() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("Hadoop-YARN"),
		CreateContext: resourceOperationCreateAppHadoopyarn,
		ReadContext:   resourceOperationReadAppHadoopyarn,
		UpdateContext: resourceOperationUpdateAppHadoopyarn,
		DeleteContext: resourceOperationDeleteAppHadoopyarn,
		Importer:      resourceOperationImportAppHadoopyarn(),
	}
}

// resourceOperationCreateAppHadoopyarn creates the sematext_app_hadoopyarn resource.
func resourceOperationCreateAppHadoopyarn(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Hadoop-YARN"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppHadoopyarn reads the sematext_app_hadoopyarn resource from Sematext Cloud.
func resourceOperationReadAppHadoopyarn(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Hadoop-YARN"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppHadoopyarn updates Sematext Cloud from the sematext_app_hadoopyarn resource.
func resourceOperationUpdateAppHadoopyarn(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Hadoop-YARN"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppHadoopyarn marks a sematext_app_hadoopyarn resource as retired.
func resourceOperationDeleteAppHadoopyarn(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Hadoop-YARN"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}

// resourceOperationImportAppHadoopyarn imports a sematext_app_hadoopyarn resource into the state file.
func resourceOperationImportAppHadoopyarn() *schema.ResourceImporter {
	apptype := "Hadoop-YARN"
	return sematext.ResourceOperationImportApp(apptype)
}
