package generated

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app.go.template
	Then run generate/generate.sh
*/

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sematext/terraform-provider-sematext/sematext"
)

// resourceAlertRuleMysql is the resource class that handles sematext_alertrule_mysql
func resourceAlertRuleMysql() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("MySQL"),
		CreateContext: resourceOperationCreateAlertRuleMysql,
		ReadContext:   resourceOperationReadAlertRuleMysql,
		UpdateContext: resourceOperationUpdateAlertRuleMysql,
		DeleteContext: resourceOperationDeleteAlertRuleMysql,
		Importer:      resourceOperationImportAlertRuleMysql(),
	}
}

// resourceOperationCreateAlertRuleMysql creates the sematext_alertrule_mysql resource.
func resourceOperationCreateAlertRuleMysql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "MySQL")
}

// resourceOperationReadAlertRuleMysql reads the sematext_alertrule_mysql resource from Sematext Cloud.
func resourceOperationReadAlertRuleMysql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "MySQL")
}

// resourceOperationUpdateAlertRuleMysql updates Sematext Cloud from the sematext_alertrule_mysql resource.
func resourceOperationUpdateAlertRuleMysql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "MySQL")
}

// resourceOperationDeleteAlertRuleMysql marks a sematext_alertrule_mysql resource as retired.
func resourceOperationDeleteAlertRuleMysql(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "MySQL")
}

// resourceOperationImportAlertRuleMysql imports a sematext_alertrule_mysql resource into the state file.
func resourceOperationImportAlertRuleMysql() *schema.ResourceImporter {
	apptype := "MySQL"
	return sematext.ResourceOperationImportAlertRule()
}
