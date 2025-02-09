package helpers

import "reflect"

// dereference dereferences a value until it is no longer a pointer or interface.
// If the value is nil, it is returned as is.
func Dereference(input interface{}) reflect.Value {
	v := reflect.ValueOf(input)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return reflect.Value{}
		}
		v = v.Elem()
	}
	return v
}
