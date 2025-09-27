package converters

import (
	"encoding/json"
	"fmt"
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

// ResetRegistry clears all registered converters - for testing only
func ResetRegistry() {
	registry = make(map[string]*Registration)
}

// GetRegistrySize returns the number of registered converters - for testing only
func GetRegistrySize() int {
	return len(registry)
}

// GetRegistryKeys returns all registered converter names - for testing only
func GetRegistryKeys() []string {
	keys := make([]string, 0, len(registry))
	for k := range registry {
		keys = append(keys, k)
	}
	return keys
}

func Register(reg *Registration) any {
	var _, exists = registry[reg.Name]
	if exists {
		err_str := "Converter " + reg.Name + " already registered"
		logrus.Error(err_str)
		panic(err_str)
	}
	registry[reg.Name] = reg
	return nil
}

func GetRegistration(name string) (*Registration, error) {
	val, ok := registry[name]
	if !ok {
		return nil, fmt.Errorf("converter '%s' not found in registry", name)
	}
	return val, nil
}

func NewConverter(name string, config_str string) (BaseConverter, error) {
	reg, err := GetRegistration(name)
	if err != nil {
		return nil, err
	}

	// Create a pointer to the config struct type
	configPtr := reflect.New(reg.Config)

	// Unmarshal JSON into the pointer
	err = json.Unmarshal([]byte(config_str), configPtr.Interface())
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config for converter '%s': %w", name, err)
	}

	// Get the config value (dereference the pointer)
	configValue := configPtr.Elem().Interface()

	// Type assert to BaseConfig
	config, ok := configValue.(BaseConfig)
	if !ok {
		return nil, fmt.Errorf("config type assertion failed for converter '%s': expected BaseConfig, got %T", name, configValue)
	}

	// Validate the config
	err = config.Validate()
	if err != nil {
		return nil, fmt.Errorf("config validation failed for converter '%s': %w", name, err)
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
