package default_rules

import "testing"

type MyString string // Custom string type for testing

func TestRequiredValidation(t *testing.T) {
	// Test variables for pointer cases
	emptyStr := ""
	nonEmptyStr := "hello"
	spaceStr := "    "

	// Double pointer setup
	nestedEmpty := &emptyStr
	nestedNonEmpty := &nonEmptyStr

	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		// String tests
		{name: "Empty string", value: "", expected: false},
		{name: "Whitespace string", value: "    ", expected: false},
		{name: "Valid string", value: "hello", expected: true},
		{name: "Padded string", value: "  hello  ", expected: true},
		{name: "Custom string type empty", value: MyString(""), expected: false},
		{name: "Custom string type valid", value: MyString("hi"), expected: true},

		// Numeric tests
		{name: "Zero int", value: 0, expected: false},
		{name: "Positive int", value: 42, expected: true},
		{name: "Zero float", value: 0.0, expected: false},
		{name: "Positive float", value: 3.14, expected: true},

		// Pointer tests
		{name: "Nil pointer", value: (*string)(nil), expected: false},
		{name: "Pointer to empty", value: &emptyStr, expected: false},
		{name: "Pointer to valid", value: &nonEmptyStr, expected: true},
		{name: "Pointer to whitespace", value: &spaceStr, expected: false},
		{name: "Double pointer empty", value: &nestedEmpty, expected: false},
		{name: "Double pointer valid", value: &nestedNonEmpty, expected: true},

		// Collection tests
		{name: "Empty slice", value: []int{}, expected: false},
		{name: "Non-empty slice", value: []int{1}, expected: true},
		{name: "Empty map", value: map[string]int{}, expected: false},
		{name: "Non-empty map", value: map[string]int{"a": 1}, expected: true},
		{name: "Empty array", value: [0]int{}, expected: false},
		{name: "Non-empty array", value: [1]int{1}, expected: true},

		// Interface and special types
		{name: "Nil interface", value: nil, expected: false},
		{name: "Valid interface", value: interface{}("test"), expected: true},
		{name: "False bool", value: false, expected: false},
		{name: "True bool", value: true, expected: true},

		// Struct tests
		{name: "Zero struct", value: struct{}{}, expected: false},
		{name: "Non-zero struct", value: struct{ Name string }{Name: "John"}, expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RequiredRule("testField", tt.value)
			if got != nil {
				t.Errorf("Case '%s': Expected %v, got %v (type: %T, value: %+v)",
					tt.name, tt.expected, got, tt.value, tt.value)
			}
		})
	}
}
