package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/errors"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"

	"github.com/okta/terraform-provider-oktapam/oktapam/utils"

	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/tomnomnom/linkheader"
)

type User struct {
	Name           *string                   `json:"name"`
	TeamName       *string                   `json:"team_name"`
	ServerUserName *string                   `json:"server_user_name,omitempty"`
	DeletedAt      *string                   `json:"deleted_at,omitempty"`
	Status         *typed_strings.UserStatus `json:"status"`
	UserType       *typed_strings.UserType   `json:"user_type"`
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
		case attributes.UserType:
			userType := typed_strings.UserType(v.(string))
			su.UserType = &userType
		case attributes.Status:
			userStatus := typed_strings.UserStatus(v.(string))
			su.Status = &userStatus
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

func (p ListUsersParameters) toQueryParametersMap() map[string][]string {
	m := make(map[string][]string, 4)

	if p.Contains != "" {
		m[attributes.Contains] = []string{p.Contains}
	}
	if p.StartsWith != "" {
		m[attributes.StartsWith] = []string{p.StartsWith}
	}
	if len(p.Status) > 0 {
		m[attributes.Status] = p.Status
	}
	if p.IncludeServiceUsers != "" {
		m[attributes.IncludeServiceUsers] = []string{p.IncludeServiceUsers}
	}
	return m
}

type UsersListResponse struct {
	Users []User `json:"list"`
}

// Commands used by both `human` and `service` users:

func (c OktaPAMClient) ListUsers(ctx context.Context, parameters ListUsersParameters) ([]User, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/users", url.PathEscape(c.Team))
	users := make([]User, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)
		resp, err := c.CreateBaseRequest(ctx).
			SetResult(&UsersListResponse{}).
			SetQueryParamsFromValues(parameters.toQueryParametersMap()).
			Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, http.StatusOK); err != nil {
			return nil, err
		}

		usersListResponse := resp.Result().(*UsersListResponse)
		users = append(users, usersListResponse.Users...)

		linkHeader := resp.Header().Get("Link")
		if linkHeader == "" {
			break
		}
		links := linkheader.Parse(linkHeader)
		requestURL = ""
		for _, link := range links {
			if link.Rel == "next" {
				requestURL = link.URL
			}
		}
		if requestURL == "" {
			break
		}
	}
	return users, nil
}

// Commands used by only `human` users:

func (c OktaPAMClient) GetHumanUser(ctx context.Context, userName string) (*User, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/users/%s", url.PathEscape(c.Team), url.PathEscape(userName))
	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&User{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == http.StatusOK {
		user := resp.Result().(*User)
		return user, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) CreateHumanUser(ctx context.Context, userName string) error {
	return fmt.Errorf(errors.HumanUserCreationError)
}

func (c OktaPAMClient) UpdateHumanUser(ctx context.Context, userName string, humanUser *User) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/users/%s", url.PathEscape(c.Team), url.PathEscape(userName))
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(humanUser).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, http.StatusOK)
	return err
}

func (c OktaPAMClient) DeleteHumanUser(ctx context.Context, userName string) error {
	return fmt.Errorf(errors.HumanUserDeletionError)
}

// Commands used by only `service` users:

func (c OktaPAMClient) ListServiceUsers(ctx context.Context, parameters ListUsersParameters) ([]User, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/service_users", url.PathEscape(c.Team))
	serviceUsers := make([]User, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)
		resp, err := c.CreateBaseRequest(ctx).
			SetResult(&UsersListResponse{}).
			SetQueryParamsFromValues(parameters.toQueryParametersMap()).
			Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, http.StatusOK); err != nil {
			return nil, err
		}

		serviceUsersListResponse := resp.Result().(*UsersListResponse)
		serviceUsers = append(serviceUsers, serviceUsersListResponse.Users...)

		linkHeader := resp.Header().Get("Link")
		if linkHeader == "" {
			break
		}
		links := linkheader.Parse(linkHeader)
		requestURL = ""
		for _, link := range links {
			if link.Rel == "next" {
				requestURL = link.URL
			}
		}
		if requestURL == "" {
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

	if statusCode, err := checkStatusCode(resp, http.StatusOK, http.StatusNoContent); err == nil {
		serviceUser := resp.Result().(*User)
		return serviceUser, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
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
	_, err = checkStatusCode(resp, http.StatusCreated)
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
	_, err = checkStatusCode(resp, http.StatusOK)
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
	_, err = checkStatusCode(resp, http.StatusNoContent, http.StatusNotFound)
	return err
}
