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

// resourceAlertRuleCassandra is the resource class that handles sematext_alertrule_cassandra
func resourceAlertRuleCassandra() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("Cassandra"),
		CreateContext: resourceOperationCreateAlertRuleCassandra,
		ReadContext:   resourceOperationReadAlertRuleCassandra,
		UpdateContext: resourceOperationUpdateAlertRuleCassandra,
		DeleteContext: resourceOperationDeleteAlertRuleCassandra,
		Importer:      resourceOperationImportAlertRuleCassandra(),
	}
}

// resourceOperationCreateAlertRuleCassandra creates the sematext_alertrule_cassandra resource.
func resourceOperationCreateAlertRuleCassandra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "Cassandra")
}

// resourceOperationReadAlertRuleCassandra reads the sematext_alertrule_cassandra resource from Sematext Cloud.
func resourceOperationReadAlertRuleCassandra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "Cassandra")
}

// resourceOperationUpdateAlertRuleCassandra updates Sematext Cloud from the sematext_alertrule_cassandra resource.
func resourceOperationUpdateAlertRuleCassandra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "Cassandra")
}

// resourceOperationDeleteAlertRuleCassandra marks a sematext_alertrule_cassandra resource as retired.
func resourceOperationDeleteAlertRuleCassandra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "Cassandra")
}

// resourceOperationImportAlertRuleCassandra imports a sematext_alertrule_cassandra resource into the state file.
func resourceOperationImportAlertRuleCassandra() *schema.ResourceImporter {
	apptype := "Cassandra"
	return sematext.ResourceOperationImportAlertRule()
}
