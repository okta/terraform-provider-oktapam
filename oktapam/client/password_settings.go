package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

type PasswordSettings struct {
	EnablePeriodicRotation            *bool             `json:"enable_periodic_rotation"`
	MinLengthInBytes                  *int              `json:"min_length_in_bytes,omitempty"`
	MaxLengthInBytes                  *int              `json:"max_length_in_bytes,omitempty"`
	PeriodicRotationDurationInSeconds *int              `json:"periodic_rotation_duration_in_seconds,omitempty"`
	CharacterOptions                  *CharacterOptions `json:"character_options,omitempty"`
	ManagedPrivilegedAccountsConfig   []string          `json:"managed_privileged_accounts_config"`
	ModifiedAt                        *string           `json:"modified_at,omitempty"`
}

type CharacterOptions struct {
	UpperCase   *bool `json:"upper_case,omitempty"`
	LowerCase   *bool `json:"lower_case,omitempty"`
	Digits      *bool `json:"digits,omitempty"`
	Punctuation *bool `json:"punctuation,omitempty"`
}

func (s PasswordSettings) ToResourceMap() map[string]any {
	m := make(map[string]any, 7)

	if s.EnablePeriodicRotation != nil {
		m[attributes.EnablePeriodicRotation] = *s.EnablePeriodicRotation
	}
	if s.MinLengthInBytes != nil {
		m[attributes.MinLength] = int(*s.MinLengthInBytes)
	}
	if s.MaxLengthInBytes != nil {
		m[attributes.MaxLength] = int(*s.MaxLengthInBytes)
	}
	if s.PeriodicRotationDurationInSeconds != nil {
		m[attributes.PeriodicRotationDurationInSeconds] = int(*s.PeriodicRotationDurationInSeconds)
	}

	acctsArr := make([]any, len(s.ManagedPrivilegedAccountsConfig))
	for idx, a := range s.ManagedPrivilegedAccountsConfig {
		acctsArr[idx] = a
	}
	m[attributes.ManagedPrivilegedAccounts] = acctsArr

	if s.CharacterOptions != nil {
		cArr := make([]any, 1)
		cArr[0] = s.CharacterOptions.ToResourceMap()
		m[attributes.CharacterOptions] = cArr
	} else {
		m[attributes.CharacterOptions] = make([]any, 0)
	}

	return m
}

func (o CharacterOptions) ToResourceMap() map[string]any {
	m := make(map[string]any, 4)

	if o.LowerCase != nil {
		m[attributes.LowerCase] = *o.LowerCase
	}
	if o.UpperCase != nil {
		m[attributes.UpperCase] = *o.UpperCase
	}
	if o.Digits != nil {
		m[attributes.Digits] = *o.Digits
	}
	if o.Punctuation != nil {
		m[attributes.Punctuation] = *o.Punctuation
	}

	return m
}

func (c OktaPAMClient) GetPasswordSettings(ctx context.Context, resourceGroupID string, projectID string) (*PasswordSettings, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s/server_password_settings", url.PathEscape(c.Team), url.PathEscape(resourceGroupID), url.PathEscape(projectID))
	logging.Tracef("making GET request to %s", requestURL)
	passwordSettings := PasswordSettings{}
	resp, err := c.CreateBaseRequest(ctx).SetResult(&passwordSettings).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == http.StatusOK {
		return &passwordSettings, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) UpdatePasswordSettings(ctx context.Context, resourceGroupID string, projectID string, passwordSettings *PasswordSettings) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/resource_groups/%s/projects/%s/server_password_settings", url.PathEscape(c.Team), url.PathEscape(resourceGroupID), url.PathEscape(projectID))
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(passwordSettings).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusOK)
	return err
}
