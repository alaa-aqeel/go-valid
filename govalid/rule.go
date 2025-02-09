package govalid

type IErrorMessage interface {
	Append(name string, message string)
	Get(name string) []string
}

type RuleCallback func(field string, value interface{}, params ...interface{}) error

type Rule struct {
	Name     string
	Callback RuleCallback
}

func (r *Rule) Valid(field string, value interface{}, errorMessages IErrorMessage, params ...interface{}) {
	if r.Callback == nil {
		return
	}
	if err := r.Callback(field, value, params...); err != nil {
		errorMessages.Append(field, err.Error())
	}
}
