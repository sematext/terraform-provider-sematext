package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/provider.go.template
	Then run generate/generate.sh
*/

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

// Provider is the connection to Sematext Cloud.
func Provider() terraform.ResourceProvider {

	provider := &schema.Provider{

		Schema: map[string]*schema.Schema{
			"sematext_region": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SEMATEXT_REGION", "US"),
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					region := val.(string)
					if val == nil || !IsValidSematextRegion(region) {
						errs = append(errs, fmt.Errorf("ERROR  : sematext_region missing or invalid in sematext provider"))
					}
					return
				},
				Description: "The Sematext region, either US or EU.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"sematext_monitor_akka":          resourceSematextMonitorAkka(),
			"sematext_monitor_apache":        resourceSematextMonitorApache(),
			"sematext_monitor_awsebs":        resourceSematextMonitorAwsebs(),
			"sematext_monitor_awsec2":        resourceSematextMonitorAwsec2(),
			"sematext_monitor_awselb":        resourceSematextMonitorAwselb(),
			"sematext_monitor_cassandra":     resourceSematextMonitorCassandra(),
			"sematext_monitor_clickhouse":    resourceSematextMonitorClickhouse(),
			"sematext_monitor_docker":        resourceSematextMonitorDocker(),
			"sematext_monitor_elasticsearch": resourceSematextMonitorElasticsearch(),
			"sematext_monitor_hadoopmrv1":    resourceSematextMonitorHadoopmrv1(),
			"sematext_monitor_hadoopyarn":    resourceSematextMonitorHadoopyarn(),
			"sematext_monitor_haproxy":       resourceSematextMonitorHaproxy(),
			"sematext_monitor_hbase":         resourceSematextMonitorHbase(),
			"sematext_monitor_jvm":           resourceSematextMonitorJvm(),
			"sematext_monitor_kafka":         resourceSematextMonitorKafka(),
			"sematext_monitor_logsene":       resourceSematextMonitorLogsene(),
			"sematext_monitor_mongodb":       resourceSematextMonitorMongodb(),
			"sematext_monitor_mysql":         resourceSematextMonitorMysql(),
			"sematext_monitor_nginx":         resourceSematextMonitorNginx(),
			"sematext_monitor_nginxplus":     resourceSematextMonitorNginxplus(),
			"sematext_monitor_nodejs":        resourceSematextMonitorNodejs(),
			"sematext_monitor_redis":         resourceSematextMonitorRedis(),
			"sematext_monitor_solr":          resourceSematextMonitorSolr(),
			"sematext_monitor_solrcloud":     resourceSematextMonitorSolrcloud(),
			"sematext_monitor_spark":         resourceSematextMonitorSpark(),
			"sematext_monitor_storm":         resourceSematextMonitorStorm(),
			"sematext_monitor_tomcat":        resourceSematextMonitorTomcat(),
			"sematext_monitor_zookeeper":     resourceSematextMonitorZookeeper(),
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {

		var err error
		var baseURL *url.URL

		/*
			Terraform 0.12 introduced this field to the protocol
			We can therefore assume that if it's missing it's 0.10 or 0.11
		*/
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			terraformVersion = "0.11+compatible"
		}

		region := d.Get("sematext_region").(string)
		if !IsValidSematextRegion(region) {
			return nil, errors.New("ERROR : Missing or invalid sematext_region parameter in provider tf")
		}

		token := os.Getenv("SEMATEXT_API_KEY")
		if !IsValidUUID(token) {
			return nil, errors.New("ERROR : Missing or invalid env SEMATEXT_API_KEY")
		}

		cfg := stcloud.NewConfiguration()

		if baseURL, err = url.Parse("https://apps.sematext.com"); err != nil {
			return nil, err
		}

		switch region {
		case "US":
			baseURL, err = url.Parse("https://apps.sematext.com")
		case "EU":
			baseURL, err = url.Parse("https://apps.eu.sematext.com")
		default:
			err = errors.New("sematext_region must be one of EU, US")
		}

		if err != nil {
			return nil, err
		}

		cfg.BasePath = baseURL.String()
		cfg.UserAgent = fmt.Sprintf("HashiCorp/1.0 Terraform/%s", terraformVersion)
		cfg.AddDefaultHeader("Authorization", fmt.Sprintf("apiKey %s", token))

		client := stcloud.NewAPIClient(cfg)

		return client, nil
	}

	return provider

}
