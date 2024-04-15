package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/hashicorp/go-multierror"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
	"github.com/tomnomnom/linkheader"
)

type CloudConnection struct {
	ID                     *string                 `json:"id,omitempty"`
	Name                   *string                 `json:"name"`
	Provider               *string                 `json:"provider"`
	CloudConnectionDetails *CloudConnectionDetails `json:"cloud_connection_details"`
}

type CloudConnectionDetails struct {
	AccountId  *string `json:"account_id"`
	ExternalId *string `json:"external_id"`
	RoleArn    *string `json:"role_arn"`
}

type CloudConnectionsListResponse struct {
	CloudConnections []CloudConnection `json:"list"`
}

func (c CloudConnection) Exists() bool {
	return utils.IsNonEmpty(c.ID)
}

func (c CloudConnection) ToResourceMap() map[string]any {
	m := make(map[string]any)

	if c.ID != nil {
		m[attributes.ID] = *c.ID
	}
	if c.Name != nil {
		m[attributes.Name] = *c.Name
	}
	if c.Provider != nil {
		m[attributes.CloudConnectionProvider] = c.Provider
	}
	if c.CloudConnectionDetails != nil {
		flattenedDetails := make([]any, 1)
		flattenedDetail := make(map[string]any)
		if c.CloudConnectionDetails.AccountId != nil {
			flattenedDetail[attributes.CloudConnectionAccountId] = *c.CloudConnectionDetails.AccountId
		}
		if c.CloudConnectionDetails.ExternalId != nil {
			flattenedDetail[attributes.CloudConnectionExternalId] = *c.CloudConnectionDetails.ExternalId
		}
		if c.CloudConnectionDetails.RoleArn != nil {
			flattenedDetail[attributes.CloudConnectionRoleARN] = *c.CloudConnectionDetails.RoleArn
		}
		flattenedDetails[0] = flattenedDetail
		m[attributes.CloudConnectionDetails] = flattenedDetails
	}

	return m
}

func isCloudConnectionDataValid(cloudConnection CloudConnection) (bool, error) {
	var errs *multierror.Error

	if cloudConnection.Name == nil || cloudConnection.CloudConnectionDetails == nil || cloudConnection.CloudConnectionDetails.AccountId == nil || cloudConnection.CloudConnectionDetails.ExternalId == nil || cloudConnection.Provider == nil || cloudConnection.CloudConnectionDetails.RoleArn == nil {
		multierror.Append(errs, fmt.Errorf("cloud connection data are not valid"))
		return false, errs
	}

	const AWS_PROVIDER_NAME = "aws"
	namePattern := regexp.MustCompile(`^[A-Za-z0-9-_.]+$`)
	accountIdPattern := regexp.MustCompile(`^\d{12}$`)
	externalIdPattern := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)

	nameValidation := len(*cloudConnection.Name) > 1 && namePattern.MatchString(*cloudConnection.Name)
	if !nameValidation {
		multierror.Append(errs, fmt.Errorf("name is not valid"))
	}
	accountIdValidation := accountIdPattern.MatchString(*cloudConnection.CloudConnectionDetails.AccountId)
	if !accountIdValidation {
		multierror.Append(errs, fmt.Errorf("account id is not valid"))
	}

	externalIdValidation := len(*cloudConnection.CloudConnectionDetails.ExternalId) != 0 && externalIdPattern.MatchString(*cloudConnection.CloudConnectionDetails.ExternalId)
	if !externalIdValidation {
		multierror.Append(errs, fmt.Errorf("external id is not valid"))
	}

	providerValidation := *cloudConnection.Provider == AWS_PROVIDER_NAME
	if !providerValidation {
		multierror.Append(errs, fmt.Errorf("provider is not valid"))
	}

	roleArnValidation := len(*cloudConnection.CloudConnectionDetails.RoleArn) != 0
	if !roleArnValidation {
		multierror.Append(errs, fmt.Errorf("role arn is not valid"))
	}

	return nameValidation && accountIdValidation && providerValidation && externalIdValidation && roleArnValidation, errs
}

func (c OktaPAMClient) ListCloudConnections(ctx context.Context) ([]CloudConnection, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/cloud_connections", url.PathEscape(c.Team))
	cloudConnections := make([]CloudConnection, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)

		resp, err := c.CreateBaseRequest(ctx).SetResult(&CloudConnectionsListResponse{}).Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, http.StatusOK); err != nil {
			return nil, err
		}

		cloudConnectionsListResponse := resp.Result().(*CloudConnectionsListResponse)
		cloudConnections = append(cloudConnections, cloudConnectionsListResponse.CloudConnections...)

		linkHeader := resp.Header().Get("Link")
		if linkHeader == "" {
			break
		}
		links := linkheader.Parse(linkHeader)
		requestURL = ""

		for _, link := range links {
			if link.Rel == "next" {
				requestURL = link.URL
				break
			}
		}
	}

	return cloudConnections, nil
}

func (c OktaPAMClient) GetCloudConnection(ctx context.Context, cloudConnectionId string) (*CloudConnection, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/cloud_connections/%s", url.PathEscape(c.Team), url.PathEscape(cloudConnectionId))
	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&CloudConnection{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == http.StatusOK {
		cloudConnection := resp.Result().(*CloudConnection)
		if cloudConnection.Exists() {
			return cloudConnection, nil
		}
		return nil, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) CreateCloudConnection(ctx context.Context, cloudConnection CloudConnection) (*CloudConnection, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/cloud_connections", url.PathEscape(c.Team))
	logging.Tracef("making POST request to %s", requestURL)

	if valid, errs := isCloudConnectionDataValid(cloudConnection); !valid {
		fmt.Println("Error validating cloud connection data", errs)
		return nil, errs
	}

	resp, err := c.CreateBaseRequest(ctx).SetBody(cloudConnection).SetResult(&CloudConnection{}).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err = checkStatusCode(resp, http.StatusCreated); err != nil {
		return nil, err
	}
	return resp.Result().(*CloudConnection), nil
}

func (c OktaPAMClient) UpdateCloudConnection(ctx context.Context, cloudConnectionID string, updates map[string]any) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/cloud_connections/%s", url.PathEscape(c.Team), url.PathEscape(cloudConnectionID))
	logging.Tracef("making PUT request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).SetBody(updates).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	if _, err = checkStatusCode(resp, http.StatusNoContent); err != nil {
		return err
	}
	return nil
}

func (c OktaPAMClient) DeleteCloudConnection(ctx context.Context, cloudConnectionID string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/cloud_connections/%s", url.PathEscape(c.Team), url.PathEscape(cloudConnectionID))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	if _, err = checkStatusCode(resp, http.StatusOK, http.StatusNoContent, http.StatusNotFound); err != nil {
		return err
	}

	return nil
}
