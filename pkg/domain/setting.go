package domain

import (
	"errors"
)

type Setting interface {
	Name() string
	Value() string
}

type setting struct {
	name         string
	defaultValue string
}

func (c setting) Name() string {
	return c.name
}

func (c setting) Value() string {
	return c.defaultValue
}

func New(name string, value string) (Setting, error) {
	if FitConfig().Validate(name) {
		return nil, errors.New("invalid param")
	}
	return &setting{name: name, defaultValue: value}, nil
}
