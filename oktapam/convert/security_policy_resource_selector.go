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

func SecurityPolicyRuleResourceSelectorFromModelToSDK(ctx context.Context, in types.Object, out *pam.SecurityPolicyRuleResourceSelector) diag.Diagnostics {
	var resourceSelectorModel SecurityPolicyRuleResourceSelectorModel
	if diags := in.As(ctx, &resourceSelectorModel, basetypes.ObjectAsOptions{}); diags.HasError() {
		return diags
	}

	if resourceSelectorModel.ServerBasedResource != nil {
		var outSelector pam.SecurityPolicyRuleServerBasedResourceSelector
		if diags := ServerBasedResourceSelectorFromModelToSDK(ctx, resourceSelectorModel.ServerBasedResource, &outSelector); diags.HasError() {
			return diags
		}
		out.SecurityPolicyRuleServerBasedResourceSelector = &outSelector
	}
	return nil
}

func SecurityPolicyRuleResourceSelectorFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRuleResourceSelector, out *types.Object) diag.Diagnostics {
	var outModel SecurityPolicyRuleResourceSelectorModel

	if in.SecurityPolicyRuleServerBasedResourceSelector != nil {
		if diags := ServerBasedResourceSelectorFromSDKToModel(ctx, in.SecurityPolicyRuleServerBasedResourceSelector, outModel.ServerBasedResource); diags.HasError() {
			return diags
		}
	} else {
		panic("missing stanza in SecurityPolicyRuleResourceSelectorFromSDKToModel")
	}

	if objectValue, diags := types.ObjectValueFrom(ctx, SecurityPolicyRuleResourceSelectorAttrTypes(), &outModel); diags.HasError() {
		return diags
	} else {
		out = &objectValue
	}
	return nil
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

func ServerBasedResourceSelectorFromSDKToModel(ctx context.Context, in *pam.SecurityPolicyRuleServerBasedResourceSelector, out *ServerBasedResourceSelectorModel) diag.Diagnostics {
	var outSelectors []ServerBasedResourceSelectorContainerModel

	//TODO(ja) this should be wrapped in a Container.

	for _, selector := range in.Selectors {
		var container ServerBasedResourceSelectorContainerModel
		if selector.Selector.SelectorServerLabel != nil {
			if diags := SelectorServerLabelFromSDKToModel(ctx, selector.Selector.SelectorServerLabel, container.SelectorServerLabel); diags.HasError() {
				return diags
			}
		} else if selector.Selector.SelectorIndividualServer != nil {
			if diags := SelectorIndividualServerFromSDKToModel(ctx, selector.Selector.SelectorIndividualServer, container.IndividualServer); diags.HasError() {
				return diags
			}
		} else if selector.Selector.SelectorIndividualServerAccount != nil {
			//if diags := SelectorIndividualServerAccountFromSDKToModel(ctx, selector.Selector.SelectorIndividualServerAccount, container.IndividualServerAccount); diags.HasError() {
			//	return diags
			//}
			// TODO(ja) -^
		} else {
			panic("missing stanza in ServerBasedResourceSelectorFromSDKToModel")
		}
		outSelectors = append(outSelectors, container)
	}

	if listValue, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: ServerBasedResourceSelectorAttrTypes()}, outSelectors); diags.HasError() {
		return diags
	} else {
		out.Selectors = listValue
	}
	return nil
}

