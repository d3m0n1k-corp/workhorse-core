package app

import "workhorse-core/internal/converters/base"

func ExecuteConverter(name string, input string, config string) (any, error) {
	conv, err := base.NewConverter(name, config)
	if err != nil {
		return nil, err
	}
	return conv.Apply(input)
}
