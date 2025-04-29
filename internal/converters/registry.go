package converters

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/sirupsen/logrus"
)

type Registration struct {
	Name        string
	DemoInput   any
	Description string
	Config      reflect.Type
	InputType   string
	OutputType  string
	Constructor func(config BaseConfig) BaseConverter
}

var registry = make(map[string]*Registration)

func Register(reg *Registration) any {
	var _, exists = registry[reg.Name]
	if exists {
		logrus.Fatalf("Converter %s already registered", reg.Name)
	}
	registry[reg.Name] = reg
	return nil
}

func GetRegistration(name string) (*Registration, error) {
	val, ok := registry[name]
	if !ok {
		return nil, errors.New("Converter " + name + " not found")
	}
	return val, nil
}

func NewConverter(name string, config_str string) (BaseConverter, error) {
	reg, err := GetRegistration(name)
	if err != nil {
		return nil, err
	}
	var instance = reflect.New(reg.Config).Interface()
	err = json.Unmarshal([]byte(config_str), &instance)
	if err != nil {
		return nil, err
	}

	config, ok := instance.(BaseConfig)
	if !ok {
		return nil, errors.New("Invalid config type")
	}
	err = config.Validate()
	if err != nil {
		return nil, err
	}

	return reg.Constructor(config), nil
}

func ListConverters() []*Registration {
	var result []*Registration
	for _, val := range registry {
		result = append(result, val)
	}
	return result
}
