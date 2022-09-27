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

// resourceAlertRuleKafka is the resource class that handles sematext_alertrule_kafka
func resourceAlertRuleKafka() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("Kafka"),
		CreateContext: resourceOperationCreateAlertRuleKafka,
		ReadContext:   resourceOperationReadAlertRuleKafka,
		UpdateContext: resourceOperationUpdateAlertRuleKafka,
		DeleteContext: resourceOperationDeleteAlertRuleKafka,
		Importer:      resourceOperationImportAlertRuleKafka(),
	}
}

// resourceOperationCreateAlertRuleKafka creates the sematext_alertrule_kafka resource.
func resourceOperationCreateAlertRuleKafka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "Kafka")
}

// resourceOperationReadAlertRuleKafka reads the sematext_alertrule_kafka resource from Sematext Cloud.
func resourceOperationReadAlertRuleKafka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "Kafka")
}

// resourceOperationUpdateAlertRuleKafka updates Sematext Cloud from the sematext_alertrule_kafka resource.
func resourceOperationUpdateAlertRuleKafka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "Kafka")
}

// resourceOperationDeleteAlertRuleKafka marks a sematext_alertrule_kafka resource as retired.
func resourceOperationDeleteAlertRuleKafka(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "Kafka")
}

// resourceOperationImportAlertRuleKafka imports a sematext_alertrule_kafka resource into the state file.
func resourceOperationImportAlertRuleKafka() *schema.ResourceImporter {
	apptype := "Kafka"
	return sematext.ResourceOperationImportAlertRule()
}
