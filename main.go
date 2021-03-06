package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/sematext/terraform-provider-sematext/sematext"
)

func main() {
	plugin.Serve(
		&plugin.ServeOpts{
			ProviderFunc: sematext.Provider,
		},
	)
}
