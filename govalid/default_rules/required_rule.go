package default_rules

import (
	"errors"
	"reflect"
	"strings"
)

func RequiredRule(field string, value interface{}, params ...interface{}) error {
	v := reflect.ValueOf(value)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return errors.New("required_error")
		}
		v = v.Elem()
	}
	if !v.IsValid() {
		return errors.New("required_error")
	}

	if v.IsZero() {
		return errors.New("required_error")
	}
	switch v.Kind() {
	case reflect.String:
		if len(strings.TrimSpace(v.String())) == 0 {
			return errors.New("required_error")
		}
	case reflect.Slice, reflect.Array, reflect.Map:
		if v.Len() == 0 {
			return errors.New("required_error")
		}
	}
	return nil
}
