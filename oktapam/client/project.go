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

type Project struct {
	Name                   *string `json:"name"`
	ID                     *string `json:"id,omitempty"`
	Team                   *string `json:"team,omitempty"`
	DeletedAt              *string `json:"deleted_at,omitempty"`
	NextUnixGID            *int    `json:"next_unix_gid,omitempty"`
	NextUnixUID            *int    `json:"next_unix_uid,omitempty"`
	CreateServerUsers      *bool   `json:"create_server_users,omitempty"`
	ForwardTraffic         *bool   `json:"forward_traffic,omitempty"`
	RDPSessionRecording    *bool   `json:"rdp_session_recording,omitempty"`
	RequirePreAuthForCreds *bool   `json:"require_preauth_for_creds,omitempty"`
	SSHSessionRecording    *bool   `json:"ssh_session_recording,omitempty"`
	GatewaySelector        *string `json:"gateway_selector,omitempty"`
	SSHCertificateType     *string `json:"ssh_certificate_type,omitempty"`
	UserOnDemandPeriod     *int    `json:"user_on_demand_period,omitempty"`
}

func (p Project) ToResourceMap() map[string]any {
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
	if p.RequirePreAuthForCreds != nil {
		m[attributes.RequirePreauthForCreds] = *p.RequirePreAuthForCreds
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

	return m
}

func (p Project) Exists() bool {
	return utils.IsNonEmpty(p.ID) && utils.IsBlank(p.DeletedAt)
}

type ListProjectsParameters struct {
	Self     bool
	Contains string
}

func (p ListProjectsParameters) toQueryParametersMap() map[string]string {
	m := make(map[string]string, 3)

	if p.Self {
		m[attributes.Self] = strconv.FormatBool(p.Self)
	}
	if p.Contains != "" {
		m[attributes.Contains] = p.Contains
	}

	return m
}

type ProjectsListResponse struct {
	Projects []Project `json:"list"`
}

func (c OktaPAMClient) ListProjects(ctx context.Context, parameters ListProjectsParameters) ([]Project, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/projects", url.PathEscape(c.Team))
	projects := make([]Project, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)

		resp, err := c.CreateBaseRequest(ctx).SetQueryParams(parameters.toQueryParametersMap()).SetResult(&ProjectsListResponse{}).Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, http.StatusOK); err != nil {
			return nil, err
		}

		projectsListResponse := resp.Result().(*ProjectsListResponse)
		projects = append(projects, projectsListResponse.Projects...)

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

func (c OktaPAMClient) GetProject(ctx context.Context, name string, allowDeleted bool) (*Project, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s", url.PathEscape(c.Team), url.PathEscape(name))
	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&Project{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == http.StatusOK {
		project := resp.Result().(*Project)
		if !project.Exists() && !allowDeleted {
			return nil, nil
		}
		return project, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) CreateProject(ctx context.Context, proj Project) error {
	// Create the project on the api server specified by project
	requestURL := fmt.Sprintf("/v1/teams/%s/projects", url.PathEscape(c.Team))
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(proj).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, 201)
	return err
}

func (c OktaPAMClient) UpdateProject(ctx context.Context, projectName string, updates map[string]any) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s", url.PathEscape(c.Team), url.PathEscape(projectName))
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(updates).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, http.StatusNoContent)
	return err
}

func (c OktaPAMClient) DeleteProject(ctx context.Context, projectName string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s", url.PathEscape(c.Team), url.PathEscape(projectName))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusNoContent, http.StatusNotFound)
	return err
}
