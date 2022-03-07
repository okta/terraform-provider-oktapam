package client

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/logging"
	"github.com/terraform-providers/terraform-provider-oktaasa/oktaasa/utils"
	"github.com/tomnomnom/linkheader"
)

type Group struct {
	Name      *string  `json:"name"`
	ID        *string  `json:"id,omitempty"`
	DeletedAt *string  `json:"deleted_at,omitempty"`
	Roles     []string `json:"roles"`
}

func (g Group) ToResourceMap() map[string]interface{} {
	m := make(map[string]interface{}, 2)

	if g.Name != nil {
		m["name"] = *g.Name
	}
	if g.ID != nil {
		m["id"] = *g.ID
	}
	if g.DeletedAt != nil {
		m["deleted_at"] = *g.DeletedAt
	}
	m["roles"] = g.Roles

	return m
}

func (g Group) Exists() bool {
	return utils.IsNonEmpty(g.ID) && utils.IsBlank(g.DeletedAt)
}

type ListGroupsParameters struct {
	Contains               string
	IncludeDeleted         bool
	OnlyIncludeDeleted     bool
	DisconnectedModeOnOnly bool
}

func (p ListGroupsParameters) toQueryParametersMap() map[string]string {
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
	requestURL := fmt.Sprintf("/v1/teams/%s/groups", url.PathEscape(c.Team))
	groups := make([]Group, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)
		resp, err := c.CreateBaseRequest(ctx).
			SetResult(&GroupsListResponse{}).
			SetQueryParams(parameters.toQueryParametersMap()).
			Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, 200); err != nil {
			return nil, err
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

func (c OktaASAClient) GetGroup(ctx context.Context, name string, allowDeleted bool) (*Group, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/groups/%s", url.PathEscape(c.Team), url.PathEscape(name))
	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&Group{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == 200 {
		group := resp.Result().(*Group)
		if group.Exists() || allowDeleted {
			return group, nil
		}
		return nil, nil
	} else if statusCode == 404 {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, 200, 404)
}

func (c OktaASAClient) CreateGroup(ctx context.Context, group Group) error {
	// Create the group on the api server specified by group
	requestURL := fmt.Sprintf("/v1/teams/%s/groups", url.PathEscape(c.Team))
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(group).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, 201)
	return err
}

func (c OktaASAClient) UpdateGroup(ctx context.Context, groupName string, updates map[string]interface{}) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/groups/%s", url.PathEscape(c.Team), url.PathEscape(groupName))
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(updates).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, 204)
	return err
}

func (c OktaASAClient) DeleteGroup(ctx context.Context, groupName string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/groups/%s", url.PathEscape(c.Team), url.PathEscape(groupName))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, 204, 404)
	return err
}
