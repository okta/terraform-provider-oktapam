package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

type TeamSettings struct {
	ReactivateUsersViaIDP           *bool   `json:"reactivate_users_via_idp,omitempty"`
	ApproveDeviceWithoutInteraction *bool   `json:"approve_device_without_interaction,omitempty"`
	UserProvisioningExactUserName   *bool   `json:"user_provisioning_exact_username,omitempty"`
	ClientSessionDuration           *int    `json:"client_session_duration,omitempty"`
	WebSessionDuration              *int    `json:"web_session_duration,omitempty"`
	IncludeUserSID                  *string `json:"include_user_sid,omitempty"`
}

func (s TeamSettings) ToResourceMap() map[string]any {
	m := make(map[string]any, 2)

	if s.ReactivateUsersViaIDP != nil {
		m[attributes.ReactivateUsersViaIDP] = *s.ReactivateUsersViaIDP
	}
	if s.ApproveDeviceWithoutInteraction != nil {
		m[attributes.ApproveDeviceWithoutInteraction] = *s.ApproveDeviceWithoutInteraction
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
		return settings, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) UpdateTeamSettings(ctx context.Context, settings TeamSettings) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/settings", url.PathEscape(c.Team))
	logging.Tracef("making PUT request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).SetBody(settings).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusOK)
	return err
}

func (c OktaPAMClient) CreateTeamSettings(ctx context.Context, settings TeamSettings) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/settings", url.PathEscape(c.Team))
	logging.Tracef("making PUT request to %s", requestURL)

	resp, err := c.CreateBaseRequest(ctx).SetBody(settings).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusOK)
	return err
}
