package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/tomnomnom/linkheader"
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

func ListSudoCommandBundles(ctx context.Context, sdkClient SDKClientWrapper) ([]pam.SudoCommandBundle, error) {
	var response []pam.SudoCommandBundle
	request := sdkClient.SDKClient.SudoCommandsAPI.ListSudoCommandBundles(ctx, sdkClient.Team)
	for {
		resp, httpResp, callErr := request.Execute()
		if httpResp != nil {
			if _, err := checkStatusCodeFromSDK(httpResp, http.StatusOK); err != nil {
				return nil, err
			}
		} else if callErr != nil {
			return nil, callErr
		}

		response = append(response, resp.List...)

		linkHeader := httpResp.Header.Get("Link")
		if linkHeader == "" {
			break
		}

		links := linkheader.Parse(linkHeader)

		var nextURLStr string
		for _, link := range links {
			if link.Rel == "next" {
				nextURLStr = link.URL
			}
		}

		if nextURLStr == "" {
			break
		}

		nextURL, err := url.Parse(nextURLStr)
		if err != nil {
			return nil, err
		}
		values := nextURL.Query()
		offset, ok := values["offset"]
		if !ok || len(offset) != 1 {
			return nil, fmt.Errorf("invalid offset")
		}

		request = request.Offset(offset[0])
	}

	return response, nil
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
