package json_prettifier

import (
	"github.com/go-playground/validator/v10"
)

var vd = validator.New()

type JsonPrettifierConfig struct {
	IndentSize int    `json:"indent_size" validate:"required" help:"Number of spaces or tabs for indentation"`
	IndentType string `json:"indent_type" validate:"required,oneof=space tab" help:"Type of indentation"`
}

func (j JsonPrettifierConfig) Validate() error {
	err := vd.Struct(j)
	if err != nil {
		return err
	}
	return nil
}
