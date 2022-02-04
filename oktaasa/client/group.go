package client

import (
	"context"
	"fmt"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/logging"
	"github.com/tomnomnom/linkheader"
)

type Group struct {
	Name                 string   `json:"name"`
	ID                   string   `json:"id,omitempty"`
	DeletedAt            string   `json:"deleted_at,omitempty"`
	FederatedFromTeam    string   `json:"federated_from_team,omitempty"`
	FederationApprovedAt string   `json:"federation_approved_at,omitempty"`
	Roles                []string `json:"roles"`
}

func (g Group) ToMap() map[string]interface{} {
	m := make(map[string]interface{}, 2)

	if g.Name != "" {
		m["name"] = g.Name
	}
	if g.ID != "" {
		m["id"] = g.ID
	}
	if g.DeletedAt != "" {
		m["deleted_at"] = g.DeletedAt
	}
	if g.FederatedFromTeam != "" {
		m["federated_from_team"] = g.FederatedFromTeam
	}
	if g.FederationApprovedAt != "" {
		m["federation_approved_at"] = g.FederationApprovedAt
	}
	m["roles"] = g.Roles

	return m
}

func (g Group) Exists() bool {
	return g.ID != "" && g.DeletedAt == ""
}

type ListGroupsParameters struct {
	Contains               string
	IncludeDeleted         bool
	OnlyIncludeDeleted     bool
	DisconnectedModeOnOnly bool
}

func (p ListGroupsParameters) toMap() map[string]string {
	m := make(map[string]string, 4)

	if p.Contains != "" {
		m["contains"] = p.Contains
	}
	if p.IncludeDeleted {
		m["IncludeDeleted"] = strconv.FormatBool(p.IncludeDeleted)
	}
	if p.OnlyIncludeDeleted {
		m["OnlyIncludeDeleted"] = strconv.FormatBool(p.OnlyIncludeDeleted)
	}
	if p.DisconnectedModeOnOnly {
		m["DisconnectedModeOnOnly"] = strconv.FormatBool(p.DisconnectedModeOnOnly)
	}

	return m
}

type GroupsListResponse struct {
	Groups []Group `json:"list"`
}

func (c OktaASAClient) ListGroups(ctx context.Context, parameters ListGroupsParameters) ([]Group, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/groups", c.Team)
	groups := make([]Group, 0)

	for {
		logging.Tracef("making GET request to %s", requestURL)
		resp, err := c.CreateBaseRequest(ctx).
			SetResult(&GroupsListResponse{}).
			SetQueryParams(parameters.toMap()).
			Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return []Group{}, err
		}
		err = checkStatusCode(resp, 200)
		if err != nil {
			return []Group{}, err
		}

		groupsListResponse := resp.Result().(*GroupsListResponse)
		groups = append(groups, groupsListResponse.Groups...)

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

	return groups, nil
}

func (c OktaASAClient) GetGroup(ctx context.Context, name string, allowDeleted bool) (Group, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/groups/%s", c.Team, name)
	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&Group{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return Group{}, err
	}
	statusCode := resp.StatusCode()

	if statusCode == 200 {
		group := resp.Result().(*Group)
		if group.Exists() || allowDeleted {
			return *group, nil
		}
		return Group{}, nil
	} else if statusCode == 404 {
		return Group{}, nil
	}

	return Group{}, createErrorForInvalidCode(resp, 200, 404)
}

func (c OktaASAClient) CreateGroup(ctx context.Context, group Group) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/groups", c.Team)
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(group).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	return checkStatusCode(resp, 201)
}

func (c OktaASAClient) UpdateGroup(ctx context.Context, groupName string, updates map[string]interface{}) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/groups/%s", c.Team, groupName)
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(updates).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	return checkStatusCode(resp, 204)
}

func (c OktaASAClient) DeleteGroup(ctx context.Context, groupName string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/groups/%s", c.Team, groupName)
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	return checkStatusCode(resp, 204, 404)
}
