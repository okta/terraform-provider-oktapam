package oktapam

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func getOkBool(attribute string, d *schema.ResourceData) (bool, error) {
	if self, ok := d.GetOk(attribute); ok {
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

func getBoolPtr(attribute string, d *schema.ResourceData, returnZero bool) *bool {
	v := d.Get(attribute).(bool)
	return utils.AsBoolPtrZero(v, returnZero)
}

func getIntPtr(attribute string, d *schema.ResourceData, returnZero bool) *int {
	v := d.Get(attribute).(int)
	return utils.AsIntPtrZero(v, returnZero)
}

func getStringPtr(attribute string, d *schema.ResourceData, returnZero bool) *string {
	v := d.Get(attribute).(string)
	return utils.AsStringPtrZero(v, returnZero)
}
