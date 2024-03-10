package domain

import (
	"errors"
)

type Config interface {
	Name() string
	Value() string
}

type config struct {
	name  string
	value string
}

func (c config) Name() string {
	return c.name
}

func (c config) Value() string {
	return c.value
}

func New(name string, value string) (Config, error) {
	if ValidateFitSetting(name) {
		return nil, errors.New("invalid param")
	}
	return &config{name: name, value: value}, nil
}
