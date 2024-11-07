package convert

import (
	"context"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NamedObjectTypeModel types.String

// NamedObjectModel is really just a string. I'm not sure if this is the right way to do this.
type NamedObjectModel types.String

func NamedObjectFromSDKToModel(_ context.Context, in *pam.NamedObject, out *NamedObjectModel) diag.Diagnostics {
	if id, ok := in.GetIdOk(); ok {
		*out = NamedObjectModel(types.StringPointerValue(id))
	}
	return nil
}

func NamedObjectFromModelToSDK(_ context.Context, in NamedObjectModel, out *pam.NamedObject) diag.Diagnostics {
	vsp := types.String(in).ValueStringPointer()
	if vsp != nil {
		out.SetId(*vsp)
	}
	return nil
}
