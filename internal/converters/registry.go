package converters

import (
	"errors"
	"workhorse-core/internal/converters/basec"
)

var registry = make(map[string]basec.BaseConverter)

func Register(name string, converter basec.BaseConverter) error {
	var _, exists = registry[name]
	if exists {
		panic("Converter " + name + " already registered")
	}
	registry[name] = converter
	return nil
}

func GetConverter(name string) (basec.BaseConverter, error) {
	val := registry[name]
	if val == nil {
		return nil, errors.New("Converter " + name + " not found")
	}
	return val, nil
}
