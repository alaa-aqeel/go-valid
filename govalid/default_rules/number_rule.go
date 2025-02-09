package default_rules

import (
	"errors"
	"math"
	"reflect"
	"strconv"

	"github.com/alaa-aqeel/go-valid/govalid/helpers"
)

func MinRule(field string, value interface{}, params ...interface{}) error {
	if len(params) < 1 {
		return errors.New("min_params_error")
	}
	boundary, err := parseToFloat(params[0])
	if err != nil {
		return errors.New("min_params_error")
	}
	v := helpers.Dereference(value)
	if !v.IsValid() {
		return errors.New("min_value_error")
	}
	val, ok := toFloat64(v)
	if !ok {
		return errors.New("min_value_error")
	}
	if val < boundary {
		return errors.New("min_error")
	}

	return nil
}

func MaxRule(field string, value interface{}, params ...interface{}) error {
	if len(params) < 1 {
		return errors.New("max_params_error")
	}
	boundary, err := parseToFloat(params[0])
	if err != nil {
		return errors.New("max_params_error")
	}
	v := helpers.Dereference(value)
	if !v.IsValid() {
		return errors.New("max_value_error")
	}
	val, ok := toFloat64(v)
	if !ok {
		return errors.New("max_value_error")
	}
	if val > boundary {
		return errors.New("max_error")
	}

	return nil
}

func IsNumericRule(field string, value interface{}, params ...interface{}) error {
	v := helpers.Dereference(value)

	_, ok := toFloat64(v)
	if !ok {
		return errors.New("numeric_error")
	}
	return nil
}

func IsIntegerRule(field string, value interface{}, params ...interface{}) error {
	v := helpers.Dereference(value)

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return nil
	case reflect.String:
		_, err := strconv.ParseUint(v.String(), 16, 64)
		if err != nil {
			return errors.New("integer_error")
		}
		return nil
	default:
		return errors.New("integer_error")
	}
}

func toFloat64(v reflect.Value) (float64, bool) {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(v.Uint()), true
	case reflect.Float32, reflect.Float64:
		f := v.Float()
		if math.IsNaN(f) || math.IsInf(f, 0) {
			return 0, false
		}
		return f, true
	case reflect.String:
		f, err := strconv.ParseFloat(v.String(), 64)
		return f, err == nil
	default:
		return 0, false
	}
}

func parseToFloat(param interface{}) (float64, error) {
	switch v := param.(type) {
	case string:
		return strconv.ParseFloat(v, 64)
	case int, int64, float32, float64:
		return reflect.ValueOf(v).Convert(reflect.TypeOf(float64(0))).Float(), nil
	case interface{}:
		return strconv.ParseFloat(v.(string), 64)
	default:
		return 0, errors.New("invalid_parameter_type")
	}
}
