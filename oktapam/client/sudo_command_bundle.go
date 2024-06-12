package client

import (
	"context"
	"net/http"

	"github.com/atko-pam/pam-sdk-go/client/pam"
)

func GetSudoCommandsBundle(ctx context.Context, sdkClient SDKClientWrapper, sudoCommandsBundleId string) (*pam.SudoCommandBundle, error) {
	request := sdkClient.SDKClient.SudoCommandsAPI.GetSudoCommandBundle(ctx, sdkClient.Team, sudoCommandsBundleId)
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

func ListSudoCommandsBundles(ctx context.Context, sdkClient SDKClientWrapper) (*pam.ListSudoCommandBundleResponse, error) {
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

func CreateSudoCommandsBundle(ctx context.Context, sdkClient SDKClientWrapper, sudoCommandsBundle *pam.SudoCommandBundle) error {
	request := sdkClient.SDKClient.SudoCommandsAPI.CreateSudoCommandBundle(ctx, sdkClient.Team).SudoCommandBundle(*sudoCommandsBundle)
	newSudoCommandBundle, httpResp, callErr := request.Execute()
	if httpResp != nil {
		if _, err := checkStatusCodeFromSDK(httpResp, http.StatusOK, http.StatusCreated); err != nil {
			return err
		}
	} else if callErr != nil {
		return callErr
	}
	sudoCommandsBundle.Id = newSudoCommandBundle.Id

	return nil
}

func UpdateSudoCommandsBundle(ctx context.Context, sdkClient SDKClientWrapper, sudoCommandsBundle *pam.SudoCommandBundle) error {
	request := sdkClient.SDKClient.SudoCommandsAPI.UpdateSudoCommandBundle(ctx, sdkClient.Team, *sudoCommandsBundle.Id).SudoCommandBundle(*sudoCommandsBundle)
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

func DeleteSudoCommandsBundle(ctx context.Context, sdkClient SDKClientWrapper, sudoCommandsBundleId string) error {
	request := sdkClient.SDKClient.SudoCommandsAPI.DeleteSudoCommandBundle(ctx, sdkClient.Team, sudoCommandsBundleId)
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
