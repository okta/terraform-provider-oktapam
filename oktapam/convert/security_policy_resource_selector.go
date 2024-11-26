package convert

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecurityPolicyRuleResourceSelectorModel struct {
	ServerBasedResource *ServerBasedResourceSelectorModel `tfsdk:"server_based_resource"`
	//ManagedSaasApp   *SecurityPolicyRuleManagedSaasAppBasedResourceSelectorModel   `tfsdk:"managed_saas_app_based_resource"`
	//UnmanagedSaasApp *SecurityPolicyRuleUnmanagedSaasAppBasedResourceSelectorModel `tfsdk:"unmanaged_saas_app_based_resource"`
	//OktaApp          *SecurityPolicyRuleOktaAppBasedResourceSelectorModel          `tfsdk:"okta_app_based_resource"`
	//Secret *SecurityPolicyRuleSecretBasedResourceSelectorModel `tfsdk:"secret_based_resource"`
	//ActiveDirectory *SecurityPolicyRuleActiveDirectoryBasedResourceSelectorModel `tfsdk:"active_directory_based_resource"`
}

func SecurityPolicyRuleResourceSelectorSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"server_based_resource": ServerBasedResourceSelectorSchema(), // server_based_resource
			//"managed_saas_app":   SecurityPolicyRuleManagedSaasAppBasedResourceSelectorSchema(),
			//"unmanaged_saas_app": SecurityPolicyRuleUnmanagedSaasAppBasedResourceSelectorSchema(),
			//"okta_app":           SecurityPolicyRuleOktaAppBasedResourceSelectorSchema(),

		},
		Required: true,
	}
}

func SecurityPolicyRuleResourceSelectorAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"server_based_resource": types.ObjectType{AttrTypes: ServerBasedResourceSelectorAttrTypes()},
	}
}

func SecurityPolicyRuleResourceSelectorFromModelToSDK(ctx context.Context, in types.Object) (*pam.SecurityPolicyRuleResourceSelector, diag.Diagnostics) {
	var out pam.SecurityPolicyRuleResourceSelector
	var diags diag.Diagnostics

	if !in.IsUnknown() && !in.IsNull() {
		var resourceSelectorModel SecurityPolicyRuleResourceSelectorModel

		diags.Append(in.As(ctx, &resourceSelectorModel, basetypes.ObjectAsOptions{})...)

		if diags.HasError() {
			return nil, diags
		}

		if resourceSelectorModel.ServerBasedResource != nil {
			outSelector, d := ServerBasedResourceSelectorFromModelToSDK(ctx, resourceSelectorModel.ServerBasedResource)
			diags.Append(d...)
			if diags.HasError() {
				return nil, diags
			}
			out.SecurityPolicyRuleServerBasedResourceSelector = outSelector
		}
	}
	return &out, diags
}

func SecurityPolicyRuleResourceSelectorFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRuleResourceSelector) (*types.Object, diag.Diagnostics) {
	var outModel SecurityPolicyRuleResourceSelectorModel
	var diags diag.Diagnostics

	if in.SecurityPolicyRuleServerBasedResourceSelector != nil {
		selector, d := ServerBasedResourceSelectorFromSDKToModel(ctx, in.SecurityPolicyRuleServerBasedResourceSelector)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		outModel.ServerBasedResource = selector

	} else {
		panic("missing stanza in SecurityPolicyRuleResourceSelectorFromSDKToModel")
	}

	objectValue, d := types.ObjectValueFrom(ctx, SecurityPolicyRuleResourceSelectorAttrTypes(), &outModel)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}
	return &objectValue, diags
}

type ServerBasedResourceSelectorModel struct {
	Selectors types.List /*ServerBasedResourceSelectorContainerModel*/ `tfsdk:"selectors"`
}

func ServerBasedResourceSelectorSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"selectors": schema.ListNestedAttribute{
				NestedObject: ServerBasedResourceSelectorContainerSchema(),
				Required:     true,
			},
		},
		Required: true,
	}
}

func ServerBasedResourceSelectorAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"selectors": types.ListType{ElemType: types.ObjectType{AttrTypes: ServerBasedResourceSelectorContainerAttrTypes()}},
	}
}

func ServerBasedResourceSelectorFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRuleServerBasedResourceSelector) (*ServerBasedResourceSelectorModel, diag.Diagnostics) {
	var out ServerBasedResourceSelectorModel
	var outSelectors []ServerBasedResourceSelectorContainerModel
	var diags diag.Diagnostics

	//TODO(ja) this should be wrapped in a Container.

	for _, selector := range in.Selectors {
		var container ServerBasedResourceSelectorContainerModel
		if selector.Selector.SelectorServerLabel != nil {
			outSelector, d := SelectorServerLabelFromSDKToModel(ctx, selector.Selector.SelectorServerLabel)
			diags.Append(d...)
			if diags.HasError() {
				return nil, diags
			}
			container.SelectorServerLabel = outSelector

		} else if selector.Selector.SelectorIndividualServer != nil {
			outSelector, d := SelectorIndividualServerFromSDKToModel(ctx, selector.Selector.SelectorIndividualServer)
			diags.Append(d...)
			if diags.HasError() {
				return nil, diags
			}
			container.IndividualServer = outSelector

		} else if selector.Selector.SelectorIndividualServerAccount != nil {
			outSelector, d := SelectorIndividualServerAccountFromSDKToModel(ctx, selector.Selector.SelectorIndividualServerAccount)
			diags.Append(d...)
			if diags.HasError() {
				return nil, diags
			}
			container.IndividualServerAccount = outSelector

		} else {
			panic("missing stanza in ServerBasedResourceSelectorFromSDKToModel")
		}
		outSelectors = append(outSelectors, container)
	}

	listValue, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: ServerBasedResourceSelectorContainerAttrTypes()}, outSelectors)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}
	out.Selectors = listValue

	return &out, diags
}

func ServerBasedResourceSelectorFromModelToSDK(ctx context.Context, in *ServerBasedResourceSelectorModel) (*pam.SecurityPolicyRuleServerBasedResourceSelector, diag.Diagnostics) {
	var out pam.SecurityPolicyRuleServerBasedResourceSelector
	var diags diag.Diagnostics

	out.Type = string(pam.SecurityPolicyRuleResourceType_SERVER_BASED_RESOURCE)
	if !in.Selectors.IsNull() && len(in.Selectors.Elements()) > 0 {
		selectorModels := make([]ServerBasedResourceSelectorContainerModel, 0, len(in.Selectors.Elements()))
		if diags := in.Selectors.ElementsAs(ctx, &selectorModels, false); diags.HasError() {
			return nil, diags
		}
		for _, selectorModel := range selectorModels {
			outSelector, d := ServerBasedResourceSelectorContainerFromModelToSDK(ctx, &selectorModel)
			diags.Append(d...)
			if diags.HasError() {
				return nil, diags
			}
			out.Selectors = append(out.Selectors, *outSelector)
		}
	}
	return &out, diags
}

type ServerBasedResourceSelectorContainerModel struct {
	IndividualServer        *SelectorIndividualServerModel        `tfsdk:"individual_server"`
	IndividualServerAccount *SelectorIndividualServerAccountModel `tfsdk:"individual_server_account"`
	SelectorServerLabel     *SelectorServerLabelModel             `tfsdk:"server_label"`
}

func ServerBasedResourceSelectorContainerSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"individual_server":         SelectorIndividualServerSchema(),
			"individual_server_account": SelectorIndividualServerAccountSchema(),
			"server_label":              SelectorServerLabelSchema(),
		},
	}
}

func ServerBasedResourceSelectorContainerAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"individual_server":         types.ObjectType{AttrTypes: SelectorIndividualServerAttrTypes()},
		"individual_server_account": types.ObjectType{AttrTypes: SelectorIndividualServerAccountAttrTypes()},
		"server_label":              types.ObjectType{AttrTypes: SelectorServerLabelAttrTypes()},
	}
}

