package main

import (
	"github.com/hashicorp/terraform/plugin"
        "github.com/suzuki-yosuke/terraform-provider-conoha"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: conoha.Provider,
	})
}
