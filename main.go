package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/plugin"
        "github.com/terraform-providers/terraform-provider-sematext/sematext"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: sematext.Provider})
}
