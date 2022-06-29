package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/utils"

	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/logging"
	"github.com/tomnomnom/linkheader"
)

type UserStatus string

const UserStatusActive UserStatus = "ACTIVE"
const UserStatusDisabled UserStatus = "DISABLED"
const UserStatusDeleted UserStatus = "DELETED"

type UserType string

const UserTypeHuman UserType = "human"
const UserTypeService UserType = "service"

type User struct {
	Name           *string `json:"name"`
	ID             *string `json:"id"`
	TeamName       *string `json:"team_name"`
	ServerUserName *string `json:"server_user_name,omitempty"`
	Status         *string `json:"status"` // TODO: Change this to using typed string
	DeletedAt      *string `json:"deleted_at,omitempty"`
	UserType       *string `json:"user_type"`
}

func UserFromMap(m map[string]interface{}) (*User, error) {
	if m == nil {
		return nil, nil
	}

	su := User{}
	for k, v := range m {
		switch k {
		case attributes.Name:
			su.Name = utils.AsStringPtr(v.(string))
		case attributes.Status:
			su.Status = utils.AsStringPtr(v.(string))
		default:
			return nil, fmt.Errorf("uknown key: %s", k)
		}
	}
	return &su, nil
}

func (su User) ToResourceMap() map[string]interface{} {
	m := make(map[string]interface{}, 2)

	if su.Name != nil {
		m[attributes.Name] = *su.Name
	}
	if su.ID != nil {
		m[attributes.ID] = *su.ID
	}
	if su.DeletedAt != nil {
		m[attributes.DeletedAt] = *su.DeletedAt
	}
	if su.TeamName != nil {
		m[attributes.TeamName] = *su.TeamName
	}
	if su.ServerUserName != nil {
		m[attributes.ServerUserName] = *su.ServerUserName
	}
	if su.Status != nil {
		m[attributes.Status] = *su.Status
	}
	if su.UserType != nil {
		m[attributes.UserType] = *su.UserType
	}
	return m
}

type ListUsersParameters struct {
	Contains            string
	StartsWith          string
	Status              []string
	IncludeServiceUsers string // NOTE: Unused in service_user endpoint
}

func (p ListUsersParameters) toQueryParametersMap() map[string]string {
	m := make(map[string]string, 4)

	if p.Contains != "" {
		m[attributes.Contains] = p.Contains
	}
	if p.StartsWith != "" {
		m[attributes.StartsWith] = p.StartsWith
	}
	if len(p.Status) > 0 {
		m[attributes.Status] = strings.Join(p.Status, ",")
	}
	if p.IncludeServiceUsers != "" {
		m[attributes.IncludeServiceUsers] = p.IncludeServiceUsers
	}
	return m
}

type ServiceUsersListResponse struct {
	ServiceUsers []User `json:"list"`
}

func (c OktaPAMClient) ListServiceUsers(ctx context.Context, parameters ListUsersParameters) ([]User, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/service_users", url.PathEscape(c.Team))
	serviceUsers := make([]User, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)
		resp, err := c.CreateBaseRequest(ctx).
			SetResult(&ServiceUsersListResponse{}).
			SetQueryParams(parameters.toQueryParametersMap()).
			Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, 200); err != nil {
			return nil, err
		}

		serviceUsersListResponse := resp.Result().(*ServiceUsersListResponse)
		serviceUsers = append(serviceUsers, serviceUsersListResponse.ServiceUsers...)

		linkHeader := resp.Header().Get("Link")
		if linkHeader == "" {
			break
		}
		links := linkheader.Parse(linkHeader)
		requestURL = ""

		isNext := false
		for _, link := range links {
			if link.Rel == "next" {
				requestURL = link.URL
				isNext = true
			}
		}
		if !isNext {
			break
		}
	}
	return serviceUsers, nil
}

func (c OktaPAMClient) GetServiceUser(ctx context.Context, serviceUserName string) (*User, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/service_users/%s", url.PathEscape(c.Team), url.PathEscape(serviceUserName))
	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&User{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == 200 {
		serviceUser := resp.Result().(*User)
		return serviceUser, nil
	} else if statusCode == 404 {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, 200, 404)
}

func (c OktaPAMClient) CreateServiceUser(ctx context.Context, userName string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/service_users", url.PathEscape(c.Team))

	request := map[string]interface{}{
		attributes.Name: userName,
	}
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(request).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, 201)
	return err
}

func (c OktaPAMClient) UpdateServiceUser(ctx context.Context, userName string, serviceUser *User) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/service_users/%s", url.PathEscape(c.Team), url.PathEscape(userName))
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(serviceUser).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, 200)
	return err
}

func (c OktaPAMClient) DeleteServiceUser(ctx context.Context, userName string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/service_users/%s", url.PathEscape(c.Team), url.PathEscape(userName))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, 204, 404)
	return err
}