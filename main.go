package main

import (
	"fmt"

	"github.com/alaa-aqeel/go-valid/govalid"
)

func main() {

	v := govalid.MakeValidator(map[string]interface{}{
		"name": []interface{}{"required", govalid.RuleCallback(func(field string, value interface{}, params ...interface{}) error {
			if value.(string) == "John Doe" {
				return fmt.Errorf("name_error")
			}
			return nil
		})},
		"age": []string{"required", "integer", "min:20", "max:100"},
	})
	v.Validate(map[string]interface{}{
		"name": "John Doe",
		"age":  "11.12",
	})
	for _, err := range v.Errors() {
		fmt.Println(err)
	}
}
