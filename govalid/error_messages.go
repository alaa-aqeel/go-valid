package govalid

type ErrorMessages map[string][]string

func (m ErrorMessages) Append(name string, message string) {
	m[name] = append(m[name], message)
}

func (m ErrorMessages) Get(name string) []string {
	value, exsit := m[name]
	if exsit {
		return value
	}
	return nil
}

func (m ErrorMessages) HasErrors() bool {

	return len(m) > 0
}
