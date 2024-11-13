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

func SecurityPolicyRuleResourceSelectorSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"servers": SecurityPolicyRuleServerBasedResourceSelectorSchema(), // server_based_resource
			//"managed_saas_app":   SecurityPolicyRuleManagedSaasAppBasedResourceSelectorSchema(),
			//"unmanaged_saas_app": SecurityPolicyRuleUnmanagedSaasAppBasedResourceSelectorSchema(),
			//"okta_app":           SecurityPolicyRuleOktaAppBasedResourceSelectorSchema(),
		},
		Required: true,
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
	SelectorServerLabel     *SelectorServerLabelModel             `tfsdk:"server_label"` // TODO(ja) - this is "server_label" in the spec, but "server_labels" in _v1
}

func SecurityPolicyRuleServerBasedResourceSelectorSchema() schema.Attribute {
	return schema.SingleNestedAttribute{ //TODO(ja) is array in spec
		Attributes: map[string]schema.Attribute{
			"individual_server":         SelectorIndividualServerSchema(),
			"individual_server_account": SelectorIndividualServerAccountSchema(),
			"server_label":              SelectorServerLabelSchema(), // TODO(ja) - this is "server_label" in the spec, but "server_labels" in _v1
		},
		Optional: true,
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
	Server types.String /*NamedObject*/ `tfsdk:"server"`
}

func SelectorIndividualServerSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"server": schema.StringAttribute{
				Required: true,
			},
		},
		Optional: true,
	}
}

func SelectorIndividualServerFromSDKToModel(ctx context.Context, in *pam.SelectorIndividualServer, out *SelectorIndividualServerModel) diag.Diagnostics {
	if val, ok := in.Server.GetIdOk(); ok {
		out.Server = types.StringPointerValue(val)
	}
	return nil
}

func SelectorIndividualServerFromModelToSDK(ctx context.Context, in *SelectorIndividualServerModel, out *pam.SelectorIndividualServer) diag.Diagnostics {
	if !in.Server.IsNull() {
		out.Server = pam.NewNamedObject().SetId(in.Server.ValueString())
	}
	return nil
}

type SelectorIndividualServerAccountModel struct {
	ServerId types.String/*NamedObject*/ `tfsdk:"server_id"`
	Username types.String `tfsdk:"username"`
}

func SelectorIndividualServerAccountSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"server_id": schema.StringAttribute{Required: true},
			"username":  schema.StringAttribute{Required: true},
		},
		Optional: true,
	}
}

func SelectorIndividualServerAccountFromModelToSDK(ctx context.Context, in *SelectorIndividualServerAccountModel, out *pam.SelectorIndividualServerAccount) diag.Diagnostics {
	out.Username = in.Username.ValueStringPointer()
	if !in.ServerId.IsNull() {
		out.ServerId = pam.NewNamedObject().SetId(in.ServerId.ValueString())
	}
	return nil
}

type SelectorServerLabelModel struct {
	ServerSelector  *SelectorServerLabelServerSelectorModel       `tfsdk:"server_selector"`
	AccountSelector types.List/*types.String*/ `tfsdk:"accounts"` // TODO(ja) "accounts" is not a list of accounts, but a list of usernames
}

func SelectorServerLabelSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"server_selector": SelectorServerLabelServerSelectorSchema(),
			"accounts":        schema.ListAttribute{ElementType: types.StringType, Optional: true},
		},
		Optional: true,
	}
}
func SelectorServerLabelFromModelToSDK(ctx context.Context, in *SelectorServerLabelModel, out *pam.SelectorServerLabel) diag.Diagnostics {
	// AccountSelector
	// ServerSelector
	return nil
}

type SelectorServerLabelServerSelectorModel struct {
	Labels types.Map `tfsdk:"labels"`
}

func SelectorServerLabelServerSelectorSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"labels": schema.MapAttribute{
				ElementType: types.StringType,
				Required:    true,
			},
		},
		Optional: true,
	}
}
