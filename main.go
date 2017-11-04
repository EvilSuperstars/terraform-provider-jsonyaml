package main

import (
	"github.com/ewbankkit/terraform-provider-yamldecode/decode"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: decode.Provider,
	})
}
