package main

import (
	"context"
	"flag"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
	"github.com/okta/terraform-provider-oktapam/oktapam/fwprovider"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/okta/terraform-provider-oktapam/oktapam"
)

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	ctx := context.Background()

	// SDKV2 used for tf plugin development is designed for maintaining tf plugins that are compatible with Plugin
	// Protocol version 5. Plugins need to communicate with Terraform CLI, protocol version 5 is supported by CLI version
	// 0.12 and later. Protocol version 6 support tf cli version 1.0 or later.

	// To start using new TF Plugin Framework(https://developer.hashicorp.com/terraform/plugin/framework) we have two options -
	// Option 1: Downgrade new plugin framework server to support protocol version 5
	// Option 2: Upgrade old SDKV2 provider server to support protocol version 6
	// If we go with Option 1, then will not be able to use some of the newer features like Nested Attributes:
	// https://developer.hashicorp.com/terraform/plugin/framework/handling-data/attributes#nested-attribute-types

	//Going with option 2, that will require upgrading tf cli version to 1.0+.

	// tf5to6server enables translating a protocol version 5 provider server into a protocol version 6 provider server.
	upgradedSdkProvider, err := tf5to6server.UpgradeServer(
		ctx,
		oktapam.Provider().GRPCProvider,
	)

	// Combine Providers
	// Refer: https://developer.hashicorp.com/terraform/plugin/mux/combining-protocol-version-6-providers
	providers := []func() tfprotov6.ProviderServer{
		func() tfprotov6.ProviderServer {
			return upgradedSdkProvider
		},

		providerserver.NewProtocol6(fwprovider.New()()),
	}

	muxServer, err := tf6muxserver.NewMuxServer(ctx, providers...)
	if err != nil {
		log.Fatal(err)
	}

	var serveOpts []tf6server.ServeOpt

	if debug {
		serveOpts = append(serveOpts, tf6server.WithManagedDebug())
	}

	err = tf6server.Serve(
		"registry.terraform.io/okta.com/pam/oktapam",
		muxServer.ProviderServer,
		serveOpts...,
	)

	if err != nil {
		log.Fatal(err)
	}
}
