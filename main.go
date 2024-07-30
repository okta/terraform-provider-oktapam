package main

import (
	"context"
	"flag"
	"log"

	"github.com/okta/terraform-provider-oktapam/oktapam"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
)

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	ctx := context.Background()

	muxServer, err := oktapam.ProviderServerFactoryV6(ctx, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	var serveOpts []tf6server.ServeOpt

	if debug {
		serveOpts = append(serveOpts, tf6server.WithManagedDebug())
	}

	_ = tf6server.Serve(
		"registry.terraform.io/okta.com/pam/oktapam",
		muxServer,
		serveOpts...,
	)
}
