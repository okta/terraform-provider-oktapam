package utils

import (
	"fmt"
	"github.com/okta/terraform-provider-oktapam/oktapam/client/wrappers"
	"strings"

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

func AsInt32PtrZero(v int, returnZero bool) *int32 {
	if !returnZero && v == 0 {
		return nil
	}
	v32 := int32(v)
	return &v32
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

func ExpandStringList(vI []any) []string {
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
	arr := make([]any, len(strings))
	for i, str := range strings {
		arr[i] = str
	}
	return schema.NewSet(schema.HashString, arr)
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

// GenerateAttributeOverrides will evaluate which attributes on the resource may have overrides set,
// It will get the existing values of those attributes, and return a map of any data found.
// The key in the map will be the single attribute value (rather than the full path).
// This is because d.Set cannot be called on nested attributes. i.e. a full list must be provided, and not a single
// element of that list.
//
// For example, the `password` field in a MySQL connection in a database resource is not provided by the GET API, so
// the initial value from creation must be copied over from the previous state on each new operation. This allows
// Terraform to detect and act on a password change if the plan is modified.
// This is not an issue for top-level fields as their initial value can automatically remain stored in state, but for
// nested attributes like `password`, if any parent attribute is set, such as attributes.ManagementConnectionDetails or
// attributes.MySQL, it will cause the nested attributes.Password to be overwritten and removed from state.
// Explicitly reading the attribute before state is modified and ensuring it gets written into the new state solves
// this problem.
func GenerateAttributeOverrides(d *schema.ResourceData, resource wrappers.ResourceWrapper) map[string]any {
	paths := resource.AttributeOverridePaths()
	existingValues := getAttributeValuesByPath(d, paths)

	// convert the full key path to only the final attribute field
	overrides := make(map[string]any, len(existingValues))
	for k, existingVal := range existingValues {
		parts := strings.Split(k, ".")
		lastAttr := parts[len(parts)-1]
		overrides[lastAttr] = existingVal
	}
	return overrides
}

func getAttributeValuesByPath(d *schema.ResourceData, keyPaths []string) map[string]any {
	keyValuePair := make(map[string]any, len(keyPaths))
	for _, key := range keyPaths {
		if val, exists := d.GetOk(key); exists {
			keyValuePair[key] = val
		}
	}
	return keyValuePair
}
