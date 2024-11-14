package convert

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecurityPolicyRuleResourceSelectorModel struct {
	Servers *SecurityPolicyRuleServerBasedResourceSelectorModel `tfsdk:"server_based_resource"`
	//ManagedSaasApp   *SecurityPolicyRuleManagedSaasAppBasedResourceSelectorModel   `tfsdk:"managed_saas_app_based_resource"`
	//UnmanagedSaasApp *SecurityPolicyRuleUnmanagedSaasAppBasedResourceSelectorModel `tfsdk:"unmanaged_saas_app_based_resource"`
	//OktaApp          *SecurityPolicyRuleOktaAppBasedResourceSelectorModel          `tfsdk:"okta_app_based_resource"`
	//Secret *SecurityPolicyRuleSecretBasedResourceSelectorModel `tfsdk:"secret_based_resource"`
	//ActiveDirectory *SecurityPolicyRuleActiveDirectoryBasedResourceSelectorModel `tfsdk:"active_directory_based_resource"`
}

func SecurityPolicyRuleResourceSelectorSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"server_based_resource": SecurityPolicyRuleServerBasedResourceSelectorSchema(), // server_based_resource
			//"managed_saas_app":   SecurityPolicyRuleManagedSaasAppBasedResourceSelectorSchema(),
			//"unmanaged_saas_app": SecurityPolicyRuleUnmanagedSaasAppBasedResourceSelectorSchema(),
			//"okta_app":           SecurityPolicyRuleOktaAppBasedResourceSelectorSchema(),

		},
		Required: true,
	}
}

func SecurityPolicyRuleResourceSelectorFromModelToSDK(ctx context.Context, in *SecurityPolicyRuleResourceSelectorModel, out *pam.SecurityPolicyRuleResourceSelector) diag.Diagnostics {
	if in.Servers != nil {
		if diags := SecurityPolicyRuleServerBasedResourceSelectorFromModelToSDK(ctx, in.Servers, out.SecurityPolicyRuleServerBasedResourceSelector); diags.HasError() {
			return diags
		}
	}
	return nil
}

func SecurityPolicyRuleResourceSelectorFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRuleResourceSelector, out *SecurityPolicyRuleResourceSelectorModel) diag.Diagnostics {
	if in.SecurityPolicyRuleServerBasedResourceSelector != nil {
		if diags := SecurityPolicyRuleServerBasedResourceSelectorFromSDKToModel(ctx, in.SecurityPolicyRuleServerBasedResourceSelector, out.Servers); diags.HasError() {
			return diags
		}
	}
	return nil
}

type SecurityPolicyRuleServerBasedResourceSelectorModel struct {
	Selectors types.List /*[]SecurityPolicyRuleServerBasedResourceSelectorContainerModel*/ `tfsdk:"selectors"`
}

func SecurityPolicyRuleServerBasedResourceSelectorSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"selectors": schema.ListNestedAttribute{
				NestedObject: SecurityPolicyRuleServerBasedResourceSelectorContainerSchema(),
				Required:     true,
			},
		},
		Required: true,
	}
}
func SecurityPolicyRuleServerBasedResourceSelectorFromModelToSDK(ctx context.Context, in *SecurityPolicyRuleServerBasedResourceSelectorModel, out *pam.SecurityPolicyRuleServerBasedResourceSelector) diag.Diagnostics {
	//TODO(ja)
	return nil
}

func SecurityPolicyRuleServerBasedResourceSelectorFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRuleServerBasedResourceSelector, out *SecurityPolicyRuleServerBasedResourceSelectorModel) diag.Diagnostics {
	//TODO(ja)
	return nil
}

type SecurityPolicyRuleServerBasedResourceSelectorContainerModel struct {
	IndividualServer        *SelectorIndividualServerModel        `tfsdk:"individual_server"`
	IndividualServerAccount *SelectorIndividualServerAccountModel `tfsdk:"individual_server_account"`
	SelectorServerLabel     *SelectorServerLabelModel             `tfsdk:"server_label"`
}

func SecurityPolicyRuleServerBasedResourceSelectorContainerSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"individual_server":         SelectorIndividualServerSchema(),
			"individual_server_account": SelectorIndividualServerAccountSchema(),
			"server_label":              SelectorServerLabelSchema(),
		},
	}
}
func SecurityPolicyRuleServerBasedResourceSelectorContainerFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRuleServerBasedResourceSelectorContainer, out *SecurityPolicyRuleServerBasedResourceSelectorContainerModel) diag.Diagnostics {
	//TODO(ja)
	return nil
}

func SecurityPolicyRuleServerBasedResourceSelectorContainerFromModelToSDK(ctx context.Context, in *SecurityPolicyRuleServerBasedResourceSelectorContainerModel, out *pam.SecurityPolicyRuleServerBasedResourceSelectorContainer) diag.Diagnostics {
	var selectorType pam.SecurityPolicyRuleServerBasedResourceSubSelectorType

	if in.IndividualServer != nil {
		selectorType = pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_INDIVIDUAL_SERVER
		SelectorIndividualServerFromModelToSDK(ctx, in.IndividualServer, out.Selector.SelectorIndividualServer)
	} else if in.IndividualServerAccount != nil {
		selectorType = pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_INDIVIDUAL_SERVER_ACCOUNT
		SelectorIndividualServerAccountFromModelToSDK(ctx, in.IndividualServerAccount, out.Selector.SelectorIndividualServerAccount)
	} else if in.SelectorServerLabel != nil {
		selectorType = pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_SERVER_LABEL
		SelectorServerLabelFromModelToSDK(ctx, in.SelectorServerLabel, out.Selector.SelectorServerLabel)
	} else {
		//TODO(ja) unhandled case
	}
	out.SelectorType = selectorType
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
		out.Server = *pam.NewNamedObject().SetId(in.Server.ValueString())
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

func SelectorIndividualServerAccountFromModelToSDK(_ context.Context, in *SelectorIndividualServerAccountModel, out *pam.SelectorIndividualServerAccount) diag.Diagnostics {
	out.Username = in.Username.ValueStringPointer()
	if !in.ServerId.IsNull() {
		out.Server = *pam.NewNamedObject().SetId(in.ServerId.ValueString())
	}
	out.Username = in.Username.ValueStringPointer()
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
