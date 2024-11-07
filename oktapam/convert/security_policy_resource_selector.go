package convert

import (
	"context"
	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecurityPolicyRuleResourceSelectorModel struct {
	Servers *SecurityPolicyRuleServerBasedResourceSelectorModel `tfsdk:"servers"`
	//ManagedSaasApp   *SecurityPolicyRuleManagedSaasAppBasedResourceSelector   `tfsdk:"managed_saas_app"`
	//UnmanagedSaasApp *SecurityPolicyRuleUnmanagedSaasAppBasedResourceSelector `tfsdk:"unmanaged_saas_app"`
	//OktaApp          *SecurityPolicyRuleOktaAppBasedResourceSelector          `tfsdk:"okta_app"`
}

func SecurityPolicyRuleResourceSelectorSchema() schema.Attribute {
	var s schema.Schema

	s.
	return schema.NestedAttributeObject{}
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"servers": SecurityPolicyRuleServerBasedResourceSelectorSchema(),
			// "managed_saas_app":
			// "unmanaged_saas_app":
			// "okta_app":
		},
	}
}

func SecurityPolicyRuleResourceSelectorFromModelToSDK(ctx context.Context, in *SecurityPolicyRuleResourceSelectorModel, out *pam.SecurityPolicyRuleResourceSelectorContainer) diag.Diagnostics {
	if in.Servers != nil {
		SecurityPolicyRuleServerBasedResourceSelectorFromModelToSDK(ctx, in.Servers, out)
	}
	return nil
}

type SecurityPolicyRuleServerBasedResourceSelectorModel struct {
	IndividualServer        *SelectorIndividualServerModel        `tfsdk:"individual_server"`
	IndividualServerAccount *SelectorIndividualServerAccountModel `tfsdk:"individual_server_account"`
	SelectorServerLabel     *SelectorServerLabelModel             `tfsdk:"server_label"`
}

func SecurityPolicyRuleServerBasedResourceSelectorSchema() schema.Attribute {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"individual_server":         SelectorIndividualServerSchema(),
			"individual_server_account": SelectorIndividualServerAccountSchema(),
			"server_label":              SelectorServerLabelSchema(),,
		},
	}
}

func SecurityPolicyRuleServerBasedResourceSelectorFromModelToSDK(ctx context.Context, in *SecurityPolicyRuleServerBasedResourceSelectorModel, out *pam.SecurityPolicyRuleResourceSelectorContainer) diag.Diagnostics {
	out.Selector = &pam.SecurityPolicyRuleResourceSelector{}

	var selectorType pam.SecurityPolicyRuleResourceSelectorType

	if in.IndividualServer != nil {
		selectorType = pam.SecurityPolicyRuleResourceSelectorType_INDIVIDUAL_SERVER
		SelectorIndividualServerFromModelToSDK(ctx, in.IndividualServer, out.Selector.SelectorIndividualServer)
	} else if in.IndividualServerAccount != nil {
		selectorType = pam.SecurityPolicyRuleResourceSelectorType_INDIVIDUAL_SERVER_ACCOUNT
		SelectorIndividualServerAccountFromModelToSDK(ctx, in.IndividualServerAccount, out.Selector.SelectorIndividualServerAccount)
	} else if in.SelectorServerLabel != nil {
		selectorType = pam.SecurityPolicyRuleResourceSelectorType_SERVER_LABEL
		SelectorServerLabelFromModelToSDK(ctx, in.SelectorServerLabel, out.Selector.SelectorServerLabel)
	} else {
		//TODO(ja) unhandled case
	}
	out.SelectorType = &selectorType
	return nil
}

type SelectorIndividualServerModel struct {
	Server NamedObjectModel `tfsdk:"server"`
}

func SelectorIndividualServerSchema() schema.Attribute {
	return schema.Attribute{
		Attributes: map[string]schema.Attribute{
			"server": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func SelectorIndividualServerFromModelToSDK(ctx context.Context, in *SelectorIndividualServerModel, out *pam.SelectorIndividualServer) diag.Diagnostics {
	out.Server = in.Server
	return nil
}

type SelectorIndividualServerAccountModel struct {
	ServerId NamedObjectModel `tfsdk:"server_id"`
	Username types.String     `tfsdk:"username"`
}

func SelectorIndividualServerAccountFromModelToSDK(ctx context.Context, in *SelectorIndividualServerAccountModel, out *pam.SelectorIndividualServerAccount) diag.Diagnostics {
	out.Username = in.Username.ValueStringPointer()
	out.ServerId = in.ServerId
	return nil
}

type SelectorServerLabelModel struct {
	ServerSelector  *SelectorServerLabelServerSelectorModel `tfsdk:"server_selector"`
	AccountSelector types.List /*types.String*/             `tfsdk:"usernames"`
}

func SelectorServerLabelFromModelToSDK(ctx context.Context, in *SelectorServerLabelModel, out *pam.SelectorServerLabel) diag.Diagnostics {
	// AccountSelector
	// ServerSelector
	return nil
}

type SelectorServerLabelServerSelectorModel struct {
	Labels types.MapType `tfsdk:"labels"`
}

type SelectorServerLabelAccountSelectorModel struct {
	Usernames []types.String `tfsdk:"usernames"`
}