func ServerBasedResourceSelectorFromModelToSDK(ctx context.Context, in *ServerBasedResourceSelectorModel, out *pam.SecurityPolicyRuleServerBasedResourceSelector) diag.Diagnostics {
	out.Type = string(pam.SecurityPolicyRuleResourceType_SERVER_BASED_RESOURCE)
	if !in.Selectors.IsNull() && len(in.Selectors.Elements()) > 0 {
		selectorModels := make([]ServerBasedResourceSelectorContainerModel, 0, len(in.Selectors.Elements()))
		if diags := in.Selectors.ElementsAs(ctx, &selectorModels, false); diags.HasError() {
			return diags
		}
		for _, selectorModel := range selectorModels {
			var outSelector pam.SecurityPolicyRuleServerBasedResourceSelectorContainer
			if diags := ServerBasedResourceSelectorContainerFromModelToSDK(ctx, &selectorModel, &outSelector); diags.HasError() {
				return diags
			}
			out.Selectors = append(out.Selectors, outSelector)
		}
	}
	return nil
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

func ServerBasedResourceSelectorContainerFromModelToSDK(ctx context.Context, in *ServerBasedResourceSelectorContainerModel, out *pam.SecurityPolicyRuleServerBasedResourceSelectorContainer) diag.Diagnostics {
	var selectorContainer *pam.SecurityPolicyRuleServerBasedResourceSelectorContainer

	if in.IndividualServer != nil {
		var outVal pam.SelectorIndividualServer
		if diags := SelectorIndividualServerFromModelToSDK(ctx, in.IndividualServer, &outVal); diags.HasError() {
			return diags
		}
		selectorContainer = pam.NewSecurityPolicyRuleServerBasedResourceSelectorContainer(
			pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_INDIVIDUAL_SERVER,
			pam.SelectorIndividualServerAsSecurityPolicyRuleResourceServerBasedResourceSubSelector(&outVal),
		)
		//out.Selector.SelectorIndividualServer = &outVal
	} else if in.IndividualServerAccount != nil {
		var outVal pam.SelectorIndividualServerAccount
		if diags := SelectorIndividualServerAccountFromModelToSDK(ctx, in.IndividualServerAccount, &outVal); diags.HasError() {
			return diags
		}
		selectorContainer = pam.NewSecurityPolicyRuleServerBasedResourceSelectorContainer(
			pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_INDIVIDUAL_SERVER_ACCOUNT,
			pam.SelectorIndividualServerAccountAsSecurityPolicyRuleResourceServerBasedResourceSubSelector(&outVal),
		)
	} else if in.SelectorServerLabel != nil {
		var outVal pam.SelectorServerLabel
		if diags := SelectorServerLabelFromModelToSDK(ctx, in.SelectorServerLabel, &outVal); diags.HasError() {
			return diags
		}
		selectorContainer = pam.NewSecurityPolicyRuleServerBasedResourceSelectorContainer(
			pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_SERVER_LABEL,
			pam.SelectorServerLabelAsSecurityPolicyRuleResourceServerBasedResourceSubSelector(&outVal))
	} else {
		panic("missing stanza in ServerBasedResourceSelectorContainerFromModelToSDK")
	}
	out = selectorContainer
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

func SelectorIndividualServerAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"server": types.StringType,
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

func SelectorIndividualServerAccountAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"server":   types.StringType,
		"username": types.StringType,
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

func SelectorServerLabelFromSDKToModel(ctx context.Context, in *pam.SelectorServerLabel, out *SelectorServerLabelModel) diag.Diagnostics {
	if diags := ServerLabelServerSelectorFromSDKToModel(ctx, &in.ServerSelector, out.ServerSelector); diags.HasError() {
		return diags
	}

	var usernames []types.String
	//TODO(ja) this is supposed to be a separate model
	for _, username := range in.AccountSelector.Usernames {
		usernames = append(usernames, types.StringValue(username))
	}
	if usernameList, diags := types.ListValueFrom(ctx, types.StringType, usernames); diags.HasError() {
		return diags
	} else {
		out.AccountSelector = usernameList
	}

	return nil
}
func SelectorServerLabelFromModelToSDK(ctx context.Context, in *SelectorServerLabelModel, out *pam.SelectorServerLabel) diag.Diagnostics {

	out.Type = string(pam.SecurityPolicyRuleServerBasedResourceSubSelectorType_SERVER_LABEL) //TODO(ja) this should probably be a hard-coded string
	if diags := ServerLabelServerSelectorFromModelToSDK(ctx, in.ServerSelector, &out.ServerSelector); diags.HasError() {
		return diags
	}

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

func ServerLabelServerSelectorFromModelToSDK(ctx context.Context, in *ServerLabelServerSelectorModel, out *pam.SelectorServerLabelServerSelector) diag.Diagnostics {
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

func ServerLabelServerSelectorFromSDKToModel(_ context.Context, in *pam.SelectorServerLabelServerSelector, out *ServerLabelServerSelectorModel) diag.Diagnostics {
	elements := make(map[string]attr.Value, len(in.Labels))
	for k, v := range in.Labels {
		elements[k] = types.StringValue(fmt.Sprintf("%s", v))
	}

	if mapValue, diags := types.MapValue(types.StringType, elements); diags.HasError() {
		return diags
	} else {
		out.Labels = mapValue
	}
	return nil
}
