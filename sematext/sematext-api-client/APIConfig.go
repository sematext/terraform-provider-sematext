package sematext

// Config TODO Doc Comment
type APIConfig struct {
	TerraformVerion string
	Client          APIClient
}

// Init TODO Doc Comment
func (apiConfig *APIConfig) Init(region string, terraformVersion string) error {

	apiClient := new(APIClient)
	err := apiClient.Init(region, terraformVersion)

	if err != nil {
		return err
	}

	apiConfig.TerraformVerion = terraformVersion
	apiConfig.Client = *apiClient

	return nil

}
