package client

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/okta/terraform-provider-oktapam/oktapam/client/wrappers"
	"github.com/tomnomnom/linkheader"
	"gopkg.in/square/go-jose.v2"
)

func RevealSecret(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID string, projectID string, secretID string) (*wrappers.SecretWrapper, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("could not generate encryption key")
	}
	jwk := jose.JSONWebKey{
		Key:       &privateKey.PublicKey,
		Algorithm: string(jose.RSA_OAEP_256),
	}
	jwkBytes, err := json.Marshal(jwk)
	if err != nil {
		return nil, fmt.Errorf("error marshaling jwk: %w", err)
	}
	rawPublicKey := pam.RawJSONWebKey{}
	if err := json.Unmarshal(jwkBytes, &rawPublicKey); err != nil {
		return nil, fmt.Errorf("error unmarshaling jwk: %w", err)
	}

	request := sdkClient.SDKClient.SecretsAPI.RevealSecret(ctx, sdkClient.Team, resourceGroupID, projectID, secretID).SecretRevealRequest(*pam.NewSecretRevealRequest(rawPublicKey))
	response, httpResponse, err := request.Execute()
	if err != nil {
		return nil, err
	}
	_, err = checkStatusCodeFromSDK(httpResponse, 200)
	if err != nil {
		return nil, err
	}

	encryptedSecret, err := jose.ParseEncrypted(response.SecretJwe)
	if err != nil {
		return nil, fmt.Errorf("could not parse jwe: %w", err)
	}
	privateJWK := jose.JSONWebKey{
		Key:       privateKey,
		Algorithm: string(jose.RSA_OAEP_256),
	}

	decryptedSecret, err := encryptedSecret.Decrypt(privateJWK)
	if err != nil {
		return nil, fmt.Errorf("could not decrypt secret: %w", err)
	}
	var result map[string]string

	if err = json.Unmarshal(decryptedSecret, &result); err != nil {
		return nil, fmt.Errorf("could not unmarshal secret json: %w", err)
	}

	secret := pam.Secret{
		Id:          response.Id,
		Name:        response.Name,
		Description: response.Description,
		Path:        response.Path,
	}

	parentFolderID := ""
	if len(response.Path) > 0 {
		parentFolderID = response.Path[len(response.Path)-1].Id
	}

	secretWrapper := wrappers.SecretWrapper{
		Secret:          &secret,
		ResourceGroupID: resourceGroupID,
		ProjectID:       projectID,
		ParentFolderID:  parentFolderID,
		SecretContents:  result,
	}

	return &secretWrapper, nil
}

func CreateSecret(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID string, projectID string, parentFolderID string, name string, description string, secret map[string]string) (*wrappers.SecretWrapper, error) {
	secretJSON, err := json.Marshal(secret)
	if err != nil {
		// do not utilize error from marshal since we do not want to leak any plaintext
		return nil, fmt.Errorf("could not marshal secret data to json")
	}
	encryptedString, err := sdkClient.SDKClient.Encrypt(string(secretJSON))
	if err != nil {
		return nil, err
	}
	request := *pam.NewSecretCreateOrUpdateRequest(name, *encryptedString, parentFolderID).SetDescription(description)
	response, httpResponse, err := sdkClient.SDKClient.SecretsAPI.CreateSecret(ctx, sdkClient.Team, resourceGroupID, projectID).SecretCreateOrUpdateRequest(request).Execute()
	if err != nil {
		return nil, err
	}
	_, err = checkStatusCodeFromSDK(httpResponse, 200, 201)
	if err != nil {
		return nil, err
	}

	secretWrapper := wrappers.SecretWrapper{
		Secret:          response,
		ResourceGroupID: resourceGroupID,
		ProjectID:       projectID,
		ParentFolderID:  parentFolderID,
		SecretContents:  secret,
	}

	return &secretWrapper, nil
}

func UpdateSecret(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID string, projectID string, parentFolderID string, secretID string, name string, description string, secret map[string]string) (*wrappers.SecretWrapper, error) {
	secretJSON, err := json.Marshal(secret)
	if err != nil {
		// do not utilize error from marshal since we do not want to leak any plaintext
		return nil, fmt.Errorf("could not marshal secret data to json")
	}
	encryptedString, err := sdkClient.SDKClient.Encrypt(string(secretJSON))
	if err != nil {
		return nil, err
	}

	request := *pam.NewSecretCreateOrUpdateRequest(name, *encryptedString, parentFolderID).SetDescription(description)
	response, httpResponse, err := sdkClient.SDKClient.SecretsAPI.UpdateSecret(ctx, sdkClient.Team, resourceGroupID, projectID, secretID).SecretCreateOrUpdateRequest(request).Execute()
	if err != nil {
		return nil, err
	}
	_, err = checkStatusCodeFromSDK(httpResponse, 200, 201)
	if err != nil {
		return nil, err
	}

	secretWrapper := wrappers.SecretWrapper{
		Secret:          response,
		ResourceGroupID: resourceGroupID,
		ProjectID:       projectID,
		ParentFolderID:  parentFolderID,
		SecretContents:  secret,
	}

	return &secretWrapper, nil
}

func DeleteSecret(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID string, projectID string, secretID string) error {
	httpResponse, err := sdkClient.SDKClient.SecretsAPI.DeleteSecret(ctx, sdkClient.Team, resourceGroupID, projectID, secretID).Execute()
	if err != nil {
		return err
	}
	_, err = checkStatusCodeFromSDK(httpResponse, 204)
	if err != nil {
		return err
	}

	return nil
}

