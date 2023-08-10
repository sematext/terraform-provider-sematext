package provider

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/provider.go.template
	Then run generate/generate.sh
*/

import (
	"context"
	"errors"
	"fmt"
	"os"
	"net/url"

	"github.com/sematext/sematext-api-client-go/stcloud"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure SematextCloudProvider satisfies various provider interfaces.
var _ provider.Provider = &SematextCloudProvider{}

// SematextCloudProvider defines the provider implementation.
type SematextCloudProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string

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
			"sematext_region": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SEMATEXT_REGION", "US"),
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					region := val.(string)
					if val == nil || !sematext.IsValidSematextRegion(region) {
						errs = append(errs, fmt.Errorf("ERROR  : sematext_region missing or invalid in sematext provider"))
					}
					return
				},
				Description: "The Sematext region, either US or EU.",
			},
		},
	}
}

func (p *SematextCloudProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	
	var model SematextCloudProviderModel
	var err error	
	var baseURL *url.URL


    // Read configuration data into model
	resp.Diagnostics.Append(req.Config.Get(ctx, &model)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get the Sematext API Token
	apiToken := os.Getenv("SEMATEXT_API_KEY")
	if !sematext.IsValidUUID(apiToken) {
		resp.Diagnostics.AddError( // @TODO - severity (diag.Error)
			"Missing or invalid env SEMATEXT_API_KEY",
			"Missing or invalid env SEMATEXT_API_KEY" //@TODO adjust brevity
		)
	}	

	// Get the Sematext API Region and set the endpoint
	region := model.SematextRegion
	if !sematext.IsValidSematextRegion(region) {
		resp.Diagnostics.AddError( // @TODO - severity (diag.Error)
			"Missing or invalid sematext_region parameter in provider stanza",
			"Missing or invalid sematext_region parameter in provider stanza",
		)
	}
	model.SematextRegion = region

	switch region {
	case "US":
		baseURL, err = url.Parse("https://apps.sematext.com")
	case "EU":
		baseURL, err = url.Parse("https://apps.eu.sematext.com")
	default:
		resp.Diagnostics.AddError( // @TODO - severity (diag.Error)
			"sematext_region must be one of EU, US",
			"sematext_region must be one of EU, US",
		)		
	}
	if err != null {
		resp.Diagnostics.AddError( // @TODO - severity (diag.Error)
			"Internal error parsing endpoint",
			"Internal error parsing endpoint",
		)
	}
	model.Endpoint = baseURL.String()	

	sematextCloudConnection := stcloud.NewConfiguration()
	sematextCloudConnection.BasePath = baseURL.String()
	sematextCloudConnection.UserAgent = fmt.Sprintf("HashiCorp/1.0 Terraform/%s", tfVersionString)
	sematextCloudConnection.AddDefaultHeader("Authorization", fmt.Sprintf("apiKey %s", apitoken))
	client := stcloud.NewAPIClient(sematextCloudConnection)

	resp.DataSourceData = client
	resp.ResourceData = client

}

func (p *SematextCloudProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		Akka,
Apache,
Awsebs,
Awsec2,
Awselb,
Cassandra,
Clickhouse,
Elasticsearch,
Hadoopmrv1,
Hadoopyarn,
Haproxy,
Hbase,
Infra,
Jvm,
Kafka,
Logsene,
Mongodb,
Mysql,
Nginx,
Nginxplus,
Nodejs,
Redis,
Solr,
Solrcloud,
Spark,
Storm,
Tomcat,
Zookeeper,
Postgresql,
Rabbitmq,
Mobilelogs,

	}
}

/* unused
func (p *SematextCloudProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NOT IMPLEMENTED
	}
}
*/

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &SematextCloudProvider{
			version: version,
		}
	}
}