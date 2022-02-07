package client

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/logging"
	"github.com/tomnomnom/linkheader"
)

type Project struct {
	Name                   string `json:"name"`
	ID                     string `json:"id,omitempty"`
	Team                   string `json:"team,omitempty"`
	DeletedAt              string `json:"deleted_at,omitempty"`
	SharedAdminUserName    string `json:"shared_admin_user_name,omitempty"`
	SharedStandardUserName string `json:"shared_standard_user_name,omitempty"`
	NextUnixGID            int    `json:"next_unix_gid,omitempty"`
	NextUnixUID            int    `json:"next_unix_uid,omitempty"`
	ForceSharedSSHUsers    bool   `json:"force_shared_ssh_users,omitempty"`
	CreateServerUsers      bool   `json:"create_server_users,omitempty"`
	ForwardTraffic         bool   `json:"forward_traffic,omitempty"`
	RDPSessionRecording    bool   `json:"rdp_session_recording,omitempty"`
	RequirePreAuthForCreds bool   `json:"require_preauth_for_creds,omitempty"`
	SSHSessionRecording    bool   `json:"ssh_session_recording,omitempty"`
	GatewaySelector        string `json:"gateway_selector,omitempty"`
	ADJoinedUsers          bool   `json:"ad_joined_users,omitempty"`
}

func (p Project) ToMap() map[string]interface{} {
	m := make(map[string]interface{}, 2)

	if p.Name != "" {
		m["project_name"] = p.Name
	}
	if p.ID != "" {
		m["id"] = p.ID
	}
	if p.Team != "" {
		m["team"] = p.Team
	}
	if p.DeletedAt != "" {
		m["deleted_at"] = p.DeletedAt
	}
	if p.NextUnixGID != 0 {
		m["next_unix_gid"] = p.NextUnixGID
	}
	if p.NextUnixUID != 0 {
		m["next_unix_uid"] = p.NextUnixUID
	}
	if p.SharedAdminUserName != "" {
		m["shared_admin_user_name"] = p.SharedAdminUserName
	}
	if p.SharedStandardUserName != "" {
		m["shared_standard_user_name"] = p.SharedStandardUserName
	}
	if p.GatewaySelector != "" {
		m["gateway_selector"] = p.GatewaySelector
	}

	m["force_shared_ssh_users"] = p.ForceSharedSSHUsers
	m["create_server_users"] = p.CreateServerUsers
	m["forward_traffic"] = p.ForwardTraffic
	m["rdp_session_recording"] = p.RDPSessionRecording
	m["require_preauth_for_creds"] = p.RequirePreAuthForCreds
	m["ssh_session_recording"] = p.SSHSessionRecording
	m["ad_joined_users"] = p.ADJoinedUsers

	return m
}

func (p Project) Exists() bool {
	return p.ID != "" && p.DeletedAt == ""
}

type ListProjectsParameters struct {
	Self          bool
	ADJoinedUsers *bool
	Contains      string
}

func (p ListProjectsParameters) toMap() map[string]string {
	m := make(map[string]string, 3)

	if p.Self {
		m["self"] = strconv.FormatBool(p.Self)
	}
	if p.ADJoinedUsers != nil {
		m["ad_joined_users"] = strconv.FormatBool(*p.ADJoinedUsers)
	}
	if p.Contains != "" {
		m["contains"] = p.Contains
	}

	return m
}

type ProjectsListResponse struct {
	Projects []Project `json:"list"`
}

func (c OktaASAClient) ListProjects(ctx context.Context, parameters ListProjectsParameters) ([]Project, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/projects", url.PathEscape(c.Team))
	projects := make([]Project, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)

		resp, err := c.CreateBaseRequest(ctx).SetQueryParams(parameters.toMap()).SetResult(&ProjectsListResponse{}).Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, 200); err != nil {
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

func (c OktaASAClient) GetProject(ctx context.Context, name string, allowDeleted bool) (*Project, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s", url.PathEscape(c.Team), url.PathEscape(name))
	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&Project{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == 200 {
		project := resp.Result().(*Project)
		if !project.Exists() && !allowDeleted {
			return nil, nil
		}
		return project, nil
	} else if statusCode == 404 {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, 200, 404)
}

func (c OktaASAClient) CreateProject(ctx context.Context, proj Project) error {
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

func (c OktaASAClient) UpdateProject(ctx context.Context, projectName string, updates map[string]interface{}) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s", url.PathEscape(c.Team), url.PathEscape(projectName))
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(updates).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, 204)
	return err
}

func (c OktaASAClient) DeleteProject(ctx context.Context, projectName string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s", url.PathEscape(c.Team), url.PathEscape(projectName))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, 204, 404)
	return err
}
