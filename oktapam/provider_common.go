package oktapam

import (
	"fmt"

	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func getSDKClientFromMetadata(meta interface{}) client.SDKClientWrapper {
	if apiClients, ok := meta.(*client.APIClients); ok {
		return apiClients.SDKClient
	}
	panic(fmt.Sprintf("expected *client.APIClients, got: %T", meta))
}

// Deprecated: Use getSDKClientFromMetadata instead of using local client
func getLocalClientFromMetadata(meta interface{}) *client.OktaPAMClient {
	if apiClients, ok := meta.(*client.APIClients); ok {
		return apiClients.LocalClient
	}
	panic(fmt.Sprintf("expected *client.APIClients, got: %T", meta))
}

func NewAPIClients(config *client.OktaPAMProviderConfig) (*client.APIClients, error) {
	sdkClient, err := client.CreateSDKClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to load sdk api client: %w", err)
	}
	sdkClientWrapper := client.SDKClientWrapper{
		SDKClient: sdkClient,
		Team:      config.Team,
	}

	localClient, err := client.CreateLocalPAMClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to load local api client: %w", err)
	}

	return &client.APIClients{
		SDKClient:   sdkClientWrapper,
		LocalClient: localClient,
	}, nil
}
