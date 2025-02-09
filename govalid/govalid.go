package govalid

import (
	"strings"

	"github.com/alaa-aqeel/go-valid/govalid/helpers"
)

var (
	RULE_SEPARATOR = ":"
)

type IMapRules interface {
	Set(name string, rule IRule)
	Get(name string) IRule
}

type Validator struct {
	rules         IMapRules
	rulesValidate map[string]interface{}
	errorMessages ErrorMessages
}

func MakeValidator(rulesValidate map[string]interface{}) *Validator {
	v := &Validator{
		rulesValidate: rulesValidate,
		errorMessages: ErrorMessages{},
	}
	v.RegisterRules()
	return v
}

func (v *Validator) RegisterRule(key string, rule IRule) {
	v.rules.Set(key, rule)
}

func (v *Validator) Validate(data map[string]interface{}) bool {
	for field, rule := range v.rulesValidate {
		value := data[field]
		switch r := rule.(type) {
		case []string:
			for _, ruleStr := range r {
				v.validateRule(field, value, ruleStr)
			}
		case []interface{}:

			for _, ruleI := range r {
				switch rx := ruleI.(type) {
				case RuleCallback:
					if err := rx(field, value); err != nil {
						v.errorMessages.Append(field, err.Error())
					}
					continue
				default:
					if ruleStrStr, ok := rx.(string); ok {
						v.validateRule(field, value, ruleStrStr)
					}
				}
			}
		case RuleCallback:
			if err := r(field, value); err != nil {
				v.errorMessages.Append(field, err.Error())
			}
		default:
			continue
		}
	}

	return v.errorMessages.HasErrors()
}

func (v *Validator) validateRule(field string, value interface{}, rule string) {
	parts := strings.SplitN(rule, ":", 2)
	ruleName := parts[0]
	var ruleParams []interface{}
	if len(parts) > 1 {
		ruleParams = helpers.ToInterfaceSlice(strings.Split(parts[1], ","))
	}

	r := v.rules.Get(ruleName)
	if r == nil {
		return
	}
	r.Valid(field, value, v.errorMessages, ruleParams...)
}

func (v *Validator) Errors() ErrorMessages {

	return v.errorMessages
}

func (v *Validator) HasErrors() bool {

	return v.errorMessages.HasErrors()
}
