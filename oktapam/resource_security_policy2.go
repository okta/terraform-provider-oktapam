package oktapam

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

type NamedObjectTypeModel types.String
type SecurityPolicyTypeModel types.String
type SecurityPolicyRuleResourceTypeModel types.String
type SecurityPolicyRulePrivilegeTypeModel types.String
type SecurityPolicyRuleConditionTypeModel types.String
type SecurityPolicyRuleResourceSelectorTypeModel types.String
type ConditionsMFAACRValuesModel types.String

var _ resource.Resource = &SecurityPolicyResource{}

type SecurityPolicyResource struct {
	client *client.SDKClientWrapper
}

func NewSecurityPolicyResource() resource.Resource {
	return &SecurityPolicyResource{}
}

func (s *SecurityPolicyResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_security_policy_v2"
}

func (s *SecurityPolicyResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityPolicyResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityPolicyResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityPolicyResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityPolicyResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	//TODO implement me
	panic("implement me")
}

func (s *SecurityPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	sdkClient := getSDKClientFromMetadata(req.ProviderData)
	s.client = &sdkClient
}

type SecurityPolicyModel struct {
	Id          types.String                  `tfsdk:"id"`
	Name        types.String                  `tfsdk:"name"`
	Type        SecurityPolicyTypeModel       `tfsdk:"type"`
	Description types.String                  `tfsdk:"description"`
	Active      types.Bool                    `tfsdk:"active"`
	Principals  SecurityPolicyPrincipalsModel `tfsdk:"principals"`
	Rules       []SecurityPolicyRuleModel     `tfsdk:"rules"`
}

type SecurityPolicyPrincipalsModel struct {
	UserGroups []NamedObjectModel `tfsdk:"user_groups"`
}

type NamedObjectModel struct {
	Id      types.String         `tfsdk:"id"`
	Name    types.String         `tfsdk:"name"`
	Type    NamedObjectTypeModel `tfsdk:"type"`
	Missing types.Bool           `tfsdk:"missing"`
}

type SecurityPolicyRuleModel struct {
	Name                              types.String                                `tfsdk:"name"`
	ResourceType                      SecurityPolicyRuleResourceTypeModel         `tfsdk:"resource_type"`
	ResourceSelector                  SecurityPolicyRuleResourceSelectorsModel    `tfsdk:"resource_selector"`
	Privileges                        []SecurityPolicyRulePrivilegeContainerModel `tfsdk:"privileges"`
	Conditions                        []SecurityPolicyRuleConditionContainerModel `tfsdk:"conditions"`
	OverrideCheckoutDurationInSeconds types.Int64                                 `tfsdk:"override_checkout_duration_in_seconds"`
	SecurityPolicyID                  types.String                                `tfsdk:"security_policy_id"`
}

type SecurityPolicyRuleResourceSelectorsModel struct {
	Selectors []SecurityPolicyRuleResourceSelectorContainerModel `tfsdk:"selectors"`
}

type SecurityPolicyRulePrivilegeContainerModel struct {
	PrivilegeType  SecurityPolicyRulePrivilegeTypeModel                    `tfsdk:"privilege_type"`
	PrivilegeValue SecurityPolicyRulePrivilegeContainerPrivilegeValueModel `tfsdk:"privilege_value"`
}

type SecurityPolicyRuleConditionContainerModel struct {
	ConditionType  SecurityPolicyRuleConditionTypeModel `tfsdk:"condition_type"`
	ConditionValue SecurityPolicyRuleConditionModel     `tfsdk:"condition_value"`
}

type SecurityPolicyRuleResourceSelectorContainerModel struct {
	SelectorType SecurityPolicyRuleResourceSelectorTypeModel `tfsdk:"selector_type"`
	Selector     SecurityPolicyRuleResourceSelectorModel     `tfsdk:"selector"`
}

