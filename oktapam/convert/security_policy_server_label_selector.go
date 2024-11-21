package convert

import (
	"context"
	"fmt"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SelectorServerLabelModel struct {
	ServerSelector      *ServerLabelServerSelectorModel  `tfsdk:"server_selector"`
	AccountSelectorType types.String                     `tfsdk:"account_selector_type"`
	AccountSelector     *ServerLabelAccountSelectorModel `tfsdk:"account_selector"`
}

func SelectorServerLabelSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"server_selector":       ServerLabelServerSelectorSchema(),
			"account_selector_type": schema.StringAttribute{Required: true},
			"account_selector":      ServerLabelAccountSelectorSchema(),
		},
		Optional: true,
	}
}

func SelectorServerLabelAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"server_selector":       types.ObjectType{AttrTypes: ServerLabelServerSelectorAttrTypes()},
		"account_selector_type": types.StringType,
		"account_selector":      types.ObjectType{AttrTypes: ServerLabelAccountSelectorAttrTypes()},
	}
}

func SelectorServerLabelFromSDKToModel(ctx context.Context, in *pam.SelectorServerLabel) (*SelectorServerLabelModel, diag.Diagnostics) {
	var out SelectorServerLabelModel
	if selector, diags := ServerLabelServerSelectorFromSDKToModel(ctx, &in.ServerSelector); diags.HasError() {
		return nil, diags
	} else {
		out.ServerSelector = selector
	}

	if val, ok := in.GetAccountSelectorTypeOk(); ok {
		valStr := string(*val)
		out.AccountSelectorType = types.StringValue(valStr)
	}

	if accountSelector, diags := ServerLabelAccountSelectorFromSDKToModel(ctx, &in.AccountSelector); diags.HasError() {
		return nil, diags
	} else {
		out.AccountSelector = accountSelector
	}

	return &out, nil
}
func ServerLabelAccountSelectorFromSDKToModel(ctx context.Context, in *pam.SelectorServerLabelAccountSelector) (*ServerLabelAccountSelectorModel, diag.Diagnostics) {
	var out ServerLabelAccountSelectorModel

	switch selector := in.GetActualInstance().(type) {
	case *pam.SecurityPolicyUsernameAccountSelector:
		var usernames []types.String
		for _, username := range selector.Usernames {
			usernames = append(usernames, types.StringValue(username))
		}
		if usernameList, diags := types.ListValueFrom(ctx, types.StringType, usernames); diags.HasError() {
			return nil, diags
		} else {
			out.Usernames = usernameList
		}
	case *pam.SecurityPolicyNoneAccountSelector:
		out.Usernames = types.ListNull(types.StringType)
		// NOTE(ja): We require API callers to provide an empty struct.
	default:
		panic("missing stanza in ServerLabelAccountSelectorFromSDKToModel")
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

	// NOTE: The way this is done is snowflake. AccountSelector is the only place where we have the type on the outside
	// but also have extra fields, and do a oneOf for the actual AccountSelector. Please don't follow this model.

	if !in.AccountSelectorType.IsNull() && !in.AccountSelectorType.IsUnknown() {
		out.AccountSelectorType = pam.SelectorServerLabelAccountSelectorType(in.AccountSelectorType.ValueString())
	}

	if outAccountSelector, diags := ServerLabelAccountSelectorFromModelToSDK(ctx, in.AccountSelector); diags.HasError() {
		return nil, diags
	} else {
		out.AccountSelector = *outAccountSelector
	}

	return &out, nil
}
func ServerLabelAccountSelectorFromModelToSDK(ctx context.Context, in *ServerLabelAccountSelectorModel) (*pam.SelectorServerLabelAccountSelector, diag.Diagnostics) {
	var out pam.SelectorServerLabelAccountSelector

	if !in.Usernames.IsNull() && !in.Usernames.IsUnknown() {
		var outUsername pam.SecurityPolicyUsernameAccountSelector
		outUsername.Type = pam.PtrString(string(pam.SelectorServerLabelAccountSelectorType_USERNAME))
		usernameModel := make([]types.String, 0, len(in.Usernames.Elements()))
		if diags := in.Usernames.ElementsAs(ctx, &usernameModel, false); diags.HasError() {
			return nil, diags
		}
		for _, elem := range usernameModel {
			if !elem.IsNull() && !elem.IsUnknown() {
				outUsername.Usernames = append(outUsername.Usernames, elem.ValueString())
			}
		}
		out.SecurityPolicyUsernameAccountSelector = &outUsername
	} else {
		var outNone pam.SecurityPolicyNoneAccountSelector
		outNone.Type = pam.PtrString(string(pam.SelectorServerLabelAccountSelectorType_NONE))
		out.SecurityPolicyNoneAccountSelector = &outNone
	}
	return &out, nil
}

type ServerLabelAccountSelectorModel struct {
	Usernames types.List `tfsdk:"usernames"`
}

func ServerLabelAccountSelectorSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"usernames": schema.ListAttribute{ElementType: types.StringType, Optional: true}, // if type==none
		},
		Required: true,
	}
}

func ServerLabelAccountSelectorAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"usernames": types.ListType{ElemType: types.StringType},
	}
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

	if len(in.Labels.Elements()) > 0 {
		elements := make(map[string]types.String, len(in.Labels.Elements()))
		if diags := in.Labels.ElementsAs(ctx, &elements, false); diags.HasError() {
			return nil, diags
		}
		outMap := make(map[string]any, len(elements))
		for k, v := range elements {
			outMap[k] = v.ValueString()
		}
		out.Labels = outMap
	}
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
