package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/vitalick/dynadotTerraformProvider/dynadot"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dynadot.Provider,
	})
}
