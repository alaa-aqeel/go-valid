package govalid

import "github.com/alaa-aqeel/go-valid/govalid/default_rules"

func (v *Validator) RegisterRules() {
	v.rules = MapRules{}
	v.rules.Set("required", &Rule{
		Name:     "required",
		Callback: default_rules.RequiredRule,
	})
	v.rules.Set("min", &Rule{
		Name:     "min",
		Callback: default_rules.MinRule,
	})
	v.rules.Set("max", &Rule{
		Name:     "max",
		Callback: default_rules.MaxRule,
	})
	v.rules.Set("numeric", &Rule{
		Name:     "numeric",
		Callback: default_rules.IsNumericRule,
	})
	v.rules.Set("integer", &Rule{
		Name:     "integer",
		Callback: default_rules.IsIntegerRule,
	})
}
