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

// resourceAlertRuleAwsec2 is the resource class that handles sematext_alertrule_awsec2
func resourceAlertRuleAwsec2() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("AWS EC2"),
		CreateContext: resourceOperationCreateAlertRuleAwsec2,
		ReadContext:   resourceOperationReadAlertRuleAwsec2,
		UpdateContext: resourceOperationUpdateAlertRuleAwsec2,
		DeleteContext: resourceOperationDeleteAlertRuleAwsec2,
		Importer:      resourceOperationImportAlertRuleAwsec2(),
	}
}

// resourceOperationCreateAlertRuleAwsec2 creates the sematext_alertrule_awsec2 resource.
func resourceOperationCreateAlertRuleAwsec2(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "AWS EC2")
}

// resourceOperationReadAlertRuleAwsec2 reads the sematext_alertrule_awsec2 resource from Sematext Cloud.
func resourceOperationReadAlertRuleAwsec2(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "AWS EC2")
}

// resourceOperationUpdateAlertRuleAwsec2 updates Sematext Cloud from the sematext_alertrule_awsec2 resource.
func resourceOperationUpdateAlertRuleAwsec2(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "AWS EC2")
}

// resourceOperationDeleteAlertRuleAwsec2 marks a sematext_alertrule_awsec2 resource as retired.
func resourceOperationDeleteAlertRuleAwsec2(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "AWS EC2")
}

// resourceOperationImportAlertRuleAwsec2 imports a sematext_alertrule_awsec2 resource into the state file.
func resourceOperationImportAlertRuleAwsec2() *schema.ResourceImporter {
	apptype := "AWS EC2"
	return sematext.ResourceOperationImportAlertRule()
}
