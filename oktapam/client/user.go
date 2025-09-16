package client

import (
	"context"
	goerrors "errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/errors"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"

	"github.com/okta/terraform-provider-oktapam/oktapam/utils"

	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/tomnomnom/linkheader"
)

type User struct {
	ID             *string                   `json:"id"`
	Name           *string                   `json:"name"`
	TeamName       *string                   `json:"team_name"`
	ServerUserName *string                   `json:"server_user_name,omitempty"`
	DeletedAt      *string                   `json:"deleted_at,omitempty"`
	Status         *typed_strings.UserStatus `json:"status"`
	UserType       *typed_strings.UserType   `json:"user_type"`
}

func UserFromMap(m map[string]any) (*User, error) {
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

func (su User) ToResourceMap() map[string]any {
	m := make(map[string]any, 2)

	if su.ID != nil {
		m[attributes.ID] = *su.ID
	}
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

func ListServiceUsers(ctx context.Context, sdkClient SDKClientWrapper) ([]User, error) {
	request := sdkClient.SDKClient.ServiceUsersAPI.ListServiceUsers(ctx, sdkClient.Team)

	users := make([]User, 0, 5)

	for {
		resp, httpResp, callErr := request.Execute()
		if _, err := checkStatusCodeFromSDK(httpResp, 200); err != nil {
			return nil, err
		} else if callErr != nil {
			return nil, callErr
		}

		for _, user := range resp.List {
			usr := User{
				ID:             &user.Id,
				Name:           &user.Name,
				TeamName:       &user.TeamName,
				UserType:       (*typed_strings.UserType)(&user.UserType),
				Status:         (*typed_strings.UserStatus)(&user.Status),
				ServerUserName: &user.Name,
			}
			users = append(users, usr)
		}

		linkHeader := httpResp.Header.Get("Link")
		if linkHeader == "" {
			break
		}

		links := linkheader.Parse(linkHeader)

		var nextURLStr string
		for _, link := range links {
			if link.Rel == "next" {
				nextURLStr = link.URL
			}
		}

		if nextURLStr == "" {
			break
		}

		nextURL, err := url.Parse(nextURLStr)
		if err != nil {
			return nil, err
		}
		values := nextURL.Query()
		offset, ok := values["offset"]
		if !ok || len(offset) != 1 {
			return nil, fmt.Errorf("invalid offset")
		}

		request = request.Offset(offset[0])
	}

	return users, nil
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
	return goerrors.New(errors.HumanUserCreationError)
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

func (c OktaPAMClient) DeleteHumanUser(_ context.Context, _ string) error {
	return goerrors.New(errors.HumanUserCreationError)
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

	request := map[string]any{
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

// NOTE: ASA does not support user hard-deletion, only soft-deletion by updating the status to `DELETED`.
// The main reason for this atypical behavior is that the user name is used to provision the corresponding account across servers.
// Following the removal of user permissions, the current behavior is to keep the user account on server while denying any short-lived certificates to that user for that server.
// Soft-deletion prevents re-usage of the same user name, and with it the potential of the new user usurping the old user account.
func (c OktaPAMClient) DeleteServiceUser(ctx context.Context, userName string, serviceUser *User) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/service_users/%s", url.PathEscape(c.Team), url.PathEscape(userName))
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(serviceUser).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}
	_, err = checkStatusCode(resp, http.StatusOK, http.StatusNotFound)
	return err
}

func GetCurrentUser(ctx context.Context, sdkClient SDKClientWrapper) (*pam.CurrentUserInfo, error) {
	userInfo, httpResp, err := sdkClient.SDKClient.UsersAPI.GetCurrentUserInfo(ctx, sdkClient.Team).Execute()
	if err != nil {
		return nil, err
	}

	if _, err := checkStatusCodeFromSDK(httpResp, 200); err != nil {
		return nil, err
	}

	return userInfo, nil
}

func AddUserToGroup(ctx context.Context, sdkClient SDKClientWrapper, groupName string, username string) error {
	request := sdkClient.SDKClient.GroupsAPI.AddUserToGroup(ctx, sdkClient.Team, groupName)
	request = request.AddUserToGroupRequest(pam.AddUserToGroupRequest{
		Name: &username,
	})

	resp, err := request.Execute()
	if err != nil {
		return err
	}

	if _, err := checkStatusCodeFromSDK(resp, 204); err != nil {
		return err
	}

	return nil
}

func RemoveUserFromGroup(ctx context.Context, sdkClient SDKClientWrapper, groupName string, username string) error {
	request := sdkClient.SDKClient.GroupsAPI.RemoveUserFromGroup(ctx, sdkClient.Team, groupName, username)
	resp, err := request.Execute()
	if err != nil {
		return err
	}

	if _, err := checkStatusCodeFromSDK(resp, 204); err != nil {
		return err
	}

	return nil
}

func GroupContainsUser(ctx context.Context, sdkClient SDKClientWrapper, groupName string, username string) (bool, error) {
	users, err := ListUsersInGroup(ctx, sdkClient, groupName, true)
	if err != nil {
		return false, err
	}

	for _, usr := range users {
		if *usr.Name == username {
			return true, nil
		}
	}

	return false, nil
}

func ListUsersInGroup(ctx context.Context, sdkClient SDKClientWrapper, groupName string, allow404 bool) ([]User, error) {
	var list []User
	request := sdkClient.SDKClient.GroupsAPI.ListUsersInGroup(ctx, sdkClient.Team, groupName)

	for {
		resp, httpResp, callErr := request.Execute()
		if httpResp != nil {
			if allow404 {
				if httpResp.StatusCode == 404 {
					return nil, nil
				}
			}
			if _, err := checkStatusCodeFromSDK(httpResp, 200); err != nil {
				return nil, err
			}
		} else if callErr != nil {
			return nil, callErr
		}

		for _, u := range resp.List {
			usr := User{
				ID:       &u.Id,
				Name:     &u.Name,
				TeamName: &u.TeamName,
				UserType: (*typed_strings.UserType)(&u.UserType),
			}

			list = append(list, usr)
		}

		linkHeader := httpResp.Header.Get("Link")
		if linkHeader == "" {
			break
		}

		links := linkheader.Parse(linkHeader)

		var nextURLStr string
		for _, link := range links {
			if link.Rel == "next" {
				nextURLStr = link.URL
			}
		}

		if nextURLStr == "" {
			break
		}

		nextURL, err := url.Parse(nextURLStr)
		if err != nil {
			return nil, err
		}
		values := nextURL.Query()
		offset, ok := values["offset"]
		if !ok || len(offset) != 1 {
			return nil, fmt.Errorf("invalid offset")
		}

		request = request.Offset(offset[0])
	}

	return list, nil
}
