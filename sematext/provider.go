package sematext

import (
	//"github.com/hashicorp/terraform-plugin-sdk/helper/logging" // TODO DEcide use of logging lib
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const (
	defaultProviderMaxOpenConnections = 4
	defaultExpectedAPIVersion         = "" //TODO Set expected API version
)

func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SEMATEXT_REGION", ""),
				Description: "The Sematext region, either US or EU.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"sematext_application_data": dataSourceSematextApplication(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"sematext_application_resource": resourceSematextApplication(),
		},

		ConfigureFunc: providerConfigure,
	}

	return p
}

type ProviderConfig struct {
	Config Config
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	region := d.Get("sematext_region").(string)
	config, err := Config.Factory(region)
	if err != "" {
		panic(err)
	}

	providerConfig := ProviderConfig{
		Config: config,
	}
	return providerConfig, nil

}
