package tests

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam"
	"github.com/okta/terraform-provider-oktapam/oktapam/fwprovider"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

const defaultRandSeqLength = 20

var randChars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

type compositeDualProviderStruct struct {
	sdkV2Provider *schema.Provider
	fwProvider *fwprovider.OktapamFrameworkProvider
}

func testAccFrameworkMuxProviders(ctx context.Context, t *testing.T) (context.Context, *compositeDualProviderStruct, map[string]func() (tfprotov6.ProviderServer, error)) {
	// Init sdkV2 provider
	sdkV2Provider := oktapam.Provider()
	// Init framework provider
	frameworkProvider := &fwprovider.OktapamFrameworkProvider{}

	// Init mux servers
	muxServer := testAccFrameworkMuxProvidersServer(ctx, sdkV2Provider, frameworkProvider)

	providers := &compositeDualProviderStruct{
		sdkV2Provider: sdkV2Provider,
		fwProvider: frameworkProvider,
	}

	return ctx, providers, muxServer
}

func testAccFrameworkMuxProvidersServer(ctx context.Context, sdkV2Provider *schema.Provider, fwProvider *fwprovider.OktapamFrameworkProvider) map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"oktapam": func() (tfprotov6.ProviderServer, error) {
			// upgrade sdk v2 provider to protocol version 6
			upgradedSdkProvider, err := tf5to6server.UpgradeServer(
				ctx,
				sdkV2Provider.GRPCProvider,
			)
			if err != nil {
				return nil, err
			}

			// combine both providers
			providers := []func() tfprotov6.ProviderServer{
				func() tfprotov6.ProviderServer {
					return upgradedSdkProvider
				},
				providerserver.NewProtocol6(fwProvider),
			}

			if muxServer, err := tf6muxserver.NewMuxServer(ctx, providers...); err != nil {
				log.Fatal(err)
				return nil, err
			}else{
				return muxServer, nil
			}
		},
	}
}

func randSeq() string {
	return randSeqWithLength(defaultRandSeqLength)
}

func randSeqWithLength(length uint) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, length) // Character length
	for i := range b {
		b[i] = randChars[r.Intn(len(randChars))]
	}
	return string(b)
}

func checkTeamApplicable(t *testing.T, isPAMTest bool) {
	if isExecutingPAMTest() != isPAMTest {
		t.Skip("skipping due to team/test mismatch")
	}
}

func isExecutingPAMTest() bool {
	pamAccEnv := os.Getenv("TF_ACC_PAM")
	return pamAccEnv != "" && pamAccEnv != "0"
}