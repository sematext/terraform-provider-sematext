package sematext

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// returns a common client setup
func sharedClientForRegion(region string) (interface{}, error) {
	if os.Getenv("SEMATEXT_TOKEN") == "" {
		return nil, fmt.Errorf("must provide environment variable SEMATEXT_TOKEN")
	}

	conf := &Config{
		MaxRetries: 5,
		Region:     region,
	}

	// configures a default client for the region, using the above env vars
	client, err := conf.Client()
	if err != nil {
		return nil, fmt.Errorf("error getting AWS client")
	}

	return client, nil
}
