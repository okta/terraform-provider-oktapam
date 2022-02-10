package utils

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
