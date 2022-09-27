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

// resourceAlertRuleRabbitmq is the resource class that handles sematext_alertrule_rabbitmq
func resourceAlertRuleRabbitmq() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("rabbitmq"),
		CreateContext: resourceOperationCreateAlertRuleRabbitmq,
		ReadContext:   resourceOperationReadAlertRuleRabbitmq,
		UpdateContext: resourceOperationUpdateAlertRuleRabbitmq,
		DeleteContext: resourceOperationDeleteAlertRuleRabbitmq,
		Importer:      resourceOperationImportAlertRuleRabbitmq(),
	}
}

// resourceOperationCreateAlertRuleRabbitmq creates the sematext_alertrule_rabbitmq resource.
func resourceOperationCreateAlertRuleRabbitmq(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "rabbitmq")
}

// resourceOperationReadAlertRuleRabbitmq reads the sematext_alertrule_rabbitmq resource from Sematext Cloud.
func resourceOperationReadAlertRuleRabbitmq(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "rabbitmq")
}

// resourceOperationUpdateAlertRuleRabbitmq updates Sematext Cloud from the sematext_alertrule_rabbitmq resource.
func resourceOperationUpdateAlertRuleRabbitmq(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "rabbitmq")
}

// resourceOperationDeleteAlertRuleRabbitmq marks a sematext_alertrule_rabbitmq resource as retired.
func resourceOperationDeleteAlertRuleRabbitmq(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "rabbitmq")
}

// resourceOperationImportAlertRuleRabbitmq imports a sematext_alertrule_rabbitmq resource into the state file.
func resourceOperationImportAlertRuleRabbitmq() *schema.ResourceImporter {
	apptype := "rabbitmq"
	return sematext.ResourceOperationImportAlertRule()
}
