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

// resourceAppClickhouse is the resource class that handles sematext_app_clickhouse
func resourceAppClickhouse() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("ClickHouse"),
		CreateContext: resourceOperationCreateAppClickhouse,
		ReadContext:   resourceOperationReadAppClickhouse,
		UpdateContext: resourceOperationUpdateAppClickhouse,
		DeleteContext: resourceOperationDeleteAppClickhouse,
		Importer:      resourceOperationImportAppClickhouse(),
	}
}

// resourceOperationCreateAppClickhouse creates the sematext_app_clickhouse resource.
func resourceOperationCreateAppClickhouse(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "ClickHouse"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppClickhouse reads the sematext_app_clickhouse resource from Sematext Cloud.
func resourceOperationReadAppClickhouse(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "ClickHouse"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppClickhouse updates Sematext Cloud from the sematext_app_clickhouse resource.
func resourceOperationUpdateAppClickhouse(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "ClickHouse"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppClickhouse marks a sematext_app_clickhouse resource as retired.
func resourceOperationDeleteAppClickhouse(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "ClickHouse"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}

// resourceOperationImportAppClickhouse imports a sematext_app_clickhouse resource into the state file.
func resourceOperationImportAppClickhouse() *schema.ResourceImporter {
	apptype := "ClickHouse"
	return sematext.ResourceOperationImportApp(apptype)
}
