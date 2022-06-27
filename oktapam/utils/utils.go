package utils

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

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
	for idx, v := range vI {
		val, ok := v.(string)
		if ok && val != "" {
			vs[idx] = val
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

type checkResourceExistsFunc func(string) (bool, error)

func CheckResourceExists(name string, checkResourceExists checkResourceExistsFunc) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resNotFoundErr := fmt.Errorf("resource not found: %s", name)
		// retrieve the resource by name from the state
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return resNotFoundErr
		}
		ID := rs.Primary.ID
		exist, err := checkResourceExists(ID)
		if err != nil {
			return err
		} else if !exist {
			return resNotFoundErr
		}
		return nil
	}
}

func CreateCheckResourceDestroy(typeName string, checkResourceExists checkResourceExistsFunc) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != typeName {
				continue
			}
			ID := rs.Primary.ID
			exists, err := checkResourceExists(ID)
			if err != nil {
				return err
			}
			if exists {
				return fmt.Errorf("resource still exists, ID: %s", ID)
			}
		}
		return nil
	}
}
