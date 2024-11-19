package convert

import (
	"context"
	"fmt"
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
	var resourceSelectorModel SecurityPolicyRuleResourceSelectorModel
	if diags := in.As(ctx, &resourceSelectorModel, basetypes.ObjectAsOptions{}); diags.HasError() {
		return nil, diags
	}

	if resourceSelectorModel.ServerBasedResource != nil {
		if outSelector, diags := ServerBasedResourceSelectorFromModelToSDK(ctx, resourceSelectorModel.ServerBasedResource); diags.HasError() {
			return nil, diags
		} else {
			out.SecurityPolicyRuleServerBasedResourceSelector = outSelector
		}
	}
	return &out, nil
}

func SecurityPolicyRuleResourceSelectorFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRuleResourceSelector) (*types.Object, diag.Diagnostics) {
	var outModel SecurityPolicyRuleResourceSelectorModel

	if in.SecurityPolicyRuleServerBasedResourceSelector != nil {
		if selector, diags := ServerBasedResourceSelectorFromSDKToModel(ctx, in.SecurityPolicyRuleServerBasedResourceSelector); diags.HasError() {
			return nil, diags
		} else {
			outModel.ServerBasedResource = selector
		}
	} else {
		panic("missing stanza in SecurityPolicyRuleResourceSelectorFromSDKToModel")
	}

	if objectValue, diags := types.ObjectValueFrom(ctx, SecurityPolicyRuleResourceSelectorAttrTypes(), &outModel); diags.HasError() {
		return nil, diags
	} else {
		return &objectValue, nil
	}
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

	//TODO(ja) this should be wrapped in a Container.

	for _, selector := range in.Selectors {
		var container ServerBasedResourceSelectorContainerModel
		if selector.Selector.SelectorServerLabel != nil {
			if outSelector, diags := SelectorServerLabelFromSDKToModel(ctx, selector.Selector.SelectorServerLabel); diags.HasError() {
				return nil, diags
			} else {
				container.SelectorServerLabel = outSelector
			}
		} else if selector.Selector.SelectorIndividualServer != nil {
			if outSelector, diags := SelectorIndividualServerFromSDKToModel(ctx, selector.Selector.SelectorIndividualServer); diags.HasError() {
				return nil, diags
			} else {
				container.IndividualServer = outSelector
			}
		} else if selector.Selector.SelectorIndividualServerAccount != nil {
			if outSelector, diags := SelectorIndividualServerAccountFromSDKToModel(ctx, selector.Selector.SelectorIndividualServerAccount); diags.HasError() {
				return nil, diags
			} else {
				container.IndividualServerAccount = outSelector
			}
		} else {
			panic("missing stanza in ServerBasedResourceSelectorFromSDKToModel")
		}
		outSelectors = append(outSelectors, container)
	}

	if listValue, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: ServerBasedResourceSelectorContainerAttrTypes()}, outSelectors); diags.HasError() {
		return nil, diags
	} else {
		out.Selectors = listValue
	}
	return &out, nil
}

func ServerBasedResourceSelectorFromModelToSDK(ctx context.Context, in *ServerBasedResourceSelectorModel) (*pam.SecurityPolicyRuleServerBasedResourceSelector, diag.Diagnostics) {
	var out pam.SecurityPolicyRuleServerBasedResourceSelector
	out.Type = string(pam.SecurityPolicyRuleResourceType_SERVER_BASED_RESOURCE)
	if !in.Selectors.IsNull() && len(in.Selectors.Elements()) > 0 {
		selectorModels := make([]ServerBasedResourceSelectorContainerModel, 0, len(in.Selectors.Elements()))
		if diags := in.Selectors.ElementsAs(ctx, &selectorModels, false); diags.HasError() {
			return nil, diags
		}
		for _, selectorModel := range selectorModels {
			if outSelector, diags := ServerBasedResourceSelectorContainerFromModelToSDK(ctx, &selectorModel); diags.HasError() {
				return nil, diags
			} else {
				out.Selectors = append(out.Selectors, *outSelector)
			}
		}
	}
	return &out, nil
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

	if in.IndividualServer != nil {
		if outVal, diags := SelectorIndividualServerFromModelToSDK(ctx, in.IndividualServer); diags.HasError() {
			return nil, diags
		} else {
			out = *pam.NewSecurityPolicyRuleServerBasedResourceSelectorContainer(
				pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_INDIVIDUAL_SERVER,
				pam.SelectorIndividualServerAsSecurityPolicyRuleResourceServerBasedResourceSubSelector(outVal),
			)
		}
	} else if in.IndividualServerAccount != nil {
		if outVal, diags := SelectorIndividualServerAccountFromModelToSDK(ctx, in.IndividualServerAccount); diags.HasError() {
			return nil, diags
		} else {
			out = *pam.NewSecurityPolicyRuleServerBasedResourceSelectorContainer(
				pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_INDIVIDUAL_SERVER_ACCOUNT,
				pam.SelectorIndividualServerAccountAsSecurityPolicyRuleResourceServerBasedResourceSubSelector(outVal),
			)
		}
	} else if in.SelectorServerLabel != nil {
		if outVal, diags := SelectorServerLabelFromModelToSDK(ctx, in.SelectorServerLabel); diags.HasError() {
			return nil, diags
		} else {
			out = *pam.NewSecurityPolicyRuleServerBasedResourceSelectorContainer(
				pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_SERVER_LABEL,
				pam.SelectorServerLabelAsSecurityPolicyRuleResourceServerBasedResourceSubSelector(outVal))
		}
	} else {
		panic("missing stanza in ServerBasedResourceSelectorContainerFromModelToSDK")
	}

	return &out, nil
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
	if val, ok := in.Server.GetIdOk(); ok {
		out.Server = types.StringPointerValue(val)
	}
	return &out, nil
}

