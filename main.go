package main

import (
	"github.com/EvilSuperstars/terraform-provider-jsonyaml/jsonyaml"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: jsonyaml.Provider,
	})
}
