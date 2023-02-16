package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"github.com/okta/terraform-provider-oktapam/oktapam/logging"
)

type PrivilegeType string
type ServerBasedResourceSubSelectorType string
type ResourceSelectorType string
type AccountSelectorType string
type ConditionType string

const (
	PrincipalAccountSSHPrivilegeType = PrivilegeType("principal_account_ssh")
	PrincipalAccountRDPPrivilegeType = PrivilegeType("principal_account_rdp")
	PasswordCheckoutSSHPrivilegeType = PrivilegeType("password_checkout_ssh")
	PasswordCheckoutRDPPrivilegeType = PrivilegeType("password_checkout_rdp")

	ServerBasedResourceSelectorType = ResourceSelectorType("server_based_resource")

	IndividualServerSubSelectorType        = ServerBasedResourceSubSelectorType("individual_server")
	IndividualServerAccountSubSelectorType = ServerBasedResourceSubSelectorType("individual_server_account")
	ServerLabelServerSubSelectorType       = ServerBasedResourceSubSelectorType("server_label")

	UsernameAccountSelectorType = AccountSelectorType("username")
	NoneAccountSelectorType     = AccountSelectorType("none")

	AccessRequestConditionType = ConditionType("access_request")
)

type SecurityPolicy struct {
	ID          *string                   `json:"id,omitempty"`
	Name        *string                   `json:"name"`
	Description *string                   `json:"description,omitempty"`
	Active      *bool                     `json:"active"`
	Principals  *SecurityPolicyPrincipals `json:"principals"`
	Rules       []*SecurityPolicyRule     `json:"rules"`
}

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

		m[attributes.Rule] = rules
	}

	return m
}

type SecurityPolicyPrincipals struct {
	UserGroups []NamedObject `json:"user_groups"`
	Users      []NamedObject `json:"users"`
}

func (p *SecurityPolicyPrincipals) ToResourceMap() map[string]any {
	m := make(map[string]any)

	userIds := make([]any, len(p.Users))
	for idx, userId := range p.Users {
		u := userId
		userIds[idx] = *u.Id
	}
	m[attributes.Users] = userIds

	groupIds := make([]any, len(p.UserGroups))
	for idx, groupId := range p.UserGroups {
		g := groupId
		groupIds[idx] = *g.Id
	}
	m[attributes.Groups] = groupIds

	return m
}

type SecurityPolicyRuleConditionContainer struct {
	ConditionType  ConditionType               `json:"condition_type"`
	ConditionValue SecurityPolicyRuleCondition `json:"condition_value"`
}

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
	default:
		return fmt.Errorf("received unknown condition type: %s", tmp.ConditionType)
	}

	if err := json.Unmarshal(tmp.ConditionValue, c.ConditionValue); err != nil {
		return err
	}

	return nil
}

type SecurityPolicyRuleCondition interface {
	ConditionType() ConditionType
	ToResourceMap() map[string]any
}

type AccessRequestCondition struct {
	RequestTypeID   *string `json:"request_type_id"`
	RequestTypeName *string `json:"request_type_name"`
}

func (c *AccessRequestCondition) ToResourceMap() map[string]any {
	m := make(map[string]any)

	m[attributes.RequestTypeId] = *c.RequestTypeID
	m[attributes.RequestTypeName] = *c.RequestTypeName

	return m
}

func (*AccessRequestCondition) ConditionType() ConditionType {
	return AccessRequestConditionType
}

type SecurityPolicyRuleResourceSelector interface {
	ResourceSelectorType() ResourceSelectorType
	ToResourceMap() map[string]any
}

type ServerBasedResourceSelector struct {
	Selectors []ServerBasedResourceSubSelectorContainer `json:"selectors"`
}

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

type ServerBasedResourceSubSelector interface {
	ServerBasedResourceSubSelectorType() ServerBasedResourceSubSelectorType
	ToResourceMap() map[string]any
}

type IndividualServerSubSelector struct {
	ServerId NamedObject `json:"server_id"`
}

func (*IndividualServerSubSelector) ServerBasedResourceSubSelectorType() ServerBasedResourceSubSelectorType {
	return IndividualServerSubSelectorType
}

func (s *IndividualServerSubSelector) ToResourceMap() map[string]any {
	m := make(map[string]any)
	m[attributes.ServerId] = *s.ServerId.Id
	return m
}

type IndividualServerAccountSubSelector struct {
	ServerId NamedObject `json:"server_id"`
	Username *string     `json:"username"`
}

func (*IndividualServerAccountSubSelector) ServerBasedResourceSubSelectorType() ServerBasedResourceSubSelectorType {
	return IndividualServerAccountSubSelectorType
}

func (s *IndividualServerAccountSubSelector) ToResourceMap() map[string]any {
	m := make(map[string]any)
	m[attributes.ServerId] = *s.ServerId.Id
	m[attributes.Username] = *s.Username
	return m
}

type ServerLabelBasedSubSelector struct {
	ServerSelector *ServerLabelServerSelector `json:"server_selector"`

	AccountSelectorType AccountSelectorType `json:"account_selector_type"`
	AccountSelector     AccountSelector     `json:"account_selector"`
}

