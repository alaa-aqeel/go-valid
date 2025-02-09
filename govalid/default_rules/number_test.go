package default_rules

import (
	"testing"
	"time"
)

type MyInt int
type MyUint uint
type MyFloat float64

func TestNumericMinMax(t *testing.T) {
	var (
		nilPtr  *int
		five    = 5
		fivePtr = &five
	)

	tests := []struct {
		name      string
		validator func(string, interface{}, ...interface{}) error
		value     interface{}
		param     interface{}
		expected  bool
	}{
		// Min
		{name: "Min int valid", validator: MinRule, value: 5, param: 3, expected: true},
		{name: "Min int equal", validator: MinRule, value: 5, param: 5, expected: true},
		{name: "Min int invalid", validator: MinRule, value: 2, param: 3, expected: false},
		{name: "Min uint valid", validator: MinRule, value: uint(5), param: 3, expected: true},
		{name: "Min float valid", validator: MinRule, value: 5.5, param: 3.2, expected: true},
		{name: "Min pointer valid", validator: MinRule, value: fivePtr, param: 3, expected: true},
		{name: "Min nil pointer invalid", validator: MinRule, value: nilPtr, param: 0, expected: false},
		{name: "Min invalid param", validator: MinRule, value: 5, param: "3", expected: false},
		{name: "Min invalid value", validator: MinRule, value: "five", param: 3, expected: false},

		// Max
		{name: "Max int valid", validator: MaxRule, value: 5, param: 5, expected: true},
		{name: "Max int invalid", validator: MaxRule, value: 6, param: 5, expected: false},
		{name: "Max uint valid", validator: MaxRule, value: uint(5), param: 5, expected: true},
		{name: "Max float valid", validator: MaxRule, value: 5.0, param: 5.0, expected: true},
		{name: "Max pointer valid", validator: MaxRule, value: fivePtr, param: 5, expected: true},
		{name: "Max nil pointer invalid", validator: MaxRule, value: nilPtr, param: 0, expected: false},
		{name: "Max invalid param", validator: MaxRule, value: 5, param: "5", expected: false},
		{name: "Max invalid value", validator: MaxRule, value: "five", param: 5, expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.validator("test", tt.value, tt.param)
			if result != nil {
				t.Errorf("%s: Expected %v, got %v (value: %v, param: %v)",
					tt.name, tt.expected, result.Error(), tt.value, tt.param)
			}
		})
	}
}

func TestIsNumeric(t *testing.T) {
	var nilPtr *int
	five := 5
	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		// Basic numeric types
		{name: "int", value: 42, expected: true},
		{name: "int8", value: int8(8), expected: true},
		{name: "float32", value: float32(3.14), expected: true},
		{name: "uintptr", value: uintptr(0xDEADBEEF), expected: true},

		// Pointers
		{name: "pointer to int", value: &five, expected: true},
		{name: "nil pointer", value: nilPtr, expected: false},
		{name: "double pointer", value: &five, expected: true},

		// Custom types
		{name: "custom int type", value: MyInt(5), expected: true},
		{name: "custom float type", value: MyFloat(3.14), expected: true},

		// Non-numeric types
		{name: "string", value: "123", expected: false},
		{name: "bool", value: true, expected: false},
		{name: "struct", value: struct{}{}, expected: false},
		{name: "time.Duration", value: time.Second, expected: true}, // Duration is int64
		{name: "slice", value: []int{1, 2}, expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNumericRule("test_field", tt.value)
			if result != nil {
				t.Errorf("%s: Expected %v, got %v (type: %T)",
					tt.name, tt.expected, result, tt.value)
			}
		})
	}
}

func TestIsInteger(t *testing.T) {
	var nilPtr *int
	five := 5
	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		// Valid integers
		{name: "int", value: 42, expected: true},
		{name: "int8", value: int8(8), expected: true},
		{name: "uint", value: uint(42), expected: true},
		{name: "uintptr", value: uintptr(0xDEADBEEF), expected: true},

		// Pointers
		{name: "pointer to int", value: &five, expected: true},
		{name: "nil pointer", value: nilPtr, expected: false},
		{name: "double pointer", value: &five, expected: true},

		// Custom types
		{name: "custom int", value: MyInt(5), expected: true},
		{name: "custom uint", value: MyUint(5), expected: true},

		// Non-integer types
		{name: "float32", value: float32(3.14), expected: false},
		{name: "string", value: "123", expected: false},
		{name: "bool", value: true, expected: false},
		{name: "time.Duration", value: time.Second, expected: true}, // int64
		{name: "slice", value: []int{1}, expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsIntegerRule("age", tt.value)
			if result != nil {
				t.Errorf("%s: Expected %v, got %v (type: %T)",
					tt.name, tt.expected, result, tt.value)
			}
		})
	}
}
