package app

import "workhorse-core/internal/converters"

func ExecuteConverter(name string, input string, config string) (any, error) {
	conv, err := converters.NewConverter(name, config)
	if err != nil {
		return nil, err
	}
	return conv.Apply(input)
}
