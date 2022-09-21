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

// resourceOperationAppMysql is the resource class that handles sematext_app_mysql
func resourceOperationAppMysql() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceOperationCreateAppMysql,
		ReadContext:   resourceOperationReadAppMysql,
		UpdateContext: resourceOperationUpdateAppMysql,
		DeleteContext: resourceOperationDeleteAppMysql,
		Schema:        sematext.ResourceSchemaApp("MySQL"),
		Importer:      sematext.ResourceImporterApp("MySQL"),
	}
}

// resourceOperationCreateAppMysql creates the sematext_app_mysql resource.
func resourceOperationCreateAppMysql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "MySQL"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppMysql reads the sematext_app_mysql resource from Sematext Cloud.
func resourceOperationReadAppMysql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "MySQL"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppMysql updates Sematext Cloud from the sematext_app_mysql resource.
func resourceOperationUpdateAppMysql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "MySQL"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppMysql marks a sematext_app_mysql resource as retired.
func resourceOperationDeleteAppMysql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "MySQL"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}
