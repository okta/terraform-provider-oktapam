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
		var outSelector pam.SecurityPolicyRuleServerBasedResourceSelector
		if diags := SecurityPolicyRuleServerBasedResourceSelectorFromModelToSDK(ctx, in.Servers, &outSelector); diags.HasError() {
			return diags
		}
		out.SecurityPolicyRuleServerBasedResourceSelector = &outSelector
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
	out.Type = string(pam.SecurityPolicyRuleResourceType_SERVER_BASED_RESOURCE)
	if !in.Selectors.IsNull() && len(in.Selectors.Elements()) > 0 {
		selectorModels := make([]SecurityPolicyRuleServerBasedResourceSelectorContainerModel, 0, len(in.Selectors.Elements()))
		if diags := in.Selectors.ElementsAs(ctx, &selectorModels, false); diags.HasError() {
			return diags
		}
		for _, selectorModel := range selectorModels {
			var outSelector pam.SecurityPolicyRuleServerBasedResourceSelectorContainer
			if diags := SecurityPolicyRuleServerBasedResourceSelectorContainerFromModelToSDK(ctx, &selectorModel, &outSelector); diags.HasError() {
				return diags
			}
			out.Selectors = append(out.Selectors, outSelector)
		}
	}
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
		var outVal pam.SelectorIndividualServer
		SelectorIndividualServerFromModelToSDK(ctx, in.IndividualServer, &outVal)
		out.Selector.SelectorIndividualServer = &outVal
	} else if in.IndividualServerAccount != nil {
		selectorType = pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_INDIVIDUAL_SERVER_ACCOUNT
		var outVal pam.SelectorIndividualServerAccount
		SelectorIndividualServerAccountFromModelToSDK(ctx, in.IndividualServerAccount, &outVal)
		out.Selector.SelectorIndividualServerAccount = &outVal
	} else if in.SelectorServerLabel != nil {
		selectorType = pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_SERVER_LABEL
		var outVal pam.SelectorServerLabel
		SelectorServerLabelFromModelToSDK(ctx, in.SelectorServerLabel, &outVal)
		out.Selector.SelectorServerLabel = &outVal
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

func SelectorIndividualServerFromSDKToModel(_ context.Context, in *pam.SelectorIndividualServer, out *SelectorIndividualServerModel) diag.Diagnostics {
	if val, ok := in.Server.GetIdOk(); ok {
		out.Server = types.StringPointerValue(val)
	}
	return nil
}

func SelectorIndividualServerFromModelToSDK(_ context.Context, in *SelectorIndividualServerModel, out *pam.SelectorIndividualServer) diag.Diagnostics {
	out.Type = string(pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_INDIVIDUAL_SERVER)
	if !in.Server.IsNull() {
		out.Server = *pam.NewNamedObject().SetId(in.Server.ValueString())
	}
	return nil
}

type SelectorIndividualServerAccountModel struct {
	Server   types.String/*NamedObject*/ `tfsdk:"server"`
	Username types.String `tfsdk:"username"`
}

func SelectorIndividualServerAccountSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"server":   schema.StringAttribute{Required: true},
			"username": schema.StringAttribute{Required: true},
		},
		Optional: true,
	}
}

func SelectorIndividualServerAccountFromModelToSDK(_ context.Context, in *SelectorIndividualServerAccountModel, out *pam.SelectorIndividualServerAccount) diag.Diagnostics {
	out.Type = string(pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_INDIVIDUAL_SERVER_ACCOUNT) //TODO(ja) this should probably be hard coded
	out.Username = in.Username.ValueStringPointer()
	if !in.Server.IsNull() && !in.Server.IsUnknown() {
		out.Server = *pam.NewNamedObject().SetId(in.Server.ValueString())
	}
	out.Username = in.Username.ValueStringPointer()
	return nil
}

type SelectorServerLabelModel struct {
	ServerSelector  *SelectorServerLabelServerSelectorModel `tfsdk:"server_selector"`
	AccountSelector types.List/*types.String*/ `tfsdk:"account_selector"`
}

func SelectorServerLabelSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"server_selector":  SelectorServerLabelServerSelectorSchema(),
			"account_selector": schema.ListAttribute{ElementType: types.StringType, Optional: true},
		},
		Optional: true,
	}
}
func SelectorServerLabelFromModelToSDK(ctx context.Context, in *SelectorServerLabelModel, out *pam.SelectorServerLabel) diag.Diagnostics {

	out.Type = string(pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_SERVER_LABEL) //TODO(ja) this should probably be a hard-coded string
	if diags := SelectorServerLabelServerSelectorFromModelToSDK(ctx, in.ServerSelector, &out.ServerSelector); diags.HasError() {
		return diags
	}

	//TODO(ja) I think this is supposed to be its own separate ModelToSDK()

	if !in.AccountSelector.IsNull() && len(in.AccountSelector.Elements()) > 0 {
		accountSelectorModel := make([]types.String, 0, len(in.AccountSelector.Elements()))
		if diags := in.AccountSelector.ElementsAs(ctx, &accountSelectorModel, false); diags.HasError() {
			return diags
		}
		var outAccountSelector pam.SelectorServerLabelAccountSelector
		for _, elem := range accountSelectorModel {
			if !elem.IsNull() && !elem.IsUnknown() {
				outAccountSelector.Usernames = append(outAccountSelector.Usernames, elem.ValueString())
			}
		}
		out.AccountSelector = &outAccountSelector
	}
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

func SelectorServerLabelServerSelectorFromModelToSDK(ctx context.Context, in *SelectorServerLabelServerSelectorModel, out *pam.SelectorServerLabelServerSelector) diag.Diagnostics {
	elements := make(map[string]types.String, len(in.Labels.Elements()))
	if diags := in.Labels.ElementsAs(ctx, &elements, false); diags.HasError() {
		return diags
	}
	outMap := make(map[string]any, len(elements))
	for k, v := range elements {
		outMap[k] = v.ValueString()
	}
	out.Labels = outMap
	return nil
}

func SelectorServerLabelServerSelectorFromSDKToModel(ctx context.Context, in *pam.SelectorServerLabelServerSelector, out *SelectorServerLabelServerSelectorModel) diag.Diagnostics {
	return nil
}