func SelectorIndividualServerFromModelToSDK(_ context.Context, in *SelectorIndividualServerModel) (*pam.SelectorIndividualServer, diag.Diagnostics) {
	var out pam.SelectorIndividualServer
	out.Type = string(pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_INDIVIDUAL_SERVER)
	if !in.Server.IsNull() {
		out.Server = *pam.NewNamedObject().SetId(in.Server.ValueString())
	}
	return &out, nil
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

	out.Username = types.StringPointerValue(in.Username)

	if server, ok := in.GetServerOk(); ok {
		out.Server = types.StringPointerValue(server.Id)
	}

	return &out, nil
}
func SelectorIndividualServerAccountFromModelToSDK(_ context.Context, in *SelectorIndividualServerAccountModel) (*pam.SelectorIndividualServerAccount, diag.Diagnostics) {
	var out pam.SelectorIndividualServerAccount
	out.Type = string(pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_INDIVIDUAL_SERVER_ACCOUNT) //TODO(ja) this should probably be hard coded
	out.Username = in.Username.ValueStringPointer()
	if !in.Server.IsNull() && !in.Server.IsUnknown() {
		out.Server = *pam.NewNamedObject().SetId(in.Server.ValueString())
	}
	out.Username = in.Username.ValueStringPointer()
	return &out, nil
}

type SelectorServerLabelModel struct {
	ServerSelector  *ServerLabelServerSelectorModel `tfsdk:"server_selector"`
	AccountSelector types.List/*types.String*/ `tfsdk:"account_selector"`
}

func SelectorServerLabelSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"server_selector":  ServerLabelServerSelectorSchema(),
			"account_selector": schema.ListAttribute{ElementType: types.StringType, Optional: true},
		},
		Optional: true,
	}
}

func SelectorServerLabelAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"server_selector":  types.ObjectType{AttrTypes: ServerLabelServerSelectorAttrTypes()},
		"account_selector": types.ListType{ElemType: types.StringType},
	}
}

func SelectorServerLabelFromSDKToModel(ctx context.Context, in *pam.SelectorServerLabel) (*SelectorServerLabelModel, diag.Diagnostics) {
	var out SelectorServerLabelModel
	if selector, diags := ServerLabelServerSelectorFromSDKToModel(ctx, &in.ServerSelector); diags.HasError() {
		return nil, diags
	} else {
		out.ServerSelector = selector
	}

	var usernames []types.String
	if in.AccountSelector.HasUsernames() {
		//TODO(ja) this is supposed to be a separate model
		for _, username := range in.AccountSelector.Usernames {
			usernames = append(usernames, types.StringValue(username))
		}
		if usernameList, diags := types.ListValueFrom(ctx, types.StringType, usernames); diags.HasError() {
			return nil, diags
		} else {
			out.AccountSelector = usernameList
		}
	} else {
		out.AccountSelector = types.ListNull(types.StringType)
	}
	return &out, nil
}
func SelectorServerLabelFromModelToSDK(ctx context.Context, in *SelectorServerLabelModel) (*pam.SelectorServerLabel, diag.Diagnostics) {
	var out pam.SelectorServerLabel

	out.Type = string(pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_SERVER_LABEL)
	if outSelector, diags := ServerLabelServerSelectorFromModelToSDK(ctx, in.ServerSelector); diags.HasError() {
		return nil, diags
	} else {
		out.ServerSelector = *outSelector
	}

	if !in.AccountSelector.IsNull() && len(in.AccountSelector.Elements()) > 0 {
		accountSelectorModel := make([]types.String, 0, len(in.AccountSelector.Elements()))
		if diags := in.AccountSelector.ElementsAs(ctx, &accountSelectorModel, false); diags.HasError() {
			return nil, diags
		}
		var outAccountSelector pam.SelectorServerLabelAccountSelector
		for _, elem := range accountSelectorModel {
			if !elem.IsNull() && !elem.IsUnknown() {
				outAccountSelector.Usernames = append(outAccountSelector.Usernames, elem.ValueString())
			}
		}
		out.AccountSelector = &outAccountSelector
	}
	return &out, nil
}

type ServerLabelServerSelectorModel struct {
	Labels types.Map `tfsdk:"labels"`
}

func ServerLabelServerSelectorSchema() schema.Attribute {
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

func ServerLabelServerSelectorAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"labels": types.MapType{ElemType: types.StringType},
	}
}

func ServerLabelServerSelectorFromModelToSDK(ctx context.Context, in *ServerLabelServerSelectorModel) (*pam.SelectorServerLabelServerSelector, diag.Diagnostics) {
	var out pam.SelectorServerLabelServerSelector
	elements := make(map[string]types.String, len(in.Labels.Elements()))
	if diags := in.Labels.ElementsAs(ctx, &elements, false); diags.HasError() {
		return nil, diags
	}
	outMap := make(map[string]any, len(elements))
	for k, v := range elements {
		outMap[k] = v.ValueString()
	}
	out.Labels = outMap
	return &out, nil
}

func ServerLabelServerSelectorFromSDKToModel(_ context.Context, in *pam.SelectorServerLabelServerSelector) (*ServerLabelServerSelectorModel, diag.Diagnostics) {
	var out ServerLabelServerSelectorModel
	elements := make(map[string]attr.Value, len(in.Labels))
	for k, v := range in.Labels {
		elements[k] = types.StringValue(fmt.Sprintf("%s", v))
	}

	if mapValue, diags := types.MapValue(types.StringType, elements); diags.HasError() {
		return nil, diags
	} else {
		out.Labels = mapValue
	}
	return &out, nil
}
