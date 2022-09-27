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

// resourceAlertRuleTomcat is the resource class that handles sematext_alertrule_tomcat
func resourceAlertRuleTomcat() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("Tomcat"),
		CreateContext: resourceOperationCreateAlertRuleTomcat,
		ReadContext:   resourceOperationReadAlertRuleTomcat,
		UpdateContext: resourceOperationUpdateAlertRuleTomcat,
		DeleteContext: resourceOperationDeleteAlertRuleTomcat,
		Importer:      resourceOperationImportAlertRuleTomcat(),
	}
}

// resourceOperationCreateAlertRuleTomcat creates the sematext_alertrule_tomcat resource.
func resourceOperationCreateAlertRuleTomcat(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "Tomcat")
}

// resourceOperationReadAlertRuleTomcat reads the sematext_alertrule_tomcat resource from Sematext Cloud.
func resourceOperationReadAlertRuleTomcat(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "Tomcat")
}

// resourceOperationUpdateAlertRuleTomcat updates Sematext Cloud from the sematext_alertrule_tomcat resource.
func resourceOperationUpdateAlertRuleTomcat(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "Tomcat")
}

// resourceOperationDeleteAlertRuleTomcat marks a sematext_alertrule_tomcat resource as retired.
func resourceOperationDeleteAlertRuleTomcat(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "Tomcat")
}

// resourceOperationImportAlertRuleTomcat imports a sematext_alertrule_tomcat resource into the state file.
func resourceOperationImportAlertRuleTomcat() *schema.ResourceImporter {
	apptype := "Tomcat"
	return sematext.ResourceOperationImportAlertRule()
}
