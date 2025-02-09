package govalid

type IRule interface {
	Valid(field string, value interface{}, errorMessages IErrorMessage, params ...interface{})
}

type MapRules map[string]IRule

func NewMapRules(rules map[string]IRule) MapRules {

	return MapRules(rules)
}

func (m MapRules) Set(name string, rule IRule) {

	m[name] = rule
}

func (m MapRules) Get(name string) IRule {
	v, ok := m[name]
	if ok {
		return v
	}
	return nil
}
