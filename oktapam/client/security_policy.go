package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
	"github.com/tomnomnom/linkheader"
)

type PrivilegeType string
type ServerBasedResourceSubSelectorType string
type SecretBasedResourceSubSelectorType string
type ResourceSelectorType string
type AccountSelectorType string
type ConditionType string

const (
	PrincipalAccountSSHPrivilegeType = PrivilegeType("principal_account_ssh")
	PrincipalAccountRDPPrivilegeType = PrivilegeType("principal_account_rdp")
	PasswordCheckoutSSHPrivilegeType = PrivilegeType("password_checkout_ssh")
	PasswordCheckoutRDPPrivilegeType = PrivilegeType("password_checkout_rdp")
	SecretPrivilegeType              = PrivilegeType("secret")

	ServerBasedResourceSelectorType = ResourceSelectorType("server_based_resource")
	SecretBasedResourceSelectorType = ResourceSelectorType("secret_based_resource")

	IndividualServerSubSelectorType        = ServerBasedResourceSubSelectorType("individual_server")
	IndividualServerAccountSubSelectorType = ServerBasedResourceSubSelectorType("individual_server_account")
	ServerLabelServerSubSelectorType       = ServerBasedResourceSubSelectorType("server_label")

	SecretSubSelectorType       = SecretBasedResourceSubSelectorType("secret")
	SecretFolderSubSelectorType = SecretBasedResourceSubSelectorType("secret_folder")

	UsernameAccountSelectorType = AccountSelectorType("username")
	NoneAccountSelectorType     = AccountSelectorType("none")

	AccessRequestConditionType = ConditionType("access_request")
	GatewayConditionType       = ConditionType("gateway")
	MFAConditionType           = ConditionType("mfa")
)

// Deprecated: use security_policy2.go instead.
type SecurityPolicy struct {
	ID            *string                   `json:"id,omitempty"`
	Name          *string                   `json:"name"`
	Description   *string                   `json:"description,omitempty"`
	Active        *bool                     `json:"active"`
	ResourceGroup *NamedObject              `json:"resource_group,omitempty"`
	Principals    *SecurityPolicyPrincipals `json:"principals"`
	Rules         []*SecurityPolicyRule     `json:"rules"`
}

// Deprecated: use security_policy2.go instead.
func (p SecurityPolicy) ToResourceMap() map[string]any {
	m := make(map[string]any)

	if p.ID != nil {
		m[attributes.ID] = *p.ID
	}
	if p.Name != nil {
		m[attributes.Name] = *p.Name
	}
	if p.Description != nil {
		m[attributes.Description] = *p.Description
	}
	if p.Active != nil {
		m[attributes.Active] = *p.Active
	}
	if p.ResourceGroup != nil {
		m[attributes.ResourceGroup] = *p.ResourceGroup.Id
	} else {
		m[attributes.ResourceGroup] = ""
	}
	if p.Principals != nil {
		principals := make([]any, 1)

		principals[0] = p.Principals.ToResourceMap()

		m[attributes.Principals] = principals
	}
	if p.Rules != nil {
		rules := make([]any, len(p.Rules))

		for idx, rule := range p.Rules {
			rules[idx] = rule.ToResourceMap()
		}

		sort.Slice(rules, func(i, j int) bool {
			return *p.Rules[i].Name < *p.Rules[j].Name
		})

		m[attributes.Rule] = rules
	}

	return m
}

// Deprecated: use security_policy2.go instead.
type SecurityPoliciesListResponse struct {
	SecurityPolicies []SecurityPolicy `json:"list"`
}