func (*ServerLabelBasedSubSelector) ServerBasedResourceSubSelectorType() ServerBasedResourceSubSelectorType {
	return ServerLabelServerSubSelectorType
}

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

	serverLabelsM[attributes.Usernames] = usernamesArr

	return m
}

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

type ServerLabelServerSelector struct {
	Labels map[string]string `json:"labels"`
}

type AccountSelector interface {
	AccountSelectorType() AccountSelectorType
}

type NoneAccountSelector struct{}

func (*NoneAccountSelector) AccountSelectorType() AccountSelectorType {
	return NoneAccountSelectorType
}

type UsernameAccountSelector struct {
	Usernames []string `json:"usernames"`
}

func (*UsernameAccountSelector) AccountSelectorType() AccountSelectorType {
	return UsernameAccountSelectorType
}

type ServerBasedResourceSubSelectorContainer struct {
	SelectorType ServerBasedResourceSubSelectorType `json:"selector_type"`
	Selector     ServerBasedResourceSubSelector     `json:"selector"`
}

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

func (*ServerBasedResourceSelector) ResourceSelectorType() ResourceSelectorType {
	return ServerBasedResourceSelectorType
}

type SecurityPolicyRulePrivilegeContainer struct {
	PrivilegeType  PrivilegeType               `json:"privilege_type"`
	PrivilegeValue SecurityPolicyRulePrivilege `json:"privilege_value"`
}

func (c *SecurityPolicyRulePrivilegeContainer) ToResourceMap() map[string]any {
	m := make(map[string]any, 4)

	return m
}

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
	default:
		return fmt.Errorf("received unknown privilege type: %s", tmp.PrivilegeValue)
	}

	if err := json.Unmarshal(tmp.PrivilegeValue, c.PrivilegeValue); err != nil {
		return err
	}
	return nil
}

type SecurityPolicyRulePrivilege interface {
	isPrivilege()
	ToResourceMap() map[string]any
}

type PrincipalAccountRDPPrivilege struct {
	Enabled *bool `json:"principal_account_rdp"`
}

func (*PrincipalAccountRDPPrivilege) isPrivilege() {}

func (p *PrincipalAccountRDPPrivilege) ToResourceMap() map[string]any {
	m := make(map[string]any, 1)
	m[attributes.Enabled] = *p.Enabled
	return m
}

type PrincipalAccountSSHPrivilege struct {
	Enabled *bool `json:"principal_account_ssh"`
}

func (*PrincipalAccountSSHPrivilege) isPrivilege() {}

func (p *PrincipalAccountSSHPrivilege) ToResourceMap() map[string]any {
	m := make(map[string]any, 1)
	m[attributes.Enabled] = *p.Enabled
	return m
}

type PasswordCheckoutRDPPrivilege struct {
	Enabled *bool `json:"password_checkout_rdp"`
}

func (p *PasswordCheckoutRDPPrivilege) ToResourceMap() map[string]any {
	m := make(map[string]any, 1)
	m[attributes.Enabled] = *p.Enabled
	return m
}

func (*PasswordCheckoutRDPPrivilege) isPrivilege() {}

type PasswordCheckoutSSHPrivilege struct {
	Enabled *bool `json:"password_checkout_ssh"`
}

func (*PasswordCheckoutSSHPrivilege) isPrivilege() {}

func (p *PasswordCheckoutSSHPrivilege) ToResourceMap() map[string]any {
	m := make(map[string]any, 1)
	m[attributes.Enabled] = *p.Enabled
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

	privileges := make([]any, len(r.Privileges))
	if r.Privileges != nil {
		privilegesM := make(map[string]any, 4)
		passwordCheckoutRDP := make([]any, 0, 1)
		passwordCheckoutSSH := make([]any, 0, 1)
		principalAccountRDP := make([]any, 0, 1)
		principalAccountSSH := make([]any, 0, 1)

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
			}
		}

		privilegesM[attributes.PasswordCheckoutRDP] = passwordCheckoutRDP
		privilegesM[attributes.PasswordCheckoutSSH] = passwordCheckoutSSH
		privilegesM[attributes.PrincipalAccountRDP] = principalAccountRDP
		privilegesM[attributes.PrincipalAccountSSH] = principalAccountSSH
	}
	m[attributes.Privileges] = privileges

	conditions := make([]any, 0, len(r.Conditions))
	if r.Conditions != nil {
		conditionsM := make(map[string]any, 1)
		accessRequests := make([]any, 0, len(r.Conditions))

		for _, condition := range r.Conditions {
			switch condition.ConditionType {
			case AccessRequestConditionType:
				accessRequests = append(accessRequests, condition.ConditionValue.ToResourceMap())
			}
		}

		conditionsM[attributes.AccessRequest] = accessRequests
		conditions = append(conditions, conditionsM)
	}
	m[attributes.Conditions] = conditions

	return m
}

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
	default:
		return fmt.Errorf("cannot unmarshal resource type: %s", tmp.ResourceType)
	}

	return nil
}

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
