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

// resourceOperationAppPostgresql is the resource class that handles sematext_app_postgresql
func resourceOperationAppPostgresql() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceOperationCreateAppPostgresql,
		ReadContext:   resourceOperationReadAppPostgresql,
		UpdateContext: resourceOperationUpdateAppPostgresql,
		DeleteContext: resourceOperationDeleteAppPostgresql,
		Schema:        sematext.ResourceSchemaApp("postgresql"),
		Importer:      sematext.ResourceImporterApp("postgresql"),
	}
}

// resourceOperationCreateAppPostgresql creates the sematext_app_postgresql resource.
func resourceOperationCreateAppPostgresql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "postgresql"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppPostgresql reads the sematext_app_postgresql resource from Sematext Cloud.
func resourceOperationReadAppPostgresql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "postgresql"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppPostgresql updates Sematext Cloud from the sematext_app_postgresql resource.
func resourceOperationUpdateAppPostgresql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "postgresql"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppPostgresql marks a sematext_app_postgresql resource as retired.
func resourceOperationDeleteAppPostgresql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "postgresql"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}