// Deprecated: use security_policy2.go instead.
func (c OktaPAMClient) ListSecurityPolicies(ctx context.Context) ([]SecurityPolicy, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/security_policy", url.PathEscape(c.Team))
	securityPolicies := make([]SecurityPolicy, 0)

	for {
		// List will paginate, so we make a request, add results to array to return, check if we get a next page, and if so loop again
		logging.Tracef("making GET request to %s", requestURL)

		resp, err := c.CreateBaseRequest(ctx).SetResult(&SecurityPoliciesListResponse{}).Get(requestURL)
		if err != nil {
			logging.Errorf("received error while making request to %s", requestURL)
			return nil, err
		}
		if _, err := checkStatusCode(resp, http.StatusOK); err != nil {
			return nil, err
		}

		securityPoliciesListResponse := resp.Result().(*SecurityPoliciesListResponse)
		securityPolicies = append(securityPolicies, securityPoliciesListResponse.SecurityPolicies...)

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

	return securityPolicies, nil
}

// Deprecated: use security_policy2.go instead.
type SecurityPolicyPrincipals struct {
	UserGroups []NamedObject `json:"user_groups"`
}

// Deprecated: use security_policy2.go instead.
func (p *SecurityPolicyPrincipals) ToResourceMap() map[string]any {
	m := make(map[string]any)

	groupIds := make([]any, len(p.UserGroups))
	for idx, groupId := range p.UserGroups {
		g := groupId
		groupIds[idx] = *g.Id
	}
	m[attributes.Groups] = groupIds

	return m
}

// Deprecated: use security_policy2.go instead.
type SecurityPolicyRuleConditionContainer struct {
	ConditionType  ConditionType               `json:"condition_type"`
	ConditionValue SecurityPolicyRuleCondition `json:"condition_value"`
}

// Deprecated: use security_policy2.go instead.
func (c *SecurityPolicyRuleConditionContainer) UnmarshalJSON(data []byte) error {
	tmp := struct {
		ConditionType  ConditionType   `json:"condition_type"`
		ConditionValue json.RawMessage `json:"condition_value"`
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	c.ConditionType = tmp.ConditionType
	switch tmp.ConditionType {
	case AccessRequestConditionType:
		c.ConditionValue = &AccessRequestCondition{}
	case GatewayConditionType:
		c.ConditionValue = &GatewayCondition{}
	case MFAConditionType:
		c.ConditionValue = &MFACondition{}
	default:
		return fmt.Errorf("received unknown condition type: %s", tmp.ConditionType)
	}

	if err := json.Unmarshal(tmp.ConditionValue, c.ConditionValue); err != nil {
		return err
	}

	return nil
}

// Deprecated: use security_policy2.go instead.
type SecurityPolicyRuleCondition interface {
	ConditionType() ConditionType
	ValidForResourceType(resourceSelectorType ResourceSelectorType) bool
	ToResourceMap() map[string]any
}

// Deprecated: use security_policy2.go instead.
type MFACondition struct {
	ReAuthFrequencyInSeconds *int    `json:"re_auth_frequency_in_seconds"`
	ACRValues                *string `json:"acr_values"`
}

// Deprecated: use security_policy2.go instead.
func (c *MFACondition) ToResourceMap() map[string]any {
	m := make(map[string]any)

	m[attributes.ReAuthFrequencyInSeconds] = *c.ReAuthFrequencyInSeconds
	m[attributes.ACRValues] = *c.ACRValues

	return m
}

// Deprecated: use security_policy2.go instead.
func (*MFACondition) ConditionType() ConditionType {
	return MFAConditionType
}

// Deprecated: use security_policy2.go instead.
func (*MFACondition) ValidForResourceType(resourceSelectorType ResourceSelectorType) bool {
	return true
}

// Deprecated: use security_policy2.go instead.
type AccessRequestCondition struct {
	RequestTypeID       *string `json:"request_type_id"`
	RequestTypeName     *string `json:"request_type_name"`
	ExpiresAfterSeconds *int    `json:"expires_after_seconds"`
}

// Deprecated: use security_policy2.go instead.
func (c *AccessRequestCondition) ToResourceMap() map[string]any {
	m := make(map[string]any)

	m[attributes.RequestTypeId] = *c.RequestTypeID
	m[attributes.RequestTypeName] = *c.RequestTypeName
	if c.ExpiresAfterSeconds != nil {
		m[attributes.ExpiresAfterSeconds] = *c.ExpiresAfterSeconds
	}

	return m
}

// Deprecated: use security_policy2.go instead.
func (*AccessRequestCondition) ConditionType() ConditionType {
	return AccessRequestConditionType
}

// Deprecated: use security_policy2.go instead.
func (*AccessRequestCondition) ValidForResourceType(resourceSelectorType ResourceSelectorType) bool {
	return true
}

// Deprecated: use security_policy2.go instead.
type GatewayCondition struct {
	TrafficForwarding *bool `json:"traffic_forwarding"`
	SessionRecording  *bool `json:"session_recording"`
}

// Deprecated: use security_policy2.go instead.
func (c *GatewayCondition) ToResourceMap() map[string]any {
	m := make(map[string]any)

	m[attributes.TrafficForwarding] = *c.TrafficForwarding
	m[attributes.SessionRecording] = *c.SessionRecording

	return m
}

// Deprecated: use security_policy2.go instead.
func (*GatewayCondition) ConditionType() ConditionType {
	return GatewayConditionType
}

// Deprecated: use security_policy2.go instead.
func (*GatewayCondition) ValidForResourceType(resourceSelectorType ResourceSelectorType) bool {
	return resourceSelectorType == ServerBasedResourceSelectorType
}

// Deprecated: use security_policy2.go instead.
type SecurityPolicyRuleResourceSelector interface {
	ResourceSelectorType() ResourceSelectorType
	ToResourceMap() map[string]any
}

// Deprecated: use security_policy2.go instead.
type ServerBasedResourceSelector struct {
	Selectors []ServerBasedResourceSubSelectorContainer `json:"selectors"`
}

// Deprecated: use security_policy2.go instead.
func (s *ServerBasedResourceSelector) ToResourceMap() map[string]any {
	m := make(map[string]any, 1)

	serversArr := make([]any, 1)
	subSelectorsMap := make(map[string]any, 3)
	serversArr[0] = subSelectorsMap

	individualServersArr := make([]any, 0, len(s.Selectors))
	individualServerAccountsArr := make([]any, 0, len(s.Selectors))
	serverLabelsArr := make([]any, 0, 1)

	for _, subSelector := range s.Selectors {
		switch subSelector.SelectorType {
		case IndividualServerSubSelectorType:
			individualServersArr = append(individualServersArr, subSelector.Selector.ToResourceMap())
		case IndividualServerAccountSubSelectorType:
			individualServerAccountsArr = append(individualServerAccountsArr, subSelector.Selector.ToResourceMap())
		case ServerLabelServerSubSelectorType:
			serverLabelsArr = append(serverLabelsArr, subSelector.Selector.ToResourceMap())
		}
	}

	subSelectorsMap[attributes.Server] = individualServersArr
	subSelectorsMap[attributes.ServerAccount] = individualServerAccountsArr
	subSelectorsMap[attributes.LabelSelectors] = serverLabelsArr

	m[attributes.Servers] = serversArr
	return m
}

// Deprecated: use security_policy2.go instead.
type ServerBasedResourceSubSelector interface {
	ServerBasedResourceSubSelectorType() ServerBasedResourceSubSelectorType
	ToResourceMap() map[string]any
}

// Deprecated: use security_policy2.go instead.
type IndividualServerSubSelector struct {
	Server NamedObject `json:"server"`
}

// Deprecated: use security_policy2.go instead.
func (*IndividualServerSubSelector) ServerBasedResourceSubSelectorType() ServerBasedResourceSubSelectorType {
	return IndividualServerSubSelectorType
}

// Deprecated: use security_policy2.go instead.
func (s *IndividualServerSubSelector) ToResourceMap() map[string]any {
	m := make(map[string]any)
	m[attributes.ServerID] = *s.Server.Id
	return m
}

// Deprecated: use security_policy2.go instead.
type IndividualServerAccountSubSelector struct {
	Server   NamedObject `json:"server"`
	Username *string     `json:"username"`
}

// Deprecated: use security_policy2.go instead.
func (*IndividualServerAccountSubSelector) ServerBasedResourceSubSelectorType() ServerBasedResourceSubSelectorType {
	return IndividualServerAccountSubSelectorType
}

// Deprecated: use security_policy2.go instead.
func (s *IndividualServerAccountSubSelector) ToResourceMap() map[string]any {
	m := make(map[string]any)
	m[attributes.ServerID] = *s.Server.Id
	m[attributes.Account] = *s.Username
	return m
}

// Deprecated: use security_policy2.go instead.
type ServerLabelBasedSubSelector struct {
	ServerSelector *ServerLabelServerSelector `json:"server_selector"`

	AccountSelectorType AccountSelectorType `json:"account_selector_type"`
	AccountSelector     AccountSelector     `json:"account_selector"`
}

// Deprecated: use security_policy2.go instead.
func (*ServerLabelBasedSubSelector) ServerBasedResourceSubSelectorType() ServerBasedResourceSubSelectorType {
	return ServerLabelServerSubSelectorType
}

// Deprecated: use security_policy2.go instead.
func (s *ServerLabelBasedSubSelector) ToResourceMap() map[string]any {
	m := make(map[string]any)

	serverLabelsM := make(map[string]any, len(s.ServerSelector.Labels))
	for k, v := range s.ServerSelector.Labels {
		serverLabelsM[k] = v
	}
	m[attributes.ServerLabels] = serverLabelsM

	usernamesArr := make([]any, 0)

	if s.AccountSelectorType == UsernameAccountSelectorType {
		usernamesArr = stringSliceToInterfaceSlice(s.AccountSelector.(*UsernameAccountSelector).Usernames)
	}

	m[attributes.Accounts] = usernamesArr

	return m
}

// Deprecated: use security_policy2.go instead.
func (ss *ServerLabelBasedSubSelector) UnmarshalJSON(data []byte) error {
	tmp := struct {
		ServerSelector      *ServerLabelServerSelector `json:"server_selector"`
		AccountSelectorType AccountSelectorType        `json:"account_selector_type"`
		AccountSelector     json.RawMessage            `json:"account_selector"`
	}{}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	ss.ServerSelector = tmp.ServerSelector
	ss.AccountSelectorType = tmp.AccountSelectorType

	switch ss.AccountSelectorType {
	case UsernameAccountSelectorType:
		ss.AccountSelector = &UsernameAccountSelector{}
		if err := json.Unmarshal(tmp.AccountSelector, ss.AccountSelector); err != nil {
			return err
		}
	case NoneAccountSelectorType:
		ss.AccountSelector = &NoneAccountSelector{}
	default:
		return fmt.Errorf("unknown account selector type %s", ss.AccountSelectorType)
	}

	return nil
}

// Deprecated: use security_policy2.go instead.
type ServerLabelServerSelector struct {
	Labels map[string]string `json:"labels"`
}

// Deprecated: use security_policy2.go instead.
type AccountSelector interface {
	AccountSelectorType() AccountSelectorType
}

// Deprecated: use security_policy2.go instead.
type NoneAccountSelector struct{}

// Deprecated: use security_policy2.go instead.
func (*NoneAccountSelector) AccountSelectorType() AccountSelectorType {
	return NoneAccountSelectorType
}

// Deprecated: use security_policy2.go instead.
type UsernameAccountSelector struct {
	Usernames []string `json:"usernames"`
}

// Deprecated: use security_policy2.go instead.
func (*UsernameAccountSelector) AccountSelectorType() AccountSelectorType {
	return UsernameAccountSelectorType
}

// Deprecated: use security_policy2.go instead.
type ServerBasedResourceSubSelectorContainer struct {
	SelectorType ServerBasedResourceSubSelectorType `json:"selector_type"`
	Selector     ServerBasedResourceSubSelector     `json:"selector"`
}

// Deprecated: use security_policy2.go instead.
func (c *ServerBasedResourceSubSelectorContainer) UnmarshalJSON(data []byte) error {
	tmp := struct {
		SelectorType ServerBasedResourceSubSelectorType `json:"selector_type"`
		Selector     json.RawMessage                    `json:"selector"`
	}{}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	c.SelectorType = tmp.SelectorType
	switch tmp.SelectorType {
	case IndividualServerSubSelectorType:
		c.Selector = &IndividualServerSubSelector{}
	case IndividualServerAccountSubSelectorType:
		c.Selector = &IndividualServerAccountSubSelector{}
	case ServerLabelServerSubSelectorType:
		c.Selector = &ServerLabelBasedSubSelector{}
	default:
		return fmt.Errorf("received unknown sub-selector type: %s", tmp.SelectorType)
	}
	if err := json.Unmarshal(tmp.Selector, c.Selector); err != nil {
		return err
	}

	return nil
}

// Deprecated: use security_policy2.go instead.
func (*ServerBasedResourceSelector) ResourceSelectorType() ResourceSelectorType {
	return ServerBasedResourceSelectorType
}

// Deprecated: use security_policy2.go instead.
type SecretBasedResourceSelector struct {
	Selectors []SecretBasedResourceSubSelectorContainer `json:"selectors"`
}

// Deprecated: use security_policy2.go instead.
func (s *SecretBasedResourceSelector) ToResourceMap() map[string]any {
	m := make(map[string]any, 1)

	secretsArr := make([]any, 1)
	subSelectorsMap := make(map[string]any, 3)
	secretsArr[0] = subSelectorsMap

	secretArr := make([]any, 0, len(s.Selectors))
	secretFolderArr := make([]any, 0, len(s.Selectors))

	for _, subSelector := range s.Selectors {
		switch subSelector.SelectorType {
		case SecretSubSelectorType:
			secretArr = append(secretArr, subSelector.Selector.ToResourceMap())
		case SecretFolderSubSelectorType:
			secretFolderArr = append(secretFolderArr, subSelector.Selector.ToResourceMap())
		}
	}

	subSelectorsMap[attributes.Secret] = secretArr
	subSelectorsMap[attributes.SecretFolder] = secretFolderArr

	m[attributes.Secrets] = secretsArr
	return m
}

// Deprecated: use security_policy2.go instead.
func (*SecretBasedResourceSelector) ResourceSelectorType() ResourceSelectorType {
	return SecretBasedResourceSelectorType
}

// Deprecated: use security_policy2.go instead.
type SecretBasedResourceSubSelector interface {
	SecretBasedResourceSubSelectorType() SecretBasedResourceSubSelectorType
	ToResourceMap() map[string]any
}

// Deprecated: use security_policy2.go instead.
type SecretSubSelector struct {
	SecretID NamedObject `json:"secret"`
}

// Deprecated: use security_policy2.go instead.
func (*SecretSubSelector) SecretBasedResourceSubSelectorType() SecretBasedResourceSubSelectorType {
	return SecretSubSelectorType
}

// Deprecated: use security_policy2.go instead.
func (s *SecretSubSelector) ToResourceMap() map[string]any {
	m := make(map[string]any)

	m[attributes.SecretID] = s.SecretID.Id

	return m
}

// Deprecated: use security_policy2.go instead.
type SecretFolderSubSelector struct {
	SecretFolderID NamedObject `json:"secret_folder"`
}

// Deprecated: use security_policy2.go instead.
func (*SecretFolderSubSelector) SecretBasedResourceSubSelectorType() SecretBasedResourceSubSelectorType {
	return SecretFolderSubSelectorType
}

// Deprecated: use security_policy2.go instead.
func (s *SecretFolderSubSelector) ToResourceMap() map[string]any {
	m := make(map[string]any)

	m[attributes.SecretFolderID] = s.SecretFolderID.Id

	return m
}

// Deprecated: use security_policy2.go instead.
type SecretBasedResourceSubSelectorContainer struct {
	SelectorType SecretBasedResourceSubSelectorType `json:"selector_type"`
	Selector     SecretBasedResourceSubSelector     `json:"selector"`
}

// Deprecated: use security_policy2.go instead.
func (c *SecretBasedResourceSubSelectorContainer) UnmarshalJSON(data []byte) error {
	tmp := struct {
		SelectorType SecretBasedResourceSubSelectorType `json:"selector_type"`
		Selector     json.RawMessage                    `json:"selector"`
	}{}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	c.SelectorType = tmp.SelectorType
	switch tmp.SelectorType {
	case SecretSubSelectorType:
		c.Selector = &SecretSubSelector{}
	case SecretFolderSubSelectorType:
		c.Selector = &SecretFolderSubSelector{}
	default:
		return fmt.Errorf("received unknown sub-selector type: %s", tmp.SelectorType)
	}
	if err := json.Unmarshal(tmp.Selector, c.Selector); err != nil {
		return err
	}

	return nil
}

// Deprecated: use security_policy2.go instead.
type SecurityPolicyRulePrivilegeContainer struct {
	PrivilegeType  PrivilegeType               `json:"privilege_type"`
	PrivilegeValue SecurityPolicyRulePrivilege `json:"privilege_value"`
}

// Deprecated: use security_policy2.go instead.
func (c *SecurityPolicyRulePrivilegeContainer) ToResourceMap() map[string]any {
	m := make(map[string]any, 4)

	return m
}

// Deprecated: use security_policy2.go instead.
func (c *SecurityPolicyRulePrivilegeContainer) UnmarshalJSON(data []byte) error {
	tmp := struct {
		PrivilegeType  PrivilegeType   `json:"privilege_type"`
		PrivilegeValue json.RawMessage `json:"privilege_value"`
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	c.PrivilegeType = tmp.PrivilegeType
	switch tmp.PrivilegeType {
	case PrincipalAccountSSHPrivilegeType:
		c.PrivilegeValue = &PrincipalAccountSSHPrivilege{}
	case PrincipalAccountRDPPrivilegeType:
		c.PrivilegeValue = &PrincipalAccountRDPPrivilege{}
	case PasswordCheckoutSSHPrivilegeType:
		c.PrivilegeValue = &PasswordCheckoutSSHPrivilege{}
	case PasswordCheckoutRDPPrivilegeType:
		c.PrivilegeValue = &PasswordCheckoutRDPPrivilege{}
	case SecretPrivilegeType:
		c.PrivilegeValue = &SecretPrivilege{}
	default:
		return fmt.Errorf("received unknown privilege type: %s", tmp.PrivilegeValue)
	}

	if err := json.Unmarshal(tmp.PrivilegeValue, c.PrivilegeValue); err != nil {
		return err
	}
	return nil
}

type SecurityPolicyRulePrivilege interface {
	ValidForResourceType(ResourceSelectorType) bool
	ToResourceMap() map[string]any
}

type PrincipalAccountRDPPrivilege struct {
	Enabled               *bool `json:"principal_account_rdp"`
	AdminLevelPermissions *bool `json:"admin_level_permissions"`
}

// Deprecated: Use security_policy2.go instead.
func (*PrincipalAccountRDPPrivilege) ValidForResourceType(resourceSelectorType ResourceSelectorType) bool {
	return resourceSelectorType == ServerBasedResourceSelectorType
}

// Deprecated: Use security_policy2.go instead.
func (p *PrincipalAccountRDPPrivilege) ToResourceMap() map[string]any {
	m := make(map[string]any, 2)
	m[attributes.Enabled] = *p.Enabled
	if p.AdminLevelPermissions != nil {
		m[attributes.AdminLevelPermissions] = *p.AdminLevelPermissions
	} else {
		m[attributes.AdminLevelPermissions] = false
	}
	return m
}

type PrincipalAccountSSHPrivilege struct {
	Enabled               *bool         `json:"principal_account_ssh"`
	AdminLevelPermissions *bool         `json:"admin_level_permissions"`
	SudoDisplayName       *string       `json:"sudo_display_name"`
	SudoCommandBundles    []NamedObject `json:"sudo_command_bundles"`
}

// Deprecated: Use security_policy2.go instead.
func (*PrincipalAccountSSHPrivilege) ValidForResourceType(resourceSelectorType ResourceSelectorType) bool {
	return resourceSelectorType == ServerBasedResourceSelectorType
}

// Deprecated: Use security_policy2.go instead.
func (p *PrincipalAccountSSHPrivilege) ToResourceMap() map[string]any {
	m := make(map[string]any)
	m[attributes.Enabled] = *p.Enabled
	if p.AdminLevelPermissions != nil {
		m[attributes.AdminLevelPermissions] = *p.AdminLevelPermissions
	} else {
		m[attributes.AdminLevelPermissions] = false
	}
	if len(p.SudoCommandBundles) > 0 {
		scbs := make([]any, len(p.SudoCommandBundles))
		for i, scb := range p.SudoCommandBundles {
			scbs[i] = *scb.Id
		}
		m[attributes.SudoCommandBundles] = scbs
		m[attributes.SudoDisplayName] = p.SudoDisplayName
	}
	return m
}

type PasswordCheckoutRDPPrivilege struct {
	Enabled *bool `json:"password_checkout_rdp"`
}

// Deprecated: Use security_policy2.go instead.
func (p *PasswordCheckoutRDPPrivilege) ToResourceMap() map[string]any {
	m := make(map[string]any, 1)
	m[attributes.Enabled] = *p.Enabled
	return m
}

// Deprecated: Use security_policy2.go instead.
func (*PasswordCheckoutRDPPrivilege) ValidForResourceType(resourceSelectorType ResourceSelectorType) bool {
	return resourceSelectorType == ServerBasedResourceSelectorType
}

type PasswordCheckoutSSHPrivilege struct {
	Enabled *bool `json:"password_checkout_ssh"`
}

// Deprecated: Use security_policy2.go instead.
func (*PasswordCheckoutSSHPrivilege) ValidForResourceType(resourceSelectorType ResourceSelectorType) bool {
	return resourceSelectorType == ServerBasedResourceSelectorType
}

// Deprecated: Use security_policy2.go instead.
func (p *PasswordCheckoutSSHPrivilege) ToResourceMap() map[string]any {
	m := make(map[string]any, 1)
	m[attributes.Enabled] = *p.Enabled
	return m
}

type SecretPrivilege struct {
	List         *bool `json:"list"`
	FolderCreate *bool `json:"folder_create"`
	FolderUpdate *bool `json:"folder_update"`
	FolderDelete *bool `json:"folder_delete"`
	SecretCreate *bool `json:"secret_create"`
	SecretUpdate *bool `json:"secret_update"`
	SecretReveal *bool `json:"secret_reveal"`
	SecretDelete *bool `json:"secret_delete"`
}

// Deprecated: Use security_policy2.go instead.
func (*SecretPrivilege) ValidForResourceType(resourceSelectorType ResourceSelectorType) bool {
	return resourceSelectorType == SecretBasedResourceSelectorType
}

// Deprecated: Use security_policy2.go instead.
func (p *SecretPrivilege) ToResourceMap() map[string]any {
	m := make(map[string]any, 8)

	m[attributes.List] = p.List != nil && *p.List
	m[attributes.FolderCreate] = p.FolderCreate != nil && *p.FolderCreate
	m[attributes.FolderUpdate] = p.FolderUpdate != nil && *p.FolderUpdate
	m[attributes.FolderDelete] = p.FolderDelete != nil && *p.FolderDelete
	m[attributes.SecretCreate] = p.SecretCreate != nil && *p.SecretCreate
	m[attributes.SecretUpdate] = p.SecretUpdate != nil && *p.SecretUpdate
	m[attributes.SecretReveal] = p.SecretReveal != nil && *p.SecretReveal
	m[attributes.SecretDelete] = p.SecretDelete != nil && *p.SecretDelete

	return m
}

type SecurityPolicyRule struct {
	ID               *string                                 `json:"id"`
	SecurityPolicyID *string                                 `json:"security_policy_id"`
	Name             *string                                 `json:"name"`
	ResourceType     ResourceSelectorType                    `json:"resource_type"`
	ResourceSelector SecurityPolicyRuleResourceSelector      `json:"resource_selector"`
	Privileges       []*SecurityPolicyRulePrivilegeContainer `json:"privileges"`
	Conditions       []*SecurityPolicyRuleConditionContainer `json:"conditions"`
}

// Deprecated: Use security_policy2.go instead.
func (r *SecurityPolicyRule) ToResourceMap() map[string]any {
	m := make(map[string]any, 7)

	if r.ID != nil {
		m[attributes.ID] = *r.ID
	}

	if r.Name != nil {
		m[attributes.Name] = *r.Name
	}

	resources := make([]any, 0, 1)
	if r.ResourceSelector != nil {
		resources = append(resources, r.ResourceSelector.ToResourceMap())
	}
	m[attributes.Resources] = resources

	if r.Privileges != nil {
		privilegesM := make(map[string]any, 4)
		passwordCheckoutRDP := make([]any, 0, 1)
		passwordCheckoutSSH := make([]any, 0, 1)
		principalAccountRDP := make([]any, 0, 1)
		principalAccountSSH := make([]any, 0, 1)
		secret := make([]any, 0, 1)

		for _, privilege := range r.Privileges {
			resourceMap := privilege.PrivilegeValue.ToResourceMap()
			switch privilege.PrivilegeType {
			case PasswordCheckoutRDPPrivilegeType:
				passwordCheckoutRDP = append(passwordCheckoutRDP, resourceMap)
			case PasswordCheckoutSSHPrivilegeType:
				passwordCheckoutSSH = append(passwordCheckoutSSH, resourceMap)
			case PrincipalAccountRDPPrivilegeType:
				principalAccountRDP = append(principalAccountRDP, resourceMap)
			case PrincipalAccountSSHPrivilegeType:
				principalAccountSSH = append(principalAccountSSH, resourceMap)
			case SecretPrivilegeType:
				secret = append(secret, resourceMap)
			}
		}

		privilegesM[attributes.PasswordCheckoutRDP] = passwordCheckoutRDP
		privilegesM[attributes.PasswordCheckoutSSH] = passwordCheckoutSSH
		privilegesM[attributes.PrincipalAccountRDP] = principalAccountRDP
		privilegesM[attributes.PrincipalAccountSSH] = principalAccountSSH
		privilegesM[attributes.Secret] = secret

		privileges := []any{privilegesM}
		m[attributes.Privileges] = privileges
	}

	conditions := make([]any, 0, 1)
	if len(r.Conditions) != 0 {
		conditionsM := make(map[string]any, 1)
		accessRequests := make([]any, 0, len(r.Conditions))
		gateways := make([]any, 0, len(r.Conditions))
		mfa := make([]any, 0, len(r.Conditions))

		for _, condition := range r.Conditions {
			switch condition.ConditionType {
			case AccessRequestConditionType:
				accessRequests = append(accessRequests, condition.ConditionValue.ToResourceMap())
			case GatewayConditionType:
				gateways = append(gateways, condition.ConditionValue.ToResourceMap())
			case MFAConditionType:
				mfa = append(mfa, condition.ConditionValue.ToResourceMap())
			}
		}

		conditionsM[attributes.AccessRequest] = accessRequests
		conditionsM[attributes.Gateway] = gateways
		conditionsM[attributes.MFA] = mfa
		conditions = append(conditions, conditionsM)
	}
	m[attributes.Conditions] = conditions

	return m
}

// Deprecated: Use security_policy2.go instead.
func (r *SecurityPolicyRule) UnmarshalJSON(data []byte) error {
	tmp := struct {
		ID               *string                                 `json:"id"`
		SecurityPolicyID *string                                 `json:"security_policy_id"`
		Name             *string                                 `json:"name"`
		ResourceType     ResourceSelectorType                    `json:"resource_type"`
		ResourceSelector json.RawMessage                         `json:"resource_selector"`
		Privileges       []*SecurityPolicyRulePrivilegeContainer `json:"privileges"`
		Conditions       []*SecurityPolicyRuleConditionContainer `json:"conditions"`
	}{}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	r.ID = tmp.ID
	r.SecurityPolicyID = tmp.SecurityPolicyID
	r.Name = tmp.Name
	r.ResourceType = tmp.ResourceType
	r.Privileges = tmp.Privileges
	r.Conditions = tmp.Conditions

	switch tmp.ResourceType {
	case ServerBasedResourceSelectorType:
		resourceSelector := &ServerBasedResourceSelector{}
		if err := json.Unmarshal(tmp.ResourceSelector, resourceSelector); err != nil {
			return err
		}
		r.ResourceSelector = resourceSelector
	case SecretBasedResourceSelectorType:
		resourceSelector := &SecretBasedResourceSelector{}
		if err := json.Unmarshal(tmp.ResourceSelector, resourceSelector); err != nil {
			return err
		}
		r.ResourceSelector = resourceSelector
	default:
		return fmt.Errorf("cannot unmarshal resource type: %s", tmp.ResourceType)
	}

	return nil
}

// Deprecated: Use security_policy2.go instead.
func (c OktaPAMClient) CreateSecurityPolicy(ctx context.Context, policy SecurityPolicy) (SecurityPolicy, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/security_policy", url.PathEscape(c.Team))
	logging.Tracef("making POST request to %s", requestURL)
	resultingPolicy := SecurityPolicy{}
	resp, err := c.CreateBaseRequest(ctx).SetBody(policy).SetResult(&resultingPolicy).Post(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return SecurityPolicy{}, err
	}

	if _, err = checkStatusCode(resp, http.StatusOK); err != nil {
		return SecurityPolicy{}, err
	}

	return resultingPolicy, nil
}

// Deprecated: Use security_policy2.go instead.
func (c OktaPAMClient) GetSecurityPolicy(ctx context.Context, securityPolicyID string) (*SecurityPolicy, error) {
	requestURL := fmt.Sprintf("/v1/teams/%s/security_policy/%s", url.PathEscape(c.Team), url.PathEscape(securityPolicyID))
	logging.Tracef("making GET request to %s", requestURL)
	policy := &SecurityPolicy{}
	resp, err := c.CreateBaseRequest(ctx).SetResult(&policy).Get(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return nil, err
	}

	statusCode := resp.StatusCode()
	if statusCode == http.StatusOK {
		return policy, nil
	} else if statusCode == http.StatusNotFound {
		return nil, nil
	}

	return nil, createErrorForInvalidCode(resp, http.StatusOK, http.StatusNotFound)
}

// Deprecated: Use security_policy2.go instead.
func (c OktaPAMClient) UpdateSecurityPolicy(ctx context.Context, securityPolicyID string, policy SecurityPolicy) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/security_policy/%s", url.PathEscape(c.Team), url.PathEscape(securityPolicyID))
	logging.Tracef("making PUt request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).SetBody(policy).Put(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	if _, err = checkStatusCode(resp, http.StatusNoContent); err != nil {
		return err
	}

	return nil
}

// Deprecated: Use security_policy2.go instead.
func (c OktaPAMClient) DeleteSecurityPolicy(ctx context.Context, securityPolicyID string) error {
	requestURL := fmt.Sprintf("/v1/teams/%s/security_policy/%s", url.PathEscape(c.Team), url.PathEscape(securityPolicyID))
	logging.Tracef("making DELETE request to %s", requestURL)
	resp, err := c.CreateBaseRequest(ctx).Delete(requestURL)
	if err != nil {
		logging.Errorf("received error while making request to %s", requestURL)
		return err
	}

	if _, err = checkStatusCode(resp, http.StatusOK, http.StatusNoContent, http.StatusNotFound); err != nil {
		return err
	}

	return nil
}
