package json_prettifier

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var vd = validator.New()

type JsonPrettifierConfig struct {
	IndentSize int    `json:"indent_size" validate:"required"`
	IndentType string `json:"indent_type" validate:"required,oneof=space tab"`
}

func (j JsonPrettifierConfig) Validate() error {
	err := vd.Struct(j)
	if err != nil {
		return err
	}

	if j.IndentType == "space" && j.IndentSize%2 != 0 {
		return fmt.Errorf("Indent size for space should be even")
	}

	if j.IndentType == "tab" && j.IndentSize != 1 {
		return fmt.Errorf("Indent size for tab should be 1")
	}

	return nil
}
