package client

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/logging"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/utils"
	"github.com/tomnomnom/linkheader"
)

type ProjectGroup struct {
	Project          *string `json:"_"`
	Group            *string `json:"group"`
	GroupID          *string `json:"group_id,omitempty"`
	DeletedAt        *string `json:"deleted_at,omitempty"`
	RemovedAt        *string `json:"removed_at,omitempty"`
	CreateServerGoup bool    `json:"create_server_group"`
	ServerAccess     bool    `json:"server_access"`
	ServerAdmin      bool    `json:"server_admin"`
	ServersSelector  string  `json:"servers_selector,omitempty"`
}

func ProjectGroupFromMap(m map[string]interface{}) (*ProjectGroup, error) {
	if m == nil {
		return nil, nil
	}

	p := ProjectGroup{}
	for k, v := range m {
		switch k {
		case attributes.ProjectName:
			p.Project = utils.AsStringPtr(v.(string))
		case attributes.GroupName:
			p.Group = utils.AsStringPtr(v.(string))
		case attributes.GroupID:
			p.GroupID = utils.AsStringPtr(v.(string))
		case attributes.DeletedAt:
			p.DeletedAt = utils.AsStringPtr(v.(string))
		case attributes.RemovedAt:
			p.RemovedAt = utils.AsStringPtr(v.(string))
		case attributes.CreateServerGroup:
			b, err := parseBool(v)
			if err != nil {
				return nil, err
			}
			p.CreateServerGoup = b
		case attributes.ServerAccess:
			b, err := parseBool(v)
			if err != nil {
				return nil, err
			}
			p.ServerAccess = b
		case attributes.ServerAdmin:
			b, err := parseBool(v)
			if err != nil {
				return nil, err
			}
			p.ServerAdmin = b
		case attributes.ServersSelector:
			serversSelectorString, err := FormatServersSelectorString(v.(map[string]interface{}))
			if err != nil {
				return nil, err
			}

			p.ServersSelector = serversSelectorString
		default:
			return nil, fmt.Errorf("uknown key: %s", k)
		}
	}

	return &p, nil
}

func invalidSelectorPart(s string) bool {
	return strings.ContainsAny(s, "=,")
}

func FormatServersSelectorString(m map[string]interface{}) (string, error) {
	selectors := make([]string, 0, len(m))
	for key, value := range m {
		if invalidSelectorPart(key) {
			return "", fmt.Errorf("servers selector key cannot contain a '=' or ',' key: %s", key)
		}
		valueString := fmt.Sprint(value)
		if invalidSelectorPart(valueString) {
			return "", fmt.Errorf("servers selector value cannot contain a '=' or ',' value: %s", valueString)
		}
		selectors = append(selectors, key+"="+valueString)
	}
	return strings.Join(selectors, ","), nil
}

func parseServersSelectorString(s string) (map[string]interface{}, error) {
	if len(s) == 0 {
		return nil, nil
	}

	m := make(map[string]interface{})

	selectors := strings.Split(s, ",")
	for _, selector := range selectors {
		parts := strings.Split(selector, "=")

		if len(parts) != 2 {
			return nil, fmt.Errorf("selectors must be in the form <key>=<value>: %s", selector)
		}

		m[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	return m, nil
}

func (p *ProjectGroup) ToResourceMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})

	if p.Project != nil {
		m[attributes.ProjectName] = *p.Project
	}
	if p.Group != nil {
		m[attributes.GroupName] = *p.Group
	}
	if p.GroupID != nil {
		m[attributes.GroupID] = *p.GroupID
	}
	if p.DeletedAt != nil {
		m[attributes.DeletedAt] = *p.DeletedAt
	}
	if p.RemovedAt != nil {
		m[attributes.RemovedAt] = *p.RemovedAt
	}
	if p.CreateServerGoup {
		m[attributes.CreateServerGroup] = p.CreateServerGoup
	}
	m[attributes.ServerAccess] = p.ServerAccess
	m[attributes.ServerAdmin] = p.ServerAdmin
	if p.ServersSelector != "" {
		selectorsMap, err := parseServersSelectorString(p.ServersSelector)
		if err != nil {
			return nil, err
		}
		m[attributes.ServersSelector] = selectorsMap
	}

	return m, nil
}

type ListProjectGroupsParameters struct {
	IncludeRemoved         bool
	CreateServerGroup      bool
	HasSelectors           bool
	HasNoSelectors         bool
	DisconnectedModeOnOnly bool
}

func (p ListProjectGroupsParameters) toQueryParametersMap() map[string]string {
	m := make(map[string]string, 5)

	if p.IncludeRemoved {
		m[attributes.IncludeRemoved] = strconv.FormatBool(p.IncludeRemoved)
	}
	if p.CreateServerGroup {
		m[attributes.CreateServerGroup] = strconv.FormatBool(p.CreateServerGroup)
	}
	if p.HasSelectors {
		m[attributes.HasSelectors] = strconv.FormatBool(p.HasSelectors)
	}
	if p.HasNoSelectors {
		m[attributes.HasNoSelectors] = strconv.FormatBool(p.HasNoSelectors)
	}
	if p.DisconnectedModeOnOnly {
		m[attributes.DisconnectedModeOnOnly] = strconv.FormatBool(p.DisconnectedModeOnOnly)
	}

	return m
}

type ProjectGroupsListResponse struct {
	ProjectGroups []ProjectGroup `json:"list"`
}

func (c OktaPAMClient) ListProjectGroups(ctx context.Context, project string, parameters ListProjectGroupsParameters) ([]ProjectGroup, error) {
	if project == "" {
		return nil, fmt.Errorf("supplied blank project name")
	}

	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/groups", url.PathEscape(c.Team), url.PathEscape(project))
	projectGroups := make([]ProjectGroup, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)
		resp, err := c.CreateBaseRequest(ctx).
			SetQueryParams(parameters.toQueryParametersMap()).
			SetResult(&ProjectGroupsListResponse{}).
			Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, 200, 404); err != nil {
			return nil, err
		}

		if resp.StatusCode() == 404 {
			logging.Warnf("received a 404 for %s, could indicate that the referenced project does not exist", requestURL)
			break
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
		projectGroup.Project = &project
		projectGroups[idx] = projectGroup
	}

	return projectGroups, nil
}

func (c OktaPAMClient) GetProjectGroup(ctx context.Context, project, group string) (*ProjectGroup, error) {
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
		projectGroup.Project = &project
		return projectGroup, nil
	} else if statusCode == 404 {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, 200, 404)
}

func (c OktaPAMClient) CreateProjectGroup(ctx context.Context, projectGroup ProjectGroup) error {
	// Create the project group on the api server specified by project group
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/groups", url.PathEscape(c.Team), url.PathEscape(*projectGroup.Project))
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(projectGroup).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, 204)
	return err
}

func (c OktaPAMClient) UpdateProjectGroup(ctx context.Context, projectGroup ProjectGroup) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/projects/%s/groups/%s", url.PathEscape(c.Team), url.PathEscape(*projectGroup.Project), url.PathEscape(*projectGroup.Group))
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(projectGroup).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, 204)
	return err
}

func (c OktaPAMClient) DeleteProjectGroup(ctx context.Context, project, group string) error {
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
