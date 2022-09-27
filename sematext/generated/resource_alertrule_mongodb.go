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

// resourceAlertRuleMongodb is the resource class that handles sematext_alertrule_mongodb
func resourceAlertRuleMongodb() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("MongoDB"),
		CreateContext: resourceOperationCreateAlertRuleMongodb,
		ReadContext:   resourceOperationReadAlertRuleMongodb,
		UpdateContext: resourceOperationUpdateAlertRuleMongodb,
		DeleteContext: resourceOperationDeleteAlertRuleMongodb,
		Importer:      resourceOperationImportAlertRuleMongodb(),
	}
}

// resourceOperationCreateAlertRuleMongodb creates the sematext_alertrule_mongodb resource.
func resourceOperationCreateAlertRuleMongodb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "MongoDB")
}

// resourceOperationReadAlertRuleMongodb reads the sematext_alertrule_mongodb resource from Sematext Cloud.
func resourceOperationReadAlertRuleMongodb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "MongoDB")
}

// resourceOperationUpdateAlertRuleMongodb updates Sematext Cloud from the sematext_alertrule_mongodb resource.
func resourceOperationUpdateAlertRuleMongodb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "MongoDB")
}

// resourceOperationDeleteAlertRuleMongodb marks a sematext_alertrule_mongodb resource as retired.
func resourceOperationDeleteAlertRuleMongodb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "MongoDB")
}

// resourceOperationImportAlertRuleMongodb imports a sematext_alertrule_mongodb resource into the state file.
func resourceOperationImportAlertRuleMongodb() *schema.ResourceImporter {
	apptype := "MongoDB"
	return sematext.ResourceOperationImportAlertRule()
}
