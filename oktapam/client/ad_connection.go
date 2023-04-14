package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
	"github.com/tomnomnom/linkheader"
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

type ListADConnectionsParameters struct {
	GatewayID          string
	CertificateID      string
	IncludeCertDetails bool
}

type ADConnectionsListResponse struct {
	ADConnections []ADConnection `json:"list"`
}

type ADUserSyncTaskSettings struct {
	ID              *string `json:"id,omitempty"`
	Name            *string `json:"name"`
	Frequency       *int    `json:"frequency"`
	StartHourUTC    *int    `json:"start_hour_utc,omitempty"`
	IsActive        *bool   `json:"is_active"`
	RunTest         *bool   `json:"run_test"`
	BaseDN          *string `json:"base_dn"`
	LDAPQueryFilter *string `json:"ldap_query_filter"`
	UPNField        *string `json:"upn_field"`
	SIDField        *string `json:"sid_field"`
}

type ADUserSyncTaskSettingsSchedule struct {
	Frequency    *int `json:"frequency"`
	StartHourUTC *int `json:"start_hour_utc"`
}

type ADUserSyncTaskSettingsState struct {
	IsActive *bool `json:"is_active"`
}

type ListADUserSyncTaskSettingsParameters struct {
	Status string
}

type ADUserSyncTaskSettingsListResponse struct {
	ADUserSyncTaskSettingsList []ADUserSyncTaskSettings `json:"list"`
}

func (adConn ADConnection) ToResourceMap() map[string]any {
	m := make(map[string]any)

	if adConn.Name != nil {
		m[attributes.Name] = *adConn.Name
	}
	if adConn.ID != nil {
		m[attributes.ID] = *adConn.ID
	}
	if adConn.GatewayID != nil {
		m[attributes.GatewayID] = *adConn.GatewayID
	}
	if adConn.Domain != nil {
		m[attributes.Domain] = *adConn.Domain
	}
	if adConn.ServiceAccountUsername != nil {
		m[attributes.ServiceAccountUsername] = *adConn.ServiceAccountUsername
	}
	if adConn.ServiceAccountPassword != nil {
		m[attributes.ServiceAccountPassword] = *adConn.ServiceAccountPassword
	}
	if adConn.DomainControllers != nil {
		m[attributes.DomainControllers] = adConn.DomainControllers
	}
	if adConn.UsePasswordless != nil {
		m[attributes.UsePasswordless] = *adConn.UsePasswordless
	}
	if adConn.CertificateId != nil {
		m[attributes.CertificateID] = *adConn.CertificateId
	}

	return m
}

func (adConn ADConnection) Exists() bool {
	return utils.IsNonEmpty(adConn.ID) && utils.IsBlank(adConn.DeletedAt)
}

func (adTaskSettings ADTaskSettings) Exists() bool {
	return utils.IsNonEmpty(adTaskSettings.ID)
}

func (adUserSyncTaskSettings ADUserSyncTaskSettings) Exists() bool {
	return utils.IsNonEmpty(adUserSyncTaskSettings.ID)
}

func (c OktaPAMClient) ListADConnections(ctx context.Context, parameters ListADConnectionsParameters) ([]ADConnection, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections", url.PathEscape(c.Team))
	adConnections := make([]ADConnection, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)

		resp, err := c.CreateBaseRequest(ctx).SetQueryParams(parameters.toQueryParametersMap()).SetResult(&ADConnectionsListResponse{}).Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, http.StatusOK); err != nil {
			return nil, err
		}

		adConnectionsListResponse := resp.Result().(*ADConnectionsListResponse)
		adConnections = append(adConnections, adConnectionsListResponse.ADConnections...)

		linkHeader := resp.Header().Get("Link")
		//No more results to fetch
		if linkHeader == "" {
			break
		}
		links := linkheader.Parse(linkHeader)
		requestURL = ""

		//Set the request url with next link
		for _, link := range links {
			if link.Rel == "next" {
				requestURL = link.URL
				break
			}
		}
	}

	return adConnections, nil
}

