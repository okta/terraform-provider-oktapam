package main

import (
	"context"
	"flag"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
	provider "github.com/okta/terraform-provider-oktapam/oktapam/fwprovider"
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

//func main() {
//	var debug bool
//
//	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
//	flag.Parse()
//
//	plugin.Serve(&plugin.ServeOpts{
//		Debug: debug,
//		ProviderFunc: func() *schema.Provider {
//			return oktapam.Provider()
//		},
//	})
//}

//var (
//	// Version can be updated by goreleaser on release
//	version string = "dev"
//)

// TF core (CLI) >= 1.0 doesn't support protocol version 5
// By default, SDKV2 providers support Protocol Version 5 and framework provider support Protocol Version 6
// Upgrading SDKV2 provider to support protocol version 6 and then serving it along with framework provider. We can downgrade
// framework provider to support protocol version 5 too but then we cannot use some of the features like NestedAttribute.
func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	ctx := context.Background()
	upgradedSdkProvider, err := tf5to6server.UpgradeServer(
		ctx,
		oktapam.Provider().GRPCProvider,
	)

	providers := []func() tfprotov6.ProviderServer{
		func() tfprotov6.ProviderServer {
			return upgradedSdkProvider
		},

		// Example terraform-plugin-framework provider
		providerserver.NewProtocol6(provider.New()()),
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