func ConvertSecretOrFolderListResponseToSecret(parentFolderPath []pam.SecretPath, resp pam.SecretOrFolderListResponse) *wrappers.SecretWrapper {
	parentFolderID := ""
	if len(parentFolderPath) > 0 {
		parentFolderID = parentFolderPath[len(parentFolderPath)-1].Id
	}
	return &wrappers.SecretWrapper{
		Secret: &pam.Secret{
			Id:          resp.Id,
			Name:        resp.Name,
			Description: resp.Description,
			Path:        parentFolderPath,
		},
		ParentFolderID:  parentFolderID,
		ResourceGroupID: *resp.ResourceGroup.Id,
		ProjectID:       *resp.Project.Id,
		SecretContents:  nil,
	}
}

func ConvertResolveSecretOrFolderResponseToSecretNoContents(resourceGroupID string, projectID string, resp *pam.ResolveSecretOrFolderResponse) *wrappers.SecretWrapper {
	parentFolderID := ""
	if len(resp.Path) > 0 {
		parentFolderID = resp.Path[len(resp.Path)-1].Id
	}

	return &wrappers.SecretWrapper{
		Secret: &pam.Secret{
			Id:          resp.Id,
			Name:        resp.Name,
			Description: resp.Description,
			Path:        resp.Path,
		},
		ParentFolderID:  parentFolderID,
		ResourceGroupID: resourceGroupID,
		ProjectID:       projectID,
		SecretContents:  nil,
	}
}

func resolveSecret(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID string, projectID string, path string) (*pam.ResolveSecretOrFolderResponse, bool, error) {
	parentFolder, name := getParentFolderPathAndNameFromPath(path)

	resolveSecretOrFolderRequest := pam.NewResolveSecretOrFolderRequest(
		pam.SecretResolveParent{Id: &resourceGroupID, Type: pam.NamedObjectType_RESOURCE_GROUP.Ptr()},
		pam.SecretResolveParent{Id: &projectID, Type: pam.NamedObjectType_PROJECT.Ptr()},
	)

	resolveSecretOrFolderRequest.SetParentFolderPath(parentFolder)
	resolveSecretOrFolderRequest.SetSecretName(name)

	return executeResolveSecretOrFolderRequest(ctx, sdkClient, resolveSecretOrFolderRequest)
}

func resolveSecretOrFolder(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID string, projectID string, path string) (*pam.ResolveSecretOrFolderResponse, error) {
	if secretFolderResp, foundSecretFolder, secretFolderErr := resolveSecretFolder(ctx, sdkClient, resourceGroupID, projectID, path); secretFolderErr != nil {
		return nil, secretFolderErr
	} else if !foundSecretFolder {
		// resolve_secret requires that you know whether it is a secret or folder, which we do not know here.  if it doesn't resolve as a secret folder, we try again as a secret
		if secretResp, foundSecret, secretErr := resolveSecret(ctx, sdkClient, resourceGroupID, projectID, path); secretErr != nil {
			return nil, secretErr
		} else if !foundSecret {
			return nil, fmt.Errorf("could not resolve secret or folder with resource group: %s, project: %s, path %s", resourceGroupID, projectID, path)
		} else {
			return secretResp, nil
		}
	} else {
		return secretFolderResp, nil
	}
}

func ListSecrets(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID string, projectID string, path string, revealSecrets bool) ([]*wrappers.SecretWrapper, error) {
	resp, err := resolveSecretOrFolder(ctx, sdkClient, resourceGroupID, projectID, path)
	if err != nil {
		return nil, err
	}

	if *resp.Type == pam.SecretType_FOLDER {
		parentFolderID := resp.Id
		parentFolderPath := make([]pam.SecretPath, len(resp.Path)+1)
		copy(parentFolderPath, resp.Path)
		parentFolderPath[len(parentFolderPath)-1] = *pam.NewSecretPath(resp.Id, resp.Name)

		return ListSecretsUnderParentFolder(ctx, sdkClient, resourceGroupID, projectID, parentFolderID, parentFolderPath, revealSecrets)
	}

	if revealSecrets {
		if secret, err := RevealSecret(ctx, sdkClient, resourceGroupID, projectID, resp.Id); err != nil {
			return nil, err
		} else {
			return []*wrappers.SecretWrapper{secret}, nil
		}
	}
	return []*wrappers.SecretWrapper{ConvertResolveSecretOrFolderResponseToSecretNoContents(resourceGroupID, projectID, resp)}, nil
}

func ListSecretsUnderParentFolder(ctx context.Context, sdkClient SDKClientWrapper, resourceGroupID string, projectID string, parentFolderID string, parentFolderPath []pam.SecretPath, revealSecrets bool) ([]*wrappers.SecretWrapper, error) {
	var list []*wrappers.SecretWrapper
	request := sdkClient.SDKClient.SecretsAPI.ListSecretFolderItems(ctx, sdkClient.Team, resourceGroupID, projectID, parentFolderID)

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
			if s.Type != pam.SecretType_FOLDER {
				if revealSecrets {
					if secret, err := RevealSecret(ctx, sdkClient, resourceGroupID, projectID, s.Id); err != nil {
						return nil, err
					} else {
						list = append(list, secret)
					}
				} else {
					list = append(list, ConvertSecretOrFolderListResponseToSecret(parentFolderPath, s))
				}
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
