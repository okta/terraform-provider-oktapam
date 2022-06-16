package utils

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func IsBlank(s *string) bool {
	return s == nil || *s == ""
}

func IsNonEmpty(s *string) bool {
	return s != nil && *s != ""
}

func AsIntPtr(v int) *int {
	return AsIntPtrZero(v, false)
}

func AsIntPtrZero(v int, returnZero bool) *int {
	if !returnZero && v == 0 {
		return nil
	}
	return &v
}

func AsBoolPtr(v bool) *bool {
	return AsBoolPtrZero(v, false)
}

func AsBoolPtrZero(v bool, returnZero bool) *bool {
	if !returnZero && !v {
		return nil
	}
	return &v
}

func AsStringPtr(v string) *string {
	return AsStringPtrZero(v, false)

}

func AsStringPtrZero(v string, returnZero bool) *string {
	if !returnZero && v == "" {
		return nil
	}
	return &v
}

func ExpandStringSet(v *schema.Set) []string {
	return ExpandStringList(v.List())
}

func ExpandStringList(vI []interface{}) []string {
	vs := make([]string, len(vI))
	for _, v := range vI {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, val)
		}
	}
	return vs
}

func ConvertStringSliceToSet(strings []string) *schema.Set {
	arr := make([]interface{}, len(strings))
	for i, str := range strings {
		arr[i] = str
	}
	return schema.NewSet(schema.HashString, arr)
}

func SetNonPrimitives(d *schema.ResourceData, valueMap map[string]interface{}) error {
	for k, v := range valueMap {
		if v != nil {
			if err := d.Set(k, v); err != nil {
				return fmt.Errorf("error setting %s for resource %s: %s", k, d.Id(), err)
			}
		}
	}
	return nil
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
