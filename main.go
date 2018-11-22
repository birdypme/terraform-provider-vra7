package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/birdypme/terraform-provider-vra7/utils"
	"github.com/birdypme/terraform-provider-vra7/vrealize"
)

func main() {
	utils.InitLog()
	opts := plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return vrealize.Provider()
		},
	}
	plugin.Serve(&opts)
}
