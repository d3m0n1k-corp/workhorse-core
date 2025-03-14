package converters

import (
	"errors"
)

type Registration struct {
	Name        string
	DemoInput   any
	Description string
	Config      BaseConfig
	InputType   string
	OutputType  string
	Constructor func(config BaseConfig) BaseConverter
}

var registry = make(map[string]Registration)

func Register(reg Registration) error {
	var _, exists = registry[reg.Name]
	if exists {
		panic("Converter " + reg.Name + " already registered")
	}
	registry[reg.Name] = reg
	return nil
}

func GetRegistration(name string) (*Registration, error) {
	val, ok := registry[name]
	if !ok {
		return nil, errors.New("Converter " + name + " not found")
	}
	return &val, nil
}
