package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/logging"
	"github.com/terraform-providers/terraform-provider-oktapam/oktapam/utils"
)

type ADConnection struct {
	Name                   *string  `json:"name"`
	ID                     *string  `json:"id,omitempty"`
	GatewayID              *string  `json:"gateway_id"`
	Domain                 *string  `json:"domain"`
	ServiceAccountUsername *string  `json:"service_account_username"`
	ServiceAccountPassword *string  `json:"service_account_password"`
	DomainControllers      []string `json:"domain_controllers,omitempty"`
	UsePasswordless        *bool    `json:"use_passwordless"`
	CertificateId          *string  `json:"certificate_id,omitempty"`
	ActiveTaskSettingsId   *string  `json:"active_task_settings_id,omitempty"`
	DeletedAt              *string  `json:"deleted_at,omitempty"`
}

type ADTaskSettings struct {
	ID                         *string                  `json:"id,omitempty"`
	Name                       *string                  `json:"name"`
	Frequency                  *int                     `json:"frequency"`
	StartHourUTC               *int                     `json:"start_hour_utc,omitempty"`
	IsActive                   *bool                    `json:"is_active"`
	RunTest                    *bool                    `json:"run_test"`
	HostnameAttribute          *string                  `json:"hostname_attribute"`
	AccessAddressAttribute     *string                  `json:"access_address_attribute,omitempty"`
	OSAttribute                *string                  `json:"os_attribute"`
	BastionAttribute           *string                  `json:"bastion_attribute,omitempty"`
	AltNamesAttributes         []string                 `json:"alt_names_attributes,omitempty"`
	AdditionalAttributeMapping []*ADAdditionalAttribute `json:"additional_attribute_mapping,omitempty"`
	//RuleAssignments are not sorted in any order, consumers may want to sort by priority.
	RuleAssignments []*ADRuleAssignment `json:"rule_assignments"`
}

type ADAdditionalAttribute struct {
	Label  string `json:"label"`
	Value  string `json:"value"`
	IsGuid bool   `json:"is_guid"`
}

type ADRuleAssignment struct {
	ID              string `json:"id"`
	BaseDN          string `json:"base_dn"`
	LDAPQueryFilter string `json:"ldap_query_filter"`
	ProjectID       string `json:"project_id"`
	ProjectName     string `json:"project_name"`
	Priority        int    `json:"priority"`
}

type ADTaskSettingsSchedule struct {
	Frequency    *int `json:"frequency"`
	StartHourUTC *int `json:"start_hour_utc"`
}

func (adConn ADConnection) Exists() bool {
	return utils.IsNonEmpty(adConn.ID) && utils.IsBlank(adConn.DeletedAt)
}

func (adTaskSettings ADTaskSettings) Exists() bool {
	return utils.IsNonEmpty(adTaskSettings.ID)
}

func (c OktaPAMClient) GetADConnection(ctx context.Context, id string) (*ADConnection, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s", url.PathEscape(c.Team), url.PathEscape(id))
	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&ADConnection{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == 200 {
		adConn := resp.Result().(*ADConnection)
		if adConn.Exists() {
			return adConn, nil
		}
		return nil, nil
	} else if statusCode == 404 {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, 200, 404)
}

func (c OktaPAMClient) CreateADConnection(ctx context.Context, adConn ADConnection) (*ADConnection, error) {
	// Create the group on the api server specified by group
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections", url.PathEscape(c.Team))
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(adConn).SetResult(&ADConnection{}).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err := checkStatusCode(resp, 201); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return nil, err
	}
	createdADConnection := resp.Result().(*ADConnection)

	return createdADConnection, nil
}

func (c OktaPAMClient) UpdateADConnection(ctx context.Context, adConnId string, adConn ADConnection) (*ADConnection, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s", url.PathEscape(c.Team), url.PathEscape(adConnId))
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(adConn).SetResult(&ADConnection{}).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err := checkStatusCode(resp, 204); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return nil, err
	}
	updatedADConnection := resp.Result().(*ADConnection)
	return updatedADConnection, nil
}

func (c OktaPAMClient) DeleteADConnection(ctx context.Context, adConnId string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s", url.PathEscape(c.Team), url.PathEscape(adConnId))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, 204, 404)
	return err
}

func (c OktaPAMClient) GetADTaskSettings(ctx context.Context, adConnId string, adTaskSettingsId string) (*ADTaskSettings, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s/task_settings/%s", url.PathEscape(c.Team),
		url.PathEscape(adConnId), url.PathEscape(adTaskSettingsId))
	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&ADTaskSettings{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == 200 {
		adTaskSettings := resp.Result().(*ADTaskSettings)
		if adTaskSettings.Exists() {
			return adTaskSettings, nil
		}
		return nil, nil
	} else if statusCode == 404 {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, 200, 404)
}

func (c OktaPAMClient) CreateADTaskSettings(ctx context.Context, adConnId string, adTaskSettings ADTaskSettings) (*ADTaskSettings, error) {
	// Create the group on the api server specified by group
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s/task_settings", url.PathEscape(c.Team),
		url.PathEscape(adConnId))
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(adTaskSettings).SetResult(&ADTaskSettings{}).Post(requestURL)

	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err := checkStatusCode(resp, 201); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return nil, err
	}

	createdADTaskSettings := resp.Result().(*ADTaskSettings)
	return createdADTaskSettings, nil
}

func (c OktaPAMClient) UpdateADTaskSettings(ctx context.Context, adConnId string, adTaskSettingsId string, adTaskSettings ADTaskSettings) (*ADTaskSettings, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s/task_settings/%s", url.PathEscape(c.Team),
		url.PathEscape(adConnId), url.PathEscape(adTaskSettingsId))
	logging.Tracef("making PUT request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(adTaskSettings).SetResult(&ADTaskSettings{}).Put(requestURL)

	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err := checkStatusCode(resp, 204); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return nil, err
	}

	updatedADTaskSettings := resp.Result().(*ADTaskSettings)
	return updatedADTaskSettings, nil
}

func (c OktaPAMClient) DeleteADTaskSettings(ctx context.Context, adConnId string, adTaskSettingsId string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s/task_settings/%s", url.PathEscape(c.Team),
		url.PathEscape(adConnId), url.PathEscape(adTaskSettingsId))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, 204, 404)
	return err
}

func (c OktaPAMClient) UpdateADTaskSettingsSchedule(ctx context.Context, adConnId string, adTaskSettingsId string, schedule ADTaskSettingsSchedule) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s/task_settings/%s/schedule", url.PathEscape(c.Team),
		url.PathEscape(adConnId), url.PathEscape(adTaskSettingsId))
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(schedule).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, 204)
	return err
}

func (c OktaPAMClient) DeactivateADTaskSettings(ctx context.Context, adConnId string, adTaskSettingsId string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s/task_settings/%s/deactivate", url.PathEscape(c.Team),
		url.PathEscape(adConnId), url.PathEscape(adTaskSettingsId))
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, 204)
	return err
}
