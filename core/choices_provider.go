package core

type ChoicesProvider interface {
	Provide() ([]string, error)
}

type dummyChoicesProvider struct {
	choices []string
}

func NewDummyChoicesProvider(choices []string) ChoicesProvider {
	return &dummyChoicesProvider{choices: choices}
}

func (d *dummyChoicesProvider) Provide() ([]string, error) {
	return d.choices, nil
}
