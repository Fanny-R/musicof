package core

type Chooser interface {
	Choose() (string, error)
}

type chooser struct {
}

func NewChooser() Chooser {
	return &chooser{}
}

func (c *chooser) Choose() (string, error) {
	return "coucou", nil
}