func (c OktaPAMClient) GetADConnection(ctx context.Context, id string, allowDeleted bool) (*ADConnection, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s", url.PathEscape(c.Team), url.PathEscape(id))
	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&ADConnection{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == http.StatusOK {
		adConn := resp.Result().(*ADConnection)
		if adConn.Exists() || allowDeleted {
			return adConn, nil
		}
		return nil, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) CreateADConnection(ctx context.Context, adConn ADConnection) (*ADConnection, error) {
	// Create the ad connection on the api server
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
	if _, err := checkStatusCode(resp, http.StatusNoContent); err != nil {
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

	_, err = checkStatusCode(resp, http.StatusNoContent, http.StatusNotFound)
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

	if statusCode == http.StatusOK {
		adTaskSettings := resp.Result().(*ADTaskSettings)
		if adTaskSettings.Exists() {
			return adTaskSettings, nil
		}
		return nil, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) CreateADTaskSettings(ctx context.Context, adConnId string, adTaskSettings ADTaskSettings) (*ADTaskSettings, error) {
	// Create the ad connection task settings on the api server
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s/task_settings", url.PathEscape(c.Team),
		url.PathEscape(adConnId))
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(adTaskSettings).SetResult(&ADTaskSettings{}).Post(requestURL)

	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err := checkStatusCode(resp, http.StatusCreated); err != nil {
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
	if _, err := checkStatusCode(resp, http.StatusNoContent); err != nil {
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

	_, err = checkStatusCode(resp, http.StatusNoContent, http.StatusNotFound)
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

	_, err = checkStatusCode(resp, http.StatusNoContent)
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

	_, err = checkStatusCode(resp, http.StatusNoContent)
	return err
}

func (p ListADConnectionsParameters) toQueryParametersMap() map[string]string {
	m := make(map[string]string, 3)

	if p.GatewayID != "" {
		m[attributes.GatewayID] = p.GatewayID
	}
	if p.CertificateID != "" {
		m[attributes.CertificateID] = p.CertificateID
	}
	if p.IncludeCertDetails {
		m[attributes.IncludeCertDetails] = strconv.FormatBool(p.IncludeCertDetails)
	}

	return m
}

func (adUserSyncTaskSettings ADUserSyncTaskSettings) ToResourceMap() map[string]interface{} {
	m := make(map[string]interface{})

	if adUserSyncTaskSettings.Name != nil {
		m[attributes.Name] = *adUserSyncTaskSettings.Name
	}
	if adUserSyncTaskSettings.ID != nil {
		m[attributes.ID] = *adUserSyncTaskSettings.ID
	}
	if adUserSyncTaskSettings.Frequency != nil {
		m[attributes.Frequency] = *adUserSyncTaskSettings.Frequency
	}
	if adUserSyncTaskSettings.StartHourUTC != nil {
		m[attributes.StartHourUTC] = *adUserSyncTaskSettings.StartHourUTC
	}
	if adUserSyncTaskSettings.BaseDN != nil {
		m[attributes.BaseDN] = *adUserSyncTaskSettings.BaseDN
	}
	if adUserSyncTaskSettings.LDAPQueryFilter != nil {
		m[attributes.LDAPQueryFilter] = *adUserSyncTaskSettings.LDAPQueryFilter
	}
	if adUserSyncTaskSettings.UPNField != nil {
		m[attributes.UPNField] = *adUserSyncTaskSettings.UPNField
	}
	if adUserSyncTaskSettings.SIDField != nil {
		m[attributes.SIDField] = *adUserSyncTaskSettings.SIDField
	}
	if adUserSyncTaskSettings.IsActive != nil {
		m[attributes.IsActive] = *adUserSyncTaskSettings.IsActive
	}

	return m
}
func (c OktaPAMClient) GetADUserSyncTaskSettings(ctx context.Context, adConnId string, adUserSyncTaskSettingsId string) (*ADUserSyncTaskSettings, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s/user_sync_task_settings/%s", url.PathEscape(c.Team),
		url.PathEscape(adConnId), url.PathEscape(adUserSyncTaskSettingsId))
	logging.Tracef("making GET request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetResult(&ADUserSyncTaskSettings{}).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	statusCode := resp.StatusCode()

	if statusCode == http.StatusOK {
		adUserSyncTaskSettings := resp.Result().(*ADUserSyncTaskSettings)
		if adUserSyncTaskSettings.Exists() {
			return adUserSyncTaskSettings, nil
		}
		return nil, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

func (c OktaPAMClient) CreateADUserSyncTaskSettings(ctx context.Context, adConnId string, adUserSyncTaskSettings ADUserSyncTaskSettings) (*ADUserSyncTaskSettings, error) {
	// Create the ad connection user sync task settings on the api server
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s/user_sync_task_settings", url.PathEscape(c.Team),
		url.PathEscape(adConnId))
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(adUserSyncTaskSettings).SetResult(&ADUserSyncTaskSettings{}).Post(requestURL)

	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}
	if _, err := checkStatusCode(resp, http.StatusCreated); err != nil {
		logging.Tracef("unexpected status code: %d", resp.StatusCode())
		return nil, err
	}

	createdADUserSyncTaskSettings := resp.Result().(*ADUserSyncTaskSettings)
	return createdADUserSyncTaskSettings, nil
}

func (c OktaPAMClient) DeleteADUserSyncTaskSettings(ctx context.Context, adConnId string, adUserSyncTaskSettingsId string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s/user_sync_task_settings/%s", url.PathEscape(c.Team),
		url.PathEscape(adConnId), url.PathEscape(adUserSyncTaskSettingsId))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusNoContent, http.StatusNotFound)
	return err
}

func (c OktaPAMClient) UpdateADUserSyncTaskSettingsSchedule(ctx context.Context, adConnId string, adUserSyncTaskSettingsId string,
	schedule ADUserSyncTaskSettingsSchedule) error {

	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s/user_sync_task_settings/%s/schedule", url.PathEscape(c.Team),
		url.PathEscape(adConnId), url.PathEscape(adUserSyncTaskSettingsId))
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(schedule).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusNoContent)
	return err
}

func (c OktaPAMClient) UpdateADUserSyncTaskSettingsState(ctx context.Context, adConnId string, adUserSyncTaskSettingsId string,
	state ADUserSyncTaskSettingsState) error {

	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s/user_sync_task_settings/%s/state", url.PathEscape(c.Team),
		url.PathEscape(adConnId), url.PathEscape(adUserSyncTaskSettingsId))
	logging.Tracef("making POST request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(state).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	_, err = checkStatusCode(resp, http.StatusNoContent)
	return err
}

func (c OktaPAMClient) ListADUserSyncTaskSettings(ctx context.Context, connectionID string, parameters ListADUserSyncTaskSettingsParameters) ([]ADUserSyncTaskSettings, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/integrations/ad_connections/%s/user_sync_task_settings", url.PathEscape(c.Team),
		connectionID)
	adUserSyncTaskSettingsList := make([]ADUserSyncTaskSettings, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)

		resp, err := c.CreateBaseRequest(ctx).SetQueryParams(parameters.toQueryParametersMap()).SetResult(&ADUserSyncTaskSettingsListResponse{}).Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, http.StatusOK); err != nil {
			return nil, err
		}

		adUserSyncTaskSettingsListResponse := resp.Result().(*ADUserSyncTaskSettingsListResponse)
		adUserSyncTaskSettingsList = append(adUserSyncTaskSettingsList, adUserSyncTaskSettingsListResponse.ADUserSyncTaskSettingsList...)

		linkHeader := resp.Header().Get("Link")
		//No more results to fetch
		if linkHeader == "" {
			break
		}
		links := linkheader.Parse(linkHeader)
		requestURL = ""

		//Set the request url with next link
		for _, link := range links {
			if link.Rel == "next" {
				requestURL = link.URL
				break
			}
		}
	}

	return adUserSyncTaskSettingsList, nil
}

func (p ListADUserSyncTaskSettingsParameters) toQueryParametersMap() map[string]string {
	m := make(map[string]string, 1)

	if p.Status != "" {
		m[attributes.ADUserSyncTaskSettingsStatus] = p.Status
	}

	return m
}
