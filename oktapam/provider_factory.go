package oktapam

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
)

func V6ProviderServerFactory(ctx context.Context) (func() tfprotov6.ProviderServer, error) {
	v5Provider := Provider()
	v6Provider := New()()
	// SDKV2 used for tf plugin development is designed for maintaining tf plugins that are compatible with Plugin
	// Protocol version 5. Plugins need to communicate with Terraform CLI, protocol version 5 is supported by CLI version
	// 0.12 and later. Protocol version 6 support tf cli version 1.0 or later.

	// We have decided to use the new TF Plugin Framework (https://developer.hashicorp.com/terraform/plugin/framework).
	// There were two options available:
	// Option 1: Downgrade the new plugin framework server to support protocol version 5.
	// Option 2: Upgrade the old SDKV2 provider server to support protocol version 6.
	// We chose Option 2 because it allows us to use some newer features like Nested Attributes:
	// https://developer.hashicorp.com/terraform/plugin/framework/handling-data/attributes#nested-attribute-types
	// This choice requires upgrading the Terraform CLI version to 1.0 or later.

	// tf5to6server enables translating a protocol version 5 provider server into a protocol version 6 provider server.
	upgradedV5Provider, err := tf5to6server.UpgradeServer(
		ctx,
		v5Provider.GRPCProvider,
	)

	if err != nil {
		return nil, err
	}

	// Combine Providers
	// Refer: https://developer.hashicorp.com/terraform/plugin/mux/combining-protocol-version-6-providers
	providers := []func() tfprotov6.ProviderServer{
		func() tfprotov6.ProviderServer {
			return upgradedV5Provider
		},
		providerserver.NewProtocol6(v6Provider),
	}

	muxServer, err := tf6muxserver.NewMuxServer(ctx, providers...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return muxServer.ProviderServer, nil
}
