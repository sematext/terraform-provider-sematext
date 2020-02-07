package sematext

// TODO - The API Key needs to be passed as Header parameter with name Authorization and value should be in the format apiKey <Value>. e.g. apiKey e5f18450-205a-48eb-8589-7d49edaea813

type Config struct {
	TerraformVerion string
	Client          SematextAPIClient
}

func (config Config) Factory(region string) (*Config, error) { // TODO Move to schema.ProviderData?

	client, err := SematextAPIClient.Factory(region)
	if err != "" {
		panic(err)
	}

	return &Config{TerraformVerion: "12", Client: client}, nil

}
