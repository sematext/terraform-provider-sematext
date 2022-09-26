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

// resourceAppAwsec2 is the resource class that handles sematext_app_awsec2
func resourceAppAwsec2() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("AWS EC2"),
		CreateContext: resourceOperationCreateAppAwsec2,
		ReadContext:   resourceOperationReadAppAwsec2,
		UpdateContext: resourceOperationUpdateAppAwsec2,
		DeleteContext: resourceOperationDeleteAppAwsec2,
		Importer:      resourceOperationImportAppAwsec2(),
	}
}

// resourceOperationCreateAppAwsec2 creates the sematext_app_awsec2 resource.
func resourceOperationCreateAppAwsec2(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS EC2"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppAwsec2 reads the sematext_app_awsec2 resource from Sematext Cloud.
func resourceOperationReadAppAwsec2(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS EC2"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppAwsec2 updates Sematext Cloud from the sematext_app_awsec2 resource.
func resourceOperationUpdateAppAwsec2(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS EC2"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppAwsec2 marks a sematext_app_awsec2 resource as retired.
func resourceOperationDeleteAppAwsec2(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "AWS EC2"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}

// resourceOperationImportAppAwsec2 imports a sematext_app_awsec2 resource into the state file.
func resourceOperationImportAppAwsec2() *schema.ResourceImporter {
	apptype := "AWS EC2"
	return sematext.ResourceOperationImportApp(apptype)
}
