package provider

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/provider.go.template
	Then run generate/generate.sh
*/

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/sematext/sematext-api-client-go/stcloud"
	"github.com/sematext/terraform-provider-sematext/internal/resources"
)

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}


// Ensure SematextCloudProvider satisfies various provider interfaces.
var _ provider.Provider = &SematextCloudProvider{}

func New(version string) func() provider.Provider {
    return func() provider.Provider {
        return &SematextCloudProvider{
            version: version,
        }
    }
}

// SematextCloudProvider defines the provider implementation.
type SematextCloudProvider struct {
	version string // set to provider version on release or "dev" on local build or "test" during testing
}

// SematextCloudProviderModel describes the provider data model.
type SematextCloudProviderModel struct {
	SematextRegion types.String `tfsdk:"sematext_region"`	
}

func (p *SematextCloudProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "SematextCloud"
	resp.Version = p.version
}

func (p *SematextCloudProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {

	resp.Schema = schema.Schema{

		Attributes: map[string]schema.Attribute{

			"sematext_region": schema.StringAttribute{
				MarkdownDescription: "The Sematext Cloud region, either US or EU.",
				Description: "The Sematext Cloud region, either US or EU.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOf("US", "EU"),
				},			
			},

		},
	}
}

func (p *SematextCloudProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	
	var model SematextCloudProviderModel
	var baseURL *url.URL

	resp.Diagnostics.Append(req.Config.Get(ctx, &model)...)

	if resp.Diagnostics.HasError() {
		return
	}


	apiToken := os.Getenv("SEMATEXT_API_KEY")
	if !isValidUUID(apiToken) {
		resp.Diagnostics.AddError( // @TODO - severity (diag.Error)
			"Missing or invalid env SEMATEXT_API_KEY",
			"Missing or invalid env SEMATEXT_API_KEY", //@TODO adjust brevity
		)
		return
	}

	if model.SematextRegion.ValueString() == "US" {
		baseURL, _ = url.Parse("https://apps.sematext.com")
	} else if model.SematextRegion.ValueString() == "EU" {
		baseURL, _ = url.Parse("https://apps.eu.sematext.com")
	} else {
		resp.Diagnostics.AddError( // @TODO - add severity (diag.Error)
			"sematext_region must be one of EU, US",
			"sematext_region must be one of EU, US",
		)
		return
	}

	sematextCloudConnection := stcloud.NewConfiguration()
	sematextCloudConnection.BasePath = baseURL.String()
	sematextCloudConnection.UserAgent = fmt.Sprintf("HashiCorp/1.0 Terraform/%s", p.version)
	sematextCloudConnection.AddDefaultHeader("Authorization", fmt.Sprintf("apiKey %s", apiToken))
	client := stcloud.NewAPIClient(sematextCloudConnection)

	resp.DataSourceData = client
	resp.ResourceData = client

}

func (p *SematextCloudProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
        resources.NewAppAkkaResource,
        resources.NewAppApacheResource,
        resources.NewAppAwsebsResource,
        resources.NewAppAwsec2Resource,
        resources.NewAppAwselbResource,
        resources.NewAppCassandraResource,
        resources.NewAppClickhouseResource,
        resources.NewAppElasticsearchResource,
        resources.NewAppHadoopmrv1Resource,
        resources.NewAppHadoopyarnResource,
        resources.NewAppHaproxyResource,
        resources.NewAppHbaseResource,
        resources.NewAppInfraResource,
        resources.NewAppJvmResource,
        resources.NewAppKafkaResource,
        resources.NewAppLogseneResource,
        resources.NewAppMongodbResource,
        resources.NewAppMysqlResource,
        resources.NewAppNginxResource,
        resources.NewAppNginxplusResource,
        resources.NewAppNodejsResource,
        resources.NewAppRedisResource,
        resources.NewAppSolrResource,
        resources.NewAppSolrcloudResource,
        resources.NewAppSparkResource,
        resources.NewAppStormResource,
        resources.NewAppTomcatResource,
        resources.NewAppZookeeperResource,
        resources.NewAppPostgresqlResource,
        resources.NewAppRabbitmqResource,
        resources.NewAppMobilelogsResource,

	}
}

func (p *SematextCloudProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
	}
}


func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