type SecurityPolicyRulePrivilegeContainerPrivilegeValueModel struct {
	SecurityPolicyPasswordCheckoutDatabasePrivilege *SecurityPolicyPasswordCheckoutDatabasePrivilegeModel `tfsdk:"password_checkout_database"`
	//SecurityPolicyPasswordCheckoutRDPPrivilege      *SecurityPolicyPasswordCheckoutRDPPrivilegeModel      `tfsdk:"password_checkout_rdp"`
	//SecurityPolicyPasswordCheckoutSSHPrivilege *SecurityPolicyPasswordCheckoutSSHPrivilegeModel `tfsdk:"password_checkout_ssh"`
	//SecurityPolicyPrincipalAccountRDPPrivilege      *SecurityPolicyPrincipalAccountRDPPrivilegeModel      `tfsdk:"principal_account_rdp"`
	SecurityPolicyPrincipalAccountSSHPrivilege *SecurityPolicyPrincipalAccountSSHPrivilegeModel `tfsdk:"principal_account_ssh"`
	//SecurityPolicyRevealPasswordPrivilege           *SecurityPolicyRevealPasswordPrivilegeModel           `tfsdk:"reveal_password"`
	//SecurityPolicySecretPrivilege                   *SecurityPolicySecretPrivilegeModel                   `tfsdk:"secret"`
	//SecurityPolicyUpdatePasswordPrivilege           *SecurityPolicyUpdatePasswordPrivilegeModel           `tfsdk:"update_password"`
}

type SecurityPolicyRuleConditionModel struct {
	ConditionsAccessRequests *ConditionsAccessRequestsModel
	ConditionsGateway        *ConditionsGatewayModel
	ConditionsMFA            *ConditionsMFAModel
}

type SecurityPolicyRuleResourceSelectorModel struct {
	SelectorIndividualServer        *SelectorIndividualServerModel
	SelectorIndividualServerAccount *SelectorIndividualServerAccountModel
	SelectorServerLabel             *SelectorServerLabelModel
}

type SecurityPolicyPrivilegeModel struct {
	Type SecurityPolicyRulePrivilegeTypeModel `tfsdk:"_type"`
}

type SecurityPolicyPasswordCheckoutDatabasePrivilegeModel struct {
	SecurityPolicyPrivilegeModel
	PasswordCheckoutDatabase types.Bool `tfsdk:"password_checkout_database"`
}

type SecurityPolicyPrincipalAccountSSHPrivilegeModel struct {
	SecurityPolicyPrivilegeModel
	PrincipalAccountSSH   types.Bool         `tfsdk:"principal_account_ssh"`
	AdminLevelPermissions types.Bool         `tfsdk:"admin_level_permissions"`
	SudoDisplayName       types.String       `tfsdk:"sudo_display_name"`
	SudoCommandBundles    []NamedObjectModel `tfsdk:"sudo_command_bundles"`
}

type ConditionsAccessRequestsModel struct {
	RequestTypeId       types.String `tfsdk:"request_type_id"`
	RequestTypeName     types.String `tfsdk:"request_type_name"`
	ExpiresAfterSeconds types.Int32  `tfsdk:"expires_after_seconds"`
}

type ConditionsGatewayModel struct {
	TrafficForwarding types.Bool `tfsdk:"traffic_forwarding"`
	SessionRecording  types.Bool `tfsdk:"session_recording"`
}

type ConditionsMFAModel struct {
	ReAuthFrequencyInSeconds types.Int32                 `tfsdk:"re_auth_frequency_in_seconds"`
	AcrValues                ConditionsMFAACRValuesModel `tfsdk:"acr_values"`
}

type SelectorIndividualServerModel struct {
	Server NamedObjectModel `tfsdk:"server"`
}

type SelectorIndividualServerAccountModel struct {
	ServerId NamedObjectModel `tfsdk:"server_id"`
	Username types.String     `tfsdk:"username"`
}

type SelectorServerLabelAccountSelectorTypeModel types.String
type SelectorServerLabelModel struct {
	ServerSelector      *SelectorServerLabelServerSelectorModel     `tfsdk:"server_selector"`
	AccountSelectorType SelectorServerLabelAccountSelectorTypeModel `tfsdk:"account_selector_type"`
	AccountSelector     SelectorServerLabelAccountSelectorModel     `tfsdk:"account_selector"`
}

type SelectorServerLabelServerSelectorModel struct {
	Labels types.MapType `tfsdk:"labels"`
}

type SelectorServerLabelAccountSelectorModel struct {
	Usernames []types.String `tfsdk:"usernames"`
}
