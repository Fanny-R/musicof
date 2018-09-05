package core

type Chooser interface {
	Choose() (string, error)
}

type chooser struct {
	provider ChoicesProvider
}

func NewChooser(provider ChoicesProvider) Chooser {
	return &chooser{provider: provider}
}

func (c *chooser) Choose() (string, error) {
	return "coucou", nil
}
