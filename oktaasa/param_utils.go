package oktaasa

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getBool(attribute string, d *schema.ResourceData) (bool, error) {
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
