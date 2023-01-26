package client

import (
	"context"
	"fmt"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
	"net/http"
	"net/url"
)

type TeamSettings struct {
	Team                            *string `json:"team,omitempty"`
	ReactivateUsersViaIDP           *bool   `json:"reactivate_users_via_idp,omitempty"`
	ApproveDeviceWithoutInteraction *bool   `json:"approve_device_without_interaction,omitempty"`
	PostDeviceEnrollmentURL         *string `json:"post_device_enrollment_url,omitempty"`
	PostLogoutURL                   *string `json:"post_logout_url,omitempty"`
	PostLoginURL                    *string `json:"post_login_url,omitempty"`
	UserProvisioningExactUserName   *bool   `json:"user_provisioning_exact_username,omitempty"`
	ClientSessionDuration           *int    `json:"client_session_duration,omitempty"`
	WebSessionDuration              *int    `json:"web_session_duration,omitempty"`
	IncludeUserSID                  *string `json:"include_user_SID,omitempty"`
}

func (s TeamSettings) ToResourceMap() map[string]interface{} {
	m := make(map[string]interface{}, 2)

	if s.Team != nil {
		m[attributes.TeamName] = *s.Team
	}
	if s.ReactivateUsersViaIDP != nil {
		m[attributes.ReactivateUsersViaIDP] = *s.ReactivateUsersViaIDP
	}
	if s.ApproveDeviceWithoutInteraction != nil {
		m[attributes.ApproveDeviceWithoutInteraction] = *s.ApproveDeviceWithoutInteraction
	}
	if s.PostDeviceEnrollmentURL != nil {
		m[attributes.PostDeviceEnrollmentURL] = *s.PostDeviceEnrollmentURL
	}
	if s.PostLoginURL != nil {
		m[attributes.PostLoginURL] = *s.PostLoginURL
	}
	if s.PostLogoutURL != nil {
		m[attributes.PostLogoutURL] = *s.PostLogoutURL
	}
	if s.UserProvisioningExactUserName != nil {
		m[attributes.UserProvisioningExactUserName] = *s.UserProvisioningExactUserName
	}
	if s.ClientSessionDuration != nil {
		m[attributes.ClientSessionDuration] = *s.ClientSessionDuration
	}
	if s.WebSessionDuration != nil {
		m[attributes.WebSessionDuration] = *s.WebSessionDuration
	}
	if s.IncludeUserSID != nil {
		m[attributes.IncludeUserSID] = *s.IncludeUserSID
	}

	return m
}

func (s TeamSettings) Exists() bool {
	return utils.IsNonEmpty(s.Team)
}

func (c OktaPAMClient) GetTeamSettings(ctx context.Context) (*TeamSettings, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/settings", url.PathEscape(c.Team))
	logging.Tracef("making GET request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).SetResult(&TeamSettings{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == http.StatusOK {
		settings := resp.Result().(*TeamSettings)
		if settings.Exists() {
			return settings, nil
		}
		return nil, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) UpdateTeamSettings(ctx context.Context, teamSettings TeamSettings) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/settings", url.PathEscape(c.Team))
	logging.Tracef("making PUT request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).SetBody(teamSettings).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusNoContent)
	return err
}
