package sematext

// TODO - The API Key needs to be passed as Header parameter with name Authorization and value should be in the format apiKey <Value>. e.g. apiKey e5f18450-205a-48eb-8589-7d49edaea813


type Config struct {
	TerraformVerion string
	Client SematextAPIClient
}


func (c Config) Factory(region string) (*Config, error) { // TODO Move to schema.ProviderData?

	region := d.Get("sematext_region").(string) // TODO validate region
	client, err := SematextAPIClientHTTPS.Factory(region)
	if err != {
		panic(err)
	}

	config := &Config {
		TerraformVersion: "12" //TODO Terraform version acceptance check
		Client: client
	}

	return config, nil
}
