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

// resourceAlertRuleInfra is the resource class that handles sematext_alertrule_infra
func resourceAlertRuleInfra() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("Infra"),
		CreateContext: resourceOperationCreateAlertRuleInfra,
		ReadContext:   resourceOperationReadAlertRuleInfra,
		UpdateContext: resourceOperationUpdateAlertRuleInfra,
		DeleteContext: resourceOperationDeleteAlertRuleInfra,
		Importer:      resourceOperationImportAlertRuleInfra(),
	}
}

// resourceOperationCreateAlertRuleInfra creates the sematext_alertrule_infra resource.
func resourceOperationCreateAlertRuleInfra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "Infra")
}

// resourceOperationReadAlertRuleInfra reads the sematext_alertrule_infra resource from Sematext Cloud.
func resourceOperationReadAlertRuleInfra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "Infra")
}

// resourceOperationUpdateAlertRuleInfra updates Sematext Cloud from the sematext_alertrule_infra resource.
func resourceOperationUpdateAlertRuleInfra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "Infra")
}

// resourceOperationDeleteAlertRuleInfra marks a sematext_alertrule_infra resource as retired.
func resourceOperationDeleteAlertRuleInfra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "Infra")
}

// resourceOperationImportAlertRuleInfra imports a sematext_alertrule_infra resource into the state file.
func resourceOperationImportAlertRuleInfra() *schema.ResourceImporter {
	apptype := "Infra"
	return sematext.ResourceOperationImportAlertRule()
}
