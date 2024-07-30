package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/tomnomnom/linkheader"
)

const secretFolderSeparator = "/"

// SecretFolder represents a secret folder entity.  We use this instead of the PAM SDK response since
// there are multiple API response types which are used for secret folders and this provides a common
// structure that can be used
type SecretFolder struct {
	Name            *string `json:"name"`
	ID              *string `json:"id,omitempty"`
	Description     *string `json:"description,omitempty"`
	ParentFolderID  *string `json:"parent_folder_id,omitempty"`
	ResourceGroupID *string `json:"-"`
	ProjectID       *string `json:"-"`
}

func (s SecretFolder) ToResourceMap() map[string]any {
	m := make(map[string]any, 8)

	if s.Name != nil {
		m[attributes.Name] = *s.Name
	}
	if s.ID != nil {
		m[attributes.ID] = *s.ID
	}
	if s.Description != nil {
		m[attributes.Description] = *s.Description
	}
	if s.ParentFolderID != nil {
		m[attributes.ParentFolder] = *s.ParentFolderID
	}

	return m
}

func ListTopLevelSecretFolders(ctx context.Context, sdkClient SDKClientWrapper) ([]SecretFolder, error) {
	var list []SecretFolder
	request := sdkClient.SDKClient.SecretsAPI.ListTopLevelSecretFoldersForUser(ctx, sdkClient.Team)

	for {
		resp, httpResp, callErr := request.Execute()
		if httpResp != nil {
			if _, err := checkStatusCodeFromSDK(httpResp, 200); err != nil {
				return nil, err
			}
		} else if callErr != nil {
			return nil, callErr
		}

		for _, s := range resp.List {
			if s.Type == pam.SecretType_FOLDER {
				list = append(list, ConvertSecretOrFolderListResponseToSecretFolder(s))
			}
		}

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

	return list, nil
}

func executeResolveSecretOrFolderRequest(ctx context.Context, sdkClient SDKClientWrapper, resolveSecretOrFolderRequest *pam.ResolveSecretOrFolderRequest) (*pam.ResolveSecretOrFolderResponse, bool, error) {
	resp, httpResp, callErr := sdkClient.SDKClient.SecretsAPI.ResolveSecretOrFolder(ctx, sdkClient.Team).ResolveSecretOrFolderRequest(*resolveSecretOrFolderRequest).Execute()
	if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
		return nil, false, nil
	} else if callErr != nil {
		return nil, false, callErr
	} else if httpResp != nil {
		if _, err := checkStatusCodeFromSDK(httpResp, 200); err != nil {
			return nil, false, err
		}
	}

	return resp, true, nil
}

func resolveSecretFolder(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID string, projectID string, path string) (*pam.ResolveSecretOrFolderResponse, bool, error) {
	parentFolder, name := getParentFolderPathAndNameFromPath(path)
	resolveSecretOrFolderRequest := pam.NewResolveSecretOrFolderRequest(
		pam.SecretResolveParent{Id: &resourceGroupID, Type: pam.NamedObjectType_RESOURCE_GROUP.Ptr()},
		pam.SecretResolveParent{Id: &projectID, Type: pam.NamedObjectType_PROJECT.Ptr()},
	)

	resolveSecretOrFolderRequest.SetParentFolderPath(parentFolder)
	resolveSecretOrFolderRequest.SetSecretFolderName(name)

	return executeResolveSecretOrFolderRequest(ctx, sdkClient, resolveSecretOrFolderRequest)
}

func ResolveSecretFolder(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID string, projectID string, path string) (*SecretFolder, error) {
	resp, found, err := resolveSecretFolder(ctx, sdkClient, resourceGroupID, projectID, path)
	if err != nil {
		return nil, err
	} else if !found {
		return nil, fmt.Errorf("could not resolve secret folder with resource group: %s, project: %s, path %s", resourceGroupID, projectID, path)
	}

	if secretFolder, err := ConvertResolveSecretOrFolderResponseToSecretFolder(resp); err != nil {
		return nil, err
	} else {
		return &secretFolder, nil
	}
}

func ListTopLevelSecretFoldersForProject(ctx context.Context, sdkCLient SDKClientWrapper, resourceGroupID string, projectID string) ([]SecretFolder, error) {
	var list []SecretFolder
	request := sdkCLient.SDKClient.SecretsAPI.ListTopLevelSecretFoldersForProject(ctx, resourceGroupID, projectID, sdkCLient.Team)
	resp, httpResp, callErr := request.Execute()
	if httpResp != nil {
		if _, err := checkStatusCodeFromSDK(httpResp, 200); err != nil {
			return nil, err
		}
	} else if callErr != nil {
		return nil, callErr
	}

	for _, s := range resp.List {
		if s.Type == pam.SecretType_FOLDER {
			list = append(list, ConvertSecretOrFolderListResponseToSecretFolder(s))
		}
	}

	return list, nil
}

