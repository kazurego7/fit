package domain

import (
	"errors"
	"slices"
)

type FitConfig struct {
	name  string
	value string
}

func (config FitConfig) GetName() string {
	return config.name
}

func (config FitConfig) GetValue() string {
	return config.value
}

func Create(name string, value string) (*FitConfig, error) {
	constNames := []string{
		FitConfigConstant.mainlineType,
	}

	if !slices.Contains(constNames, name) {
		return nil, errors.New("invalid param")
	}
	return &FitConfig{name: name, value: value}, nil
}
