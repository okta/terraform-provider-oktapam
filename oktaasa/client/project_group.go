package client

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/logging"
	"github.com/tomnomnom/linkheader"
)

type ProjectGroup struct {
	Project          string `json:"_"`
	Group            string `json:"group"`
	GroupID          string `json:"group_id,omitempty"`
	DeletedAt        string `json:"deleted_at,omitempty"`
	RemovedAt        string `json:"removed_at,omitempty"`
	CreateServerGoup bool   `json:"create_server_group"`
	ServerAccess     bool   `json:"server_access"`
	ServerAdmin      bool   `json:"server_admin"`
	ServersSelector  string `json:"servers_selector,omitempty"`
}

func ProjectGroupFromMap(m map[string]interface{}) (ProjectGroup, error) {
	p := ProjectGroup{}
	if m == nil {
		return p, nil
	}

	for k, v := range m {
		switch k {
		case "project_name":
			p.Project = v.(string)
		case "group_name":
			p.Group = v.(string)
		case "group_id":
			p.GroupID = v.(string)
		case "deleted_at":
			p.DeletedAt = v.(string)
		case "removed_at":
			p.RemovedAt = v.(string)
		case "create_server_group":
			b, err := parseBool(v)
			if err != nil {
				return ProjectGroup{}, err
			}
			p.CreateServerGoup = b
		case "server_access":
			b, err := parseBool(v)
			if err != nil {
				return ProjectGroup{}, err
			}
			p.ServerAccess = b
		case "server_admin":
			b, err := parseBool(v)
			if err != nil {
				return ProjectGroup{}, err
			}
			p.ServerAdmin = b
		case "servers_selector":
			p.ServersSelector = FormatServersSelectorString(v.(map[string]interface{}))
		default:
			return ProjectGroup{}, fmt.Errorf("uknown key: %s", k)
		}
	}

	return p, nil
}

func FormatServersSelectorString(m map[string]interface{}) string {
	selectors := make([]string, 0, len(m))
	for key, value := range m {
		selectors = append(selectors, key+"="+fmt.Sprint(value))
	}
	return strings.Join(selectors, ",")
}

func parseServersSelectorString(s string) (map[string]interface{}, error) {
	m := make(map[string]interface{})

	if len(s) == 0 {
		return m, nil
	}

	selectors := strings.Split(s, ",")
	for _, selector := range selectors {
		parts := strings.Split(selector, "=")

		if len(parts) != 2 {
			return map[string]interface{}{}, fmt.Errorf("selectors must be in the form <key>=<value>: %s", selector)
		}

		m[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	return m, nil
}

func (p *ProjectGroup) ToMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})

	if p.Project != "" {
		m["project_name"] = p.Project
	}
	if p.Group != "" {
		m["group_name"] = p.Group
	}
	if p.GroupID != "" {
		m["group_id"] = p.GroupID
	}
	if p.DeletedAt != "" {
		m["deleted_at"] = p.DeletedAt
	}
	if p.RemovedAt != "" {
		m["removed_at"] = p.RemovedAt
	}
	if p.CreateServerGoup {
		m["create_server_group"] = p.CreateServerGoup
	}
	m["server_access"] = p.ServerAccess
	m["server_admin"] = p.ServerAdmin
	if p.ServersSelector != "" {
		selectorsMap, err := parseServersSelectorString(p.ServersSelector)
		if err != nil {
			return map[string]interface{}{}, err
		}
		m["servers_selector"] = selectorsMap
	}

	return m, nil
}

type ListProjectGroupsParameters struct {
	IncludeRemoved    bool
	CreateServerGroup bool
	HasSelectors      bool
	HasNoSelectors    bool
	OfflineEnabled    bool
}

func (p ListProjectGroupsParameters) toMap() map[string]string {
	m := make(map[string]string, 4)

	if p.IncludeRemoved {
		m["include_removed"] = strconv.FormatBool(p.IncludeRemoved)
	}
	if p.CreateServerGroup {
		m["create_server_group"] = strconv.FormatBool(p.CreateServerGroup)
	}
	if p.HasSelectors {
		m["HasSelectors"] = strconv.FormatBool(p.HasSelectors)
	}
	if p.HasNoSelectors {
		m["HasNoSelectors"] = strconv.FormatBool(p.HasNoSelectors)
	}
	if p.OfflineEnabled {
		m["offline_enabled"] = strconv.FormatBool(p.OfflineEnabled)
	}

	return m
}

type ProjectGroupsListResponse struct {
	ProjectGroups []ProjectGroup `json:"list"`
}

func (c OktaASAClient) ListProjectGroups(ctx context.Context, project string, parameters ListProjectGroupsParameters) ([]ProjectGroup, error) {
	if project == "" {
		return []ProjectGroup{}, fmt.Errorf("supplied blank project name")
	}

	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/groups", url.PathEscape(c.Team), url.PathEscape(project))
	projectGroups := make([]ProjectGroup, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)
		resp, err := c.CreateBaseRequest(ctx).
			SetQueryParams(parameters.toMap()).
			SetResult(&ProjectGroupsListResponse{}).
			Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return []ProjectGroup{}, err
		}
		if _, err := checkStatusCode(resp, 200); err != nil {
			return []ProjectGroup{}, err
		}

		projectGroupsListResponse := resp.Result().(*ProjectGroupsListResponse)
		projectGroups = append(projectGroups, projectGroupsListResponse.ProjectGroups...)

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

	for idx, projectGroup := range projectGroups {
		projectGroup.Project = project
		projectGroups[idx] = projectGroup
	}

	return projectGroups, nil
}

func (c OktaASAClient) GetProjectGroup(ctx context.Context, project, group string) (*ProjectGroup, error) {
	if project == "" {
		return nil, fmt.Errorf("supplied blank project name")
	}
	if group == "" {
		return nil, fmt.Errorf("supplied blank group name")
	}

	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/groups/%s", url.PathEscape(c.Team), url.PathEscape(project), url.PathEscape(group))
	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&ProjectGroup{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == 200 {
		projectGroup := resp.Result().(*ProjectGroup)
		projectGroup.Project = project
		return projectGroup, nil
	} else if statusCode == 404 {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, 200, 404)
}

func (c OktaASAClient) CreateProjectGroup(ctx context.Context, projectGroup ProjectGroup) error {
	// Create the project group on the api server specified by project group
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/groups", url.PathEscape(c.Team), url.PathEscape(projectGroup.Project))
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(projectGroup).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, 204)
	return err
}

func (c OktaASAClient) UpdateProjectGroup(ctx context.Context, projectGroup ProjectGroup) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/groups/%s", url.PathEscape(c.Team), url.PathEscape(projectGroup.Project), url.PathEscape(projectGroup.Group))
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(projectGroup).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, 204)
	return err
}

func (c OktaASAClient) DeleteProjectGroup(ctx context.Context, project, group string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/groups/%s", url.PathEscape(c.Team), url.PathEscape(project), url.PathEscape(group))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, 204, 404)
	return err
}
