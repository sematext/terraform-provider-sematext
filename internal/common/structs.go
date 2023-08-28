package common

import "github.com/hashicorp/terraform-plugin-framework/types"

type AppTokenType struct {
	Names []string `tdfsk:"name"`
}

type AppTokenEntryType struct {
	CreatedAt string `tdfsk:"createdAt"`
	Enabled   bool   `tdfsk:"enabled"`
	ID        int64  `tdfsk:"id"`
	Name      string `tdfsk:"name"`
	Readable  string `tdfsk:"readable"`
	Token     string `tdfsk:"token"`
	Writeable bool   `tdfsk:"writable"`
}

type AppResourceModelDefault struct {
	Id                string `tfsdk:"id"`
	Name              string `tfsdk:"name"`
	BillingPlanId     int64  `tfsdk:"billing_plan_id"`
	DiscountCode      string `tfsdk:"discount_code"`
	AppToken          []AppTokenType
	ApptokenEntry     []AppTokenEntryType
	ScAppTokenEntries map[string]string `tfsdk:"sc_apptoken_entries"`
}

type AppResourceModelAWS struct {
	Id                string `tfsdk:"id"`
	Name              string `tfsdk:"name"`
	BillingPlanId     int64  `tfsdk:"billing_plan_id"`
	DiscountCode      string `tfsdk:"discount_code"`
	AppToken          []AppTokenType
	ApptokenEntry     []AppTokenEntryType
	ScAppTokenEntries map[string]string `tfsdk:"sc_apptoken_entries"`
	AwsAccessKey      types.String      `tfsdk:"aws_access_key"`
	AwsSecretKey      types.String      `tfsdk:"aws_secret_key"`
	AwsFetchFrequency types.String      `tfsdk:"aws_fetch_frequency"`
	AwsRegion         types.String      `tfsdk:"aws_region"`
}

/*
type AppResourceInterface interface {
	Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse)
	Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse)
	Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse)
	ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse)
	Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse)
	Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse)
	Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse)
	Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse)
}
*/
