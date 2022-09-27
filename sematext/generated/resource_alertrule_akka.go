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

// resourceAlertRuleAkka is the resource class that handles sematext_alertrule_akka
func resourceAlertRuleAkka() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("Akka"),
		CreateContext: resourceOperationCreateAlertRuleAkka,
		ReadContext:   resourceOperationReadAlertRuleAkka,
		UpdateContext: resourceOperationUpdateAlertRuleAkka,
		DeleteContext: resourceOperationDeleteAlertRuleAkka,
		Importer:      resourceOperationImportAlertRuleAkka(),
	}
}

// resourceOperationCreateAlertRuleAkka creates the sematext_alertrule_akka resource.
func resourceOperationCreateAlertRuleAkka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "Akka")
}

// resourceOperationReadAlertRuleAkka reads the sematext_alertrule_akka resource from Sematext Cloud.
func resourceOperationReadAlertRuleAkka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "Akka")
}

// resourceOperationUpdateAlertRuleAkka updates Sematext Cloud from the sematext_alertrule_akka resource.
func resourceOperationUpdateAlertRuleAkka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "Akka")
}

// resourceOperationDeleteAlertRuleAkka marks a sematext_alertrule_akka resource as retired.
func resourceOperationDeleteAlertRuleAkka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "Akka")
}

// resourceOperationImportAlertRuleAkka imports a sematext_alertrule_akka resource into the state file.
func resourceOperationImportAlertRuleAkka() *schema.ResourceImporter {
	apptype := "Akka"
	return sematext.ResourceOperationImportAlertRule()
}
