package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
	"github.com/tomnomnom/linkheader"
)

type ResourceGroupProject struct {
	Name                *string `json:"name"`
	ID                  *string `json:"id,omitempty"`
	Team                *string `json:"team,omitempty"`
	ResourceGroupID     *string `json:"resource_group_id,omitempty"`
	DeletedAt           *string `json:"deleted_at,omitempty"`
	NextUnixGID         *int    `json:"next_unix_gid,omitempty"`
	NextUnixUID         *int    `json:"next_unix_uid,omitempty"`
	CreateServerUsers   *bool   `json:"create_server_users,omitempty"`
	ForwardTraffic      *bool   `json:"forward_traffic,omitempty"`
	RDPSessionRecording *bool   `json:"rdp_session_recording,omitempty"`
	SSHSessionRecording *bool   `json:"ssh_session_recording,omitempty"`
	GatewaySelector     *string `json:"gateway_selector,omitempty"`
	SSHCertificateType  *string `json:"ssh_certificate_type,omitempty"`
	UserOnDemandPeriod  *int    `json:"user_on_demand_period,omitempty"`
	AccountDiscovery    *bool   `json:"server_account_management,omitempty"`
}

func (p ResourceGroupProject) ToResourceMap() map[string]any {
	m := make(map[string]any, 2)

	if p.Name != nil {
		m[attributes.Name] = *p.Name
	}
	if p.ID != nil {
		m[attributes.ID] = *p.ID
	}
	if p.Team != nil {
		m[attributes.Team] = *p.Team
	}
	if p.DeletedAt != nil {
		m[attributes.DeletedAt] = *p.DeletedAt
	}
	if p.NextUnixGID != nil {
		m[attributes.NextUnixGID] = *p.NextUnixGID
	}
	if p.NextUnixUID != nil {
		m[attributes.NextUnixUID] = *p.NextUnixUID
	}
	if p.GatewaySelector != nil {
		m[attributes.GatewaySelector] = *p.GatewaySelector
	}
	if p.CreateServerUsers != nil {
		m[attributes.CreateServerUsers] = *p.CreateServerUsers
	}
	if p.ForwardTraffic != nil {
		m[attributes.ForwardTraffic] = *p.ForwardTraffic
	}
	if p.RDPSessionRecording != nil {
		m[attributes.RDPSessionRecording] = *p.RDPSessionRecording
	}
	if p.SSHSessionRecording != nil {
		m[attributes.SSHSessionRecording] = *p.SSHSessionRecording
	}
	if p.SSHCertificateType != nil {
		m[attributes.SSHCertificateType] = *p.SSHCertificateType
	}
	if p.UserOnDemandPeriod != nil {
		m[attributes.UserOnDemandPeriod] = *p.UserOnDemandPeriod
	}
	if p.AccountDiscovery != nil {
		m[attributes.AccountDiscovery] = *p.AccountDiscovery
	}

	return m
}

func (p ResourceGroupProject) Exists() bool {
	return utils.IsNonEmpty(p.ID) && utils.IsBlank(p.DeletedAt)
}

type ListResourceGroupProjectsParameters struct {
	ResourceGroupID string
	Self            bool
	Contains        string
}

func (p ListResourceGroupProjectsParameters) toQueryParametersMap() map[string]string {
	m := make(map[string]string, 2)

	if p.Self {
		m[attributes.Self] = strconv.FormatBool(p.Self)
	}
	if p.Contains != "" {
		m[attributes.Contains] = p.Contains
	}

	return m
}

type ResourceGroupProjectsListResponse struct {
	ResourceGroupProjects []ResourceGroupProject `json:"list"`
}

func (c OktaPAMClient) ListResourceGroupProjects(ctx context.Context, parameters ListResourceGroupProjectsParameters) ([]ResourceGroupProject, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects", url.PathEscape(c.Team), url.PathEscape(parameters.ResourceGroupID))
	projects := make([]ResourceGroupProject, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)

		resp, err := c.CreateBaseRequest(ctx).SetQueryParams(parameters.toQueryParametersMap()).SetResult(&ResourceGroupProjectsListResponse{}).Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, http.StatusOK); err != nil {
			return nil, err
		}

		projectsListResponse := resp.Result().(*ResourceGroupProjectsListResponse)
		projects = append(projects, projectsListResponse.ResourceGroupProjects...)

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

	return projects, nil
}

func (c OktaPAMClient) GetResourceGroupProject(ctx context.Context, resourceGroupID string, id string, allowDeleted bool) (*ResourceGroupProject, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s", url.PathEscape(c.Team), url.PathEscape(resourceGroupID), url.PathEscape(id))
	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&ResourceGroupProject{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == http.StatusOK {
		project := resp.Result().(*ResourceGroupProject)
		if !project.Exists() && !allowDeleted {
			return nil, nil
		}
		return project, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) CreateResourceGroupProject(ctx context.Context, proj ResourceGroupProject) (*ResourceGroupProject, error) {
	// Create the project on the api server specified by project
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects", url.PathEscape(c.Team), url.PathEscape(*proj.ResourceGroupID))
	logging.Tracef("making POST request to %s", requestURL)
	resultingProject := &ResourceGroupProject{}
	resp, err := c.CreateBaseRequest(ctx).SetBody(proj).SetResult(resultingProject).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err := checkStatusCode(resp, 201); err == nil {
		return resultingProject, nil
	} else {
		return nil, err
	}
}

func (c OktaPAMClient) UpdateResourceGroupProject(ctx context.Context, resourceGroupID string, id string, updates map[string]any) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s", url.PathEscape(c.Team), url.PathEscape(resourceGroupID), url.PathEscape(id))
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(updates).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, http.StatusNoContent)
	return err
}

func (c OktaPAMClient) DeleteResourceGroupProject(ctx context.Context, resourceGroupID string, id string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s", url.PathEscape(c.Team), url.PathEscape(resourceGroupID), url.PathEscape(id))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusNoContent, http.StatusNotFound)
	return err
}