func ServerBasedResourceSelectorContainerFromModelToSDK(ctx context.Context, in *ServerBasedResourceSelectorContainerModel) (*pam.SecurityPolicyRuleServerBasedResourceSelectorContainer, diag.Diagnostics) {
	var out pam.SecurityPolicyRuleServerBasedResourceSelectorContainer
	var diags diag.Diagnostics

	if in.IndividualServer != nil {
		outVal, d := SelectorIndividualServerFromModelToSDK(ctx, in.IndividualServer)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		out = *pam.NewSecurityPolicyRuleServerBasedResourceSelectorContainer(
			pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_INDIVIDUAL_SERVER,
			pam.SelectorIndividualServerAsSecurityPolicyRuleResourceServerBasedResourceSubSelector(outVal),
		)

	} else if in.IndividualServerAccount != nil {
		outVal, d := SelectorIndividualServerAccountFromModelToSDK(ctx, in.IndividualServerAccount)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		out = *pam.NewSecurityPolicyRuleServerBasedResourceSelectorContainer(
			pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_INDIVIDUAL_SERVER_ACCOUNT,
			pam.SelectorIndividualServerAccountAsSecurityPolicyRuleResourceServerBasedResourceSubSelector(outVal),
		)

	} else if in.SelectorServerLabel != nil {
		outVal, d := SelectorServerLabelFromModelToSDK(ctx, in.SelectorServerLabel)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		out = *pam.NewSecurityPolicyRuleServerBasedResourceSelectorContainer(
			pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_SERVER_LABEL,
			pam.SelectorServerLabelAsSecurityPolicyRuleResourceServerBasedResourceSubSelector(outVal))

	} else {
		panic("missing stanza in ServerBasedResourceSelectorContainerFromModelToSDK")
	}

	return &out, diags
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

func SelectorIndividualServerAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"server": types.StringType,
	}
}

func SelectorIndividualServerFromSDKToModel(_ context.Context, in *pam.SelectorIndividualServer) (*SelectorIndividualServerModel, diag.Diagnostics) {
	var out SelectorIndividualServerModel
	var diags diag.Diagnostics

	if val, ok := in.Server.GetIdOk(); ok {
		out.Server = types.StringPointerValue(val)
	}
	return &out, diags
}

func SelectorIndividualServerFromModelToSDK(_ context.Context, in *SelectorIndividualServerModel) (*pam.SelectorIndividualServer, diag.Diagnostics) {
	var out pam.SelectorIndividualServer
	var diags diag.Diagnostics

	out.Type = string(pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_INDIVIDUAL_SERVER)
	if !in.Server.IsNull() && !in.Server.IsUnknown() {
		out.Server = *pam.NewNamedObject().SetId(in.Server.ValueString())
	}
	return &out, diags
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

func SelectorIndividualServerAccountAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"server":   types.StringType,
		"username": types.StringType,
	}
}

func SelectorIndividualServerAccountFromSDKToModel(_ context.Context, in *pam.SelectorIndividualServerAccount) (*SelectorIndividualServerAccountModel, diag.Diagnostics) {
	var out SelectorIndividualServerAccountModel
	var diags diag.Diagnostics

	out.Username = types.StringPointerValue(in.Username)

	if server, ok := in.GetServerOk(); ok {
		out.Server = types.StringPointerValue(server.Id)
	}

	return &out, diags
}
func SelectorIndividualServerAccountFromModelToSDK(_ context.Context, in *SelectorIndividualServerAccountModel) (*pam.SelectorIndividualServerAccount, diag.Diagnostics) {
	var out pam.SelectorIndividualServerAccount
	var diags diag.Diagnostics

	out.Type = string(pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_INDIVIDUAL_SERVER_ACCOUNT)
	if !in.Username.IsNull() && !in.Username.IsUnknown() {
		out.Username = in.Username.ValueStringPointer()
	}

	if !in.Server.IsNull() && !in.Server.IsUnknown() {
		out.Server = *pam.NewNamedObject().SetId(in.Server.ValueString())
	}
	return &out, diags
}
