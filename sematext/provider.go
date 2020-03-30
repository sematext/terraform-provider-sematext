package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/provider.go.template
	Then run generate/generate.sh
*/

import (
	"errors"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/sematext/sematext-api-client/api"
)

// Provider  - TODO Doc Comment
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

		token := os.Getenv("SEMATEXT_API_TOKEN")
		if !IsValidUUID(token) {
			return nil, errors.New("ERROR : Missing or invalid env SEMATEXT_API_TOKEN")
		}

		client := new(api.Client)
		err := client.Init(region, terraformVersion)
		if err != nil {
			return nil, err
		}
		client.SetAuthorization(token)

		return client, nil
	}

	return provider

}
