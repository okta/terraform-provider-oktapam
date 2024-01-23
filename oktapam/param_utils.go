package oktapam

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

var uuidRegex = regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
var nameMatcher = regexp.MustCompile(`^[\w\-_.]+$`)

func MatchesUUID(s string) bool {
	return uuidRegex.MatchString(s)
}

func MatchesSimpleName(s string) bool {
	return nameMatcher.MatchString(s)
}

func GetStringSliceFromResource(attrName string, d *schema.ResourceData, optional bool) ([]string, diag.Diagnostics) {
	attrsI, ok := d.GetOk(attrName)
	if !ok {
		if !optional {
			return nil, diag.FromErr(fmt.Errorf("value for %s was not present", attrName))
		} else {
			return nil, nil
		}
	}

	return GetStringSlice(attrsI, attrName)
}

func GetUUIDSlice(attr any, attrName string) ([]string, diag.Diagnostics) {
	uuids, diags := GetStringSlice(attr, attrName)
	if diags != nil {
		return nil, diags
	}

	for _, uuid := range uuids {
		if !MatchesUUID(uuid) {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("value for %s must be a UUID: %s", attrName, uuid),
			})
		}
	}

	return uuids, diags
}

func GetStringSlice(attr any, attrName string) ([]string, diag.Diagnostics) {
	var diags diag.Diagnostics

	switch attrsI := attr.(type) {
	case []any:
		strs := make([]string, len(attrsI))
		for idx, strI := range attrsI {
			if str, ok := strI.(string); ok {
				strs[idx] = str
			} else {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  fmt.Sprintf("values for %s must be strings", attrName),
				})
			}
		}

		return strs, diags
	case *schema.Set:
		return GetStringSlice(attrsI.List(), attrName)
	default:
		panic(fmt.Sprintf("cannot convert type to string slice: %T", attr))
	}

}

func ConvertToNamedObject(id string, objectType client.NamedObjectType) client.NamedObject {
	return client.NamedObject{
		Id:   &id,
		Type: objectType,
	}
}

func ConvertInterfaceSliceToNamedObjectSlice(ids []any, objectType client.NamedObjectType) []client.NamedObject {
	namedObjects := make([]client.NamedObject, len(ids))
	for idx, id := range ids {
		id := id.(string)
		namedObjects[idx] = client.NamedObject{
			Id:   &id,
			Type: objectType,
		}
	}
	return namedObjects
}

func ConvertToNamedObjectSlice(ids []string, objectType client.NamedObjectType) []client.NamedObject {
	namedObjects := make([]client.NamedObject, len(ids))
	for idx, id := range ids {
		id := id
		namedObjects[idx] = client.NamedObject{
			Id:   &id,
			Type: objectType,
		}
	}
	return namedObjects
}

func GetTypeListMapFromResource(attr string, d *schema.ResourceData) map[string]any {
	listI := d.Get(attr)
	if listI == nil {
		return nil
	}

	listArr := listI.([]any)
	if len(listArr) == 0 {
		return nil
	}

	return listArr[0].(map[string]any)
}

func GetTypeListMapFromResourceElement(attr string, data map[string]any) map[string]any {
	listI := data[attr]
	if listI == nil {
		return nil
	}

	listArr := listI.([]any)
	if len(listArr) == 0 {
		return nil
	}

	return listArr[0].(map[string]any)
}

func GetMapPtrFromResource[V any](attr string, d *schema.ResourceData) (*map[string]V, diag.Diagnostics) {
	var returnedMap *map[string]V
	if v, ok := d.GetOk(attr); ok {
		if m, ok := v.(map[string]any); ok {
			tmpMap := make(map[string]V, len(m))
			for k, v := range m {
				if value, ok := v.(V); ok {
					tmpMap[k] = value
				}
			}
			returnedMap = &tmpMap
		} else {
			return nil, diag.FromErr(fmt.Errorf("invalid %s", attr))
		}
	}
	return returnedMap, nil
}

func GetOkBoolFromResource(attr string, d *schema.ResourceData) (bool, error) {
	if self, ok := d.GetOk(attr); ok {
		switch v := self.(type) {
		case bool:
			return v, nil
		case int:
			return v != 0, nil
		case string:
			b, err := strconv.ParseBool(v)
			if err != nil {
				return false, err
			}
			return b, nil
		default:
			return false, fmt.Errorf("cannot convert %T to bool", v)
		}
	}
	// zero value will result in !ok
	return false, nil
}

func GetBoolPtrFromResource(attr string, d *schema.ResourceData, returnZero bool) *bool {
	v := d.Get(attr).(bool)
	return utils.AsBoolPtrZero(v, returnZero)
}

func GetBoolPtrFromElement(attr string, data map[string]any, returnZero bool) *bool {
	if data[attr] == nil {
		return utils.AsBoolPtrZero(false, returnZero)
	}
	v := data[attr].(bool)
	return utils.AsBoolPtrZero(v, returnZero)
}

func GetInt32FromResource(attr string, d *schema.ResourceData) int32 {
	val := d.Get(attr).(int)
	return int32(val)
}

func GetInt32PtrFromResource(attr string, d *schema.ResourceData, returnZero bool) *int32 {
	val := GetIntPtrFromResource(attr, d, returnZero)
	if val == nil {
		return nil
	}
	int32Val := int32(*val)
	return &int32Val
}

func GetIntPtrFromResource(attr string, d *schema.ResourceData, returnZero bool) *int {
	v := d.Get(attr).(int)
	return utils.AsIntPtrZero(v, returnZero)
}

func GetIntPtrFromElement(attr string, data map[string]any, returnZero bool) *int {
	v := data[attr].(int)
	return utils.AsIntPtrZero(v, returnZero)
}

func GetStringPtrFromElement(attr string, data map[string]any, returnZero bool) *string {
	v := data[attr].(string)
	return utils.AsStringPtrZero(v, returnZero)
}

func GetStringPtrFromResource(attr string, d *schema.ResourceData, returnZero bool) *string {
	v := d.Get(attr).(string)
	return utils.AsStringPtrZero(v, returnZero)
}
