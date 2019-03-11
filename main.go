package main

import (
	"github.com/hashicorp/terraform-provider-vra7/utils"
	"github.com/hashicorp/terraform-provider-vra7/vrealize"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
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
