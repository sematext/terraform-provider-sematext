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

	    "github.com/sematext/terraform-provider-sematext/internal/resource/akka"
    "github.com/sematext/terraform-provider-sematext/internal/resource/apache"
    "github.com/sematext/terraform-provider-sematext/internal/resource/awsebs"
    "github.com/sematext/terraform-provider-sematext/internal/resource/awsec2"
    "github.com/sematext/terraform-provider-sematext/internal/resource/awselb"
    "github.com/sematext/terraform-provider-sematext/internal/resource/cassandra"
    "github.com/sematext/terraform-provider-sematext/internal/resource/clickhouse"
    "github.com/sematext/terraform-provider-sematext/internal/resource/elasticsearch"
    "github.com/sematext/terraform-provider-sematext/internal/resource/hadoopmrv1"
    "github.com/sematext/terraform-provider-sematext/internal/resource/hadoopyarn"
    "github.com/sematext/terraform-provider-sematext/internal/resource/haproxy"
    "github.com/sematext/terraform-provider-sematext/internal/resource/hbase"
    "github.com/sematext/terraform-provider-sematext/internal/resource/infra"
    "github.com/sematext/terraform-provider-sematext/internal/resource/jvm"
    "github.com/sematext/terraform-provider-sematext/internal/resource/kafka"
    "github.com/sematext/terraform-provider-sematext/internal/resource/logsene"
    "github.com/sematext/terraform-provider-sematext/internal/resource/mongodb"
    "github.com/sematext/terraform-provider-sematext/internal/resource/mysql"
    "github.com/sematext/terraform-provider-sematext/internal/resource/nginx"
    "github.com/sematext/terraform-provider-sematext/internal/resource/nginxplus"
    "github.com/sematext/terraform-provider-sematext/internal/resource/nodejs"
    "github.com/sematext/terraform-provider-sematext/internal/resource/redis"
    "github.com/sematext/terraform-provider-sematext/internal/resource/solr"
    "github.com/sematext/terraform-provider-sematext/internal/resource/solrcloud"
    "github.com/sematext/terraform-provider-sematext/internal/resource/spark"
    "github.com/sematext/terraform-provider-sematext/internal/resource/storm"
    "github.com/sematext/terraform-provider-sematext/internal/resource/tomcat"
    "github.com/sematext/terraform-provider-sematext/internal/resource/zookeeper"
    "github.com/sematext/terraform-provider-sematext/internal/resource/postgresql"
    "github.com/sematext/terraform-provider-sematext/internal/resource/rabbitmq"
    "github.com/sematext/terraform-provider-sematext/internal/resource/mobilelogs"

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
        akka.NewResource,
        apache.NewResource,
        awsebs.NewResource,
        awsec2.NewResource,
        awselb.NewResource,
        cassandra.NewResource,
        clickhouse.NewResource,
        elasticsearch.NewResource,
        hadoopmrv1.NewResource,
        hadoopyarn.NewResource,
        haproxy.NewResource,
        hbase.NewResource,
        infra.NewResource,
        jvm.NewResource,
        kafka.NewResource,
        logsene.NewResource,
        mongodb.NewResource,
        mysql.NewResource,
        nginx.NewResource,
        nginxplus.NewResource,
        nodejs.NewResource,
        redis.NewResource,
        solr.NewResource,
        solrcloud.NewResource,
        spark.NewResource,
        storm.NewResource,
        tomcat.NewResource,
        zookeeper.NewResource,
        postgresql.NewResource,
        rabbitmq.NewResource,
        mobilelogs.NewResource,

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

