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
	//ManagedSaasApp   *SecurityPolicyRuleManagedSaasAppBasedResourceSelectorModel   `tfsdk:"managed_saas_app"`
	//UnmanagedSaasApp *SecurityPolicyRuleUnmanagedSaasAppBasedResourceSelectorModel `tfsdk:"unmanaged_saas_app"`
	//OktaApp          *SecurityPolicyRuleOktaAppBasedResourceSelectorModel          `tfsdk:"okta_app"`
}

func SecurityPolicyRuleResourceSelectorBlock() schema.Block {
	return schema.SingleNestedBlock{
		Blocks: map[string]schema.Block{
			"servers": SecurityPolicyRuleServerBasedResourceSelectorBlock(),
			//"managed_saas_app":   SecurityPolicyRuleManagedSaasAppBasedResourceSelectorBlock(),
			//"unmanaged_saas_app": SecurityPolicyRuleUnmanagedSaasAppBasedResourceSelectorBlock(),
			//"okta_app":           SecurityPolicyRuleOktaAppBasedResourceSelectorBlock(),
		}}
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

func SecurityPolicyRuleServerBasedResourceSelectorBlock() schema.Block {
	return schema.SingleNestedBlock{
		Blocks: map[string]schema.Block{
			"individual_server":         SelectorIndividualServerBlock(),
			"individual_server_account": SelectorIndividualServerAccountBlock(),
			"server_label":              SelectorServerLabelBlock(),
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

func SelectorIndividualServerBlock() schema.Block {
	return schema.SingleNestedBlock{
		Attributes: map[string]schema.Attribute{
			"server": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func SelectorIndividualServerFromSDKToModel(ctx context.Context, in *pam.SelectorIndividualServer, out *SelectorIndividualServerModel) diag.Diagnostics {
	return NamedObjectFromSDKToModel(ctx, in.Server, &out.Server)
}

func SelectorIndividualServerFromModelToSDK(ctx context.Context, in *SelectorIndividualServerModel, out *pam.SelectorIndividualServer) diag.Diagnostics {
	return NamedObjectFromModelToSDK(ctx, in.Server, out.Server)
}

type SelectorIndividualServerAccountModel struct {
	ServerId NamedObjectModel `tfsdk:"server_id"`
	Username types.String     `tfsdk:"username"`
}

func SelectorIndividualServerAccountBlock() schema.Block {
	return schema.SingleNestedBlock{
		Attributes: map[string]schema.Attribute{
			"server_id": schema.StringAttribute{Required: true},
			"username":  schema.StringAttribute{Required: true},
		},
	}
}

func SelectorIndividualServerAccountFromModelToSDK(ctx context.Context, in *SelectorIndividualServerAccountModel, out *pam.SelectorIndividualServerAccount) diag.Diagnostics {
	out.Username = in.Username.ValueStringPointer()
	return NamedObjectFromModelToSDK(ctx, in.ServerId, out.ServerId)
}

type SelectorServerLabelModel struct {
	ServerSelector  *SelectorServerLabelServerSelectorModel `tfsdk:"server_selector"`
	AccountSelector types.List/*types.String*/ `tfsdk:"usernames"`
}

func SelectorServerLabelBlock() schema.Block {
	return schema.SingleNestedBlock{
		//TODO(ja)
	}
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
