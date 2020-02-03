package sematext

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/logging"
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
			}
		},

		DataSourcesMap: map[string]*schema.Resource{
			"sematext_application" dataSourceSematextApplication()
		},

		ResourcesMap: map[string]*schema.Resource{
			"sematext_application_resource" resourceSematextApplication()
		},

		ConfigureFunc: providerConfigure

	}

	return p
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := sematext.NewClient(d.Get("region").(string))
	log.Println("[INFO] Sematext client successfully initialized, now validating...")
	return client, nil
}
