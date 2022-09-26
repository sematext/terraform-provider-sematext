package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/sematext/terraform-provider-sematext/sematext/generated"
)

func main() {
	plugin.Serve(
		&plugin.ServeOpts{
			ProviderFunc: generated.Provider,
		},
	)
}