func ListSecretFoldersUnderPath(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID string, projectID string, path string) ([]SecretFolder, error) {
	parentFolder, err := ResolveSecretFolder(ctx, sdkClient, resourceGroupID, projectID, path)
	if err != nil {
		return nil, err
	}

	var list []SecretFolder
	request := sdkClient.SDKClient.SecretsAPI.ListSecretFolderItems(ctx, sdkClient.Team, resourceGroupID, projectID, *parentFolder.ID)

	for {
		resp, httpResp, callErr := request.Execute()
		if httpResp != nil {
			if _, err := checkStatusCodeFromSDK(httpResp, 200); err != nil {
				return nil, err
			}
		} else if callErr != nil {
			return nil, callErr
		}

		for _, s := range resp.List {
			if s.Type == pam.SecretType_FOLDER {
				list = append(list, ConvertSecretOrFolderListResponseToSecretFolder(s))
			}
		}

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

	return list, nil
}

func ConvertSecretOrFolderListResponseToSecretFolder(resp pam.SecretOrFolderListResponse) SecretFolder {
	return SecretFolder{
		Name:            &resp.Name,
		ID:              &resp.Id,
		Description:     resp.Description.Get(),
		ResourceGroupID: resp.ResourceGroup.Id,
		ProjectID:       resp.Project.Id,
	}
}

func ConvertResolveSecretOrFolderResponseToSecretFolder(resp *pam.ResolveSecretOrFolderResponse) (SecretFolder, error) {
	if resp.Type != nil && *resp.Type != pam.SecretType_FOLDER {
		return SecretFolder{}, fmt.Errorf("resolved entity was not a secret folder.  id: %s", resp.Id)
	}

	var parentSecretFolderID *string

	if len(resp.Path) != 0 {
		parentSecretFolderID = &resp.Path[len(resp.Path)-1].Id
	}

	return SecretFolder{
		Name:            &resp.Name,
		ID:              &resp.Id,
		Description:     resp.Description.Get(),
		ResourceGroupID: resp.ResourceGroup.Id,
		ProjectID:       resp.Project.Id,
		ParentFolderID:  parentSecretFolderID,
	}, nil
}

func (c OktaPAMClient) GetSecretFolder(ctx context.Context, resourceGroupID string, projectID string, id string) (*SecretFolder, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s/secret_folders/%s", c.Team, resourceGroupID, projectID, id)

	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&SecretFolder{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == http.StatusOK {
		secretFolder := resp.Result().(*SecretFolder)
		secretFolder.ResourceGroupID = &resourceGroupID
		secretFolder.ProjectID = &projectID
		return secretFolder, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) CreateSecretFolder(ctx context.Context, secretFolder SecretFolder) (*SecretFolder, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s/secret_folders", url.PathEscape(c.Team), url.PathEscape(*secretFolder.ResourceGroupID), url.PathEscape(*secretFolder.ProjectID))
	logging.Tracef("making POST request to %s", requestURL)
	resultingSecretFolder := &SecretFolder{}
	resp, err := c.CreateBaseRequest(ctx).SetBody(secretFolder).SetResult(resultingSecretFolder).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err := checkStatusCode(resp, 201); err == nil {
		return resultingSecretFolder, nil
	} else {
		return nil, err
	}
}

func (c OktaPAMClient) UpdateSecretFolder(ctx context.Context, secretFolder SecretFolder) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s/secret_folders/%s", url.PathEscape(c.Team), url.PathEscape(*secretFolder.ResourceGroupID), url.PathEscape(*secretFolder.ProjectID), url.PathEscape(*secretFolder.ID))
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(secretFolder).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, http.StatusNoContent, http.StatusOK)
	return err
}

func (c OktaPAMClient) DeleteSecretFolder(ctx context.Context, resourceGroupID string, projectID string, id string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s/secret_folders/%s", url.PathEscape(c.Team), url.PathEscape(resourceGroupID), url.PathEscape(projectID), url.PathEscape(id))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusNoContent, http.StatusNotFound)
	return err
}

func getParentFolderPathAndNameFromPath(path string) (string, string) {
	path = strings.TrimRight(path, secretFolderSeparator)
	pathParts := strings.Split(path, secretFolderSeparator)
	lastElement := pathParts[len(pathParts)-1]
	var parentFolder string
	if lastSeparatorIndex := strings.LastIndex(path, secretFolderSeparator); lastSeparatorIndex != -1 {
		parentFolder = path[:lastSeparatorIndex+1]
	}

	return parentFolder, lastElement
}
