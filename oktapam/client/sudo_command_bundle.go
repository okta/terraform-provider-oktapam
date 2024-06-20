package client

import (
	"context"
	"net/http"

	"github.com/atko-pam/pam-sdk-go/client/pam"
)

func GetSudoCommandBundle(ctx context.Context, sdkClient SDKClientWrapper, sudoCommandBundleId string) (*pam.SudoCommandBundle, error) {
	request := sdkClient.SDKClient.SudoCommandsAPI.GetSudoCommandBundle(ctx, sdkClient.Team, sudoCommandBundleId)
	resp, httpResp, callErr := request.Execute()
	if httpResp != nil {
		if status, err := checkStatusCodeFromSDK(httpResp, http.StatusOK, http.StatusNotFound); err != nil {
			return nil, err
		} else if status == http.StatusNotFound {
			return nil, nil
		}
	} else if callErr != nil {
		return nil, callErr
	}

	return resp, nil
}

func ListSudoCommandBundles(ctx context.Context, sdkClient SDKClientWrapper) (*pam.ListSudoCommandBundleResponse, error) {
	request := sdkClient.SDKClient.SudoCommandsAPI.ListSudoCommandBundles(ctx, sdkClient.Team)
	resp, httpResp, callErr := request.Execute()
	if httpResp != nil {
		if _, err := checkStatusCodeFromSDK(httpResp, http.StatusOK); err != nil {
			return nil, err
		}
	} else if callErr != nil {
		return nil, callErr
	}

	return resp, nil
}

func CreateSudoCommandBundle(ctx context.Context, sdkClient SDKClientWrapper, sudoCommandBundle *pam.SudoCommandBundle) error {
	request := sdkClient.SDKClient.SudoCommandsAPI.CreateSudoCommandBundle(ctx, sdkClient.Team).SudoCommandBundle(*sudoCommandBundle)
	newSudoCommandBundle, httpResp, callErr := request.Execute()
	if httpResp != nil {
		if _, err := checkStatusCodeFromSDK(httpResp, http.StatusOK, http.StatusCreated); err != nil {
			return err
		}
	} else if callErr != nil {
		return callErr
	}
	sudoCommandBundle.Id = newSudoCommandBundle.Id

	return nil
}

func UpdateSudoCommandBundle(ctx context.Context, sdkClient SDKClientWrapper, sudoCommandBundle *pam.SudoCommandBundle) error {
	request := sdkClient.SDKClient.SudoCommandsAPI.UpdateSudoCommandBundle(ctx, sdkClient.Team, *sudoCommandBundle.Id).SudoCommandBundle(*sudoCommandBundle)
	httpResp, callErr := request.Execute()
	if httpResp != nil {
		if _, err := checkStatusCodeFromSDK(httpResp, http.StatusOK, http.StatusNoContent); err != nil {
			return err
		}
	} else if callErr != nil {
		return callErr
	}

	return nil
}

func DeleteSudoCommandBundle(ctx context.Context, sdkClient SDKClientWrapper, sudoCommandBundleId string) error {
	request := sdkClient.SDKClient.SudoCommandsAPI.DeleteSudoCommandBundle(ctx, sdkClient.Team, sudoCommandBundleId)
	httpResp, callErr := request.Execute()
	if httpResp != nil {
		if _, err := checkStatusCodeFromSDK(httpResp, http.StatusOK, http.StatusNoContent); err != nil {
			return err
		}
	} else if callErr != nil {
		return callErr
	}

	return nil
}
